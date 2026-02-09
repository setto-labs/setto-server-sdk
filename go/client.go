// Package setto provides the Setto Server SDK for partner server integration.
//
// This SDK communicates with Setto Wallet Server via REST API.
// Used by partner servers to manage merchants, verify users, check payment status,
// and verify Setto Wallet ID Tokens (JWKS/OIDC).
//
// Quick start:
//
//	client, err := setto.NewClient(setto.Config{
//	    APIKey:      "sk_partner.xxx",
//	    Environment: setto.Production,
//	})
//
//	merchant, err := client.CreateMerchant(ctx, &setto.CreateMerchantRequest{...})
package setto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Environment represents the Setto environment.
type Environment int

const (
	// Production is the live environment. HTTPS is enforced.
	Production Environment = iota
	// Development is the test environment.
	Development
)

const (
	productionURL  = "https://wallet.settopay.com"
	developmentURL = "https://dev-wallet.settopay.com"
	defaultTimeout = 30 * time.Second
	sdkVersion     = "0.1.0"
)

// Config holds configuration for creating a Client.
type Config struct {
	APIKey      string      // Partner API Key (sk_partner.xxx)
	Environment Environment // Production or Development
}

// Option configures the Client.
type Option func(*clientOptions)

type clientOptions struct {
	timeout    time.Duration
	httpClient *http.Client
	baseURL    string
}

// WithTimeout sets the HTTP client timeout. Default: 30s.
func WithTimeout(d time.Duration) Option {
	return func(o *clientOptions) { o.timeout = d }
}

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(c *http.Client) Option {
	return func(o *clientOptions) { o.httpClient = c }
}

// WithBaseURL overrides the environment URL.
func WithBaseURL(url string) Option {
	return func(o *clientOptions) { o.baseURL = url }
}

// Client is the Setto Wallet SDK client.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Setto SDK client.
func NewClient(cfg Config, opts ...Option) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("setto: API key is required")
	}
	if !strings.HasPrefix(cfg.APIKey, "sk_partner.") {
		return nil, fmt.Errorf("setto: API key must start with 'sk_partner.'")
	}

	options := &clientOptions{timeout: defaultTimeout}
	for _, opt := range opts {
		opt(options)
	}

	baseURL := options.baseURL
	if baseURL == "" {
		switch cfg.Environment {
		case Production:
			baseURL = productionURL
		case Development:
			baseURL = developmentURL
		}
	}

	if cfg.Environment == Production && !strings.HasPrefix(baseURL, "https://") {
		return nil, fmt.Errorf("setto: HTTPS is required in production (got %s)", baseURL)
	}

	httpClient := options.httpClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: options.timeout}
	}

	return &Client{
		apiKey:     cfg.APIKey,
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}, nil
}

// NewVerifier creates a JWT Verifier configured from this client's base URL.
func (c *Client) NewVerifier() *Verifier {
	return NewVerifier(
		c.baseURL+"/.well-known/jwks.json",
		c.baseURL,
	)
}

// do executes an HTTP request with authentication and error handling.
func (c *Client) do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("setto: failed to marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return &NetworkError{Cause: err}
	}

	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "setto-server-sdk-go/"+sdkVersion)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &NetworkError{Cause: err}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &NetworkError{Cause: fmt.Errorf("read response: %w", err)}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseHTTPError(resp.StatusCode, respBody)
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("setto: failed to parse response: %w", err)
		}
	}

	return nil
}

// parseHTTPError parses a non-2xx HTTP response into a WalletError.
func parseHTTPError(statusCode int, body []byte) error {
	we := &WalletError{HTTPStatus: statusCode}
	if len(body) > 0 {
		_ = json.Unmarshal(body, we)
	}
	if we.SystemError != "" && we.SystemError != SystemOK {
		we.Code = we.SystemError
	} else if we.PaymentError != "" && we.PaymentError != PaymentOK {
		we.Code = we.PaymentError
	} else if we.ValidationError != "" && we.ValidationError != ValidationOK {
		we.Code = we.ValidationError
	}
	return we
}
