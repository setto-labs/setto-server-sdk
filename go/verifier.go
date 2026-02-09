package setto

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// Verifier verifies Setto Wallet ID Tokens using JWKS.
// Thread-safe and handles JWKS caching/refresh per OIDC standard.
type Verifier struct {
	jwksURL string
	issuer  string

	mu        sync.RWMutex
	cache     *jwk.Cache
	cachedSet jwk.Set
	cacheCtx  context.Context
}

// NewVerifier creates a new Wallet ID Token verifier.
// JWKS is NOT fetched at this point; it's fetched lazily on first VerifyIDToken call.
func NewVerifier(jwksURL, issuer string) *Verifier {
	return &Verifier{
		jwksURL: jwksURL,
		issuer:  issuer,
	}
}

// VerifyIDToken verifies a Wallet ID Token and returns the claims.
// JWKS is fetched lazily and cached. If kid is not found, JWKS is re-fetched.
func (v *Verifier) VerifyIDToken(ctx context.Context, idToken string) (*Claims, error) {
	if err := v.ensureCache(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize JWKS cache: %w", err)
	}

	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid not found in token header")
		}

		return v.getKeyByKid(ctx, kid)
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, fmt.Errorf("%w: %v", ErrTokenInvalid, err)
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenInvalid
	}

	iss, _ := mapClaims["iss"].(string)
	if iss != v.issuer {
		return nil, fmt.Errorf("%w: expected %s, got %s", ErrIssuerMismatch, v.issuer, iss)
	}

	claims := &Claims{}
	if sub, ok := mapClaims["sub"].(string); ok {
		claims.UserID = sub
	}
	if email, ok := mapClaims["email"].(string); ok {
		claims.Email = email
	}
	if verified, ok := mapClaims["email_verified"].(bool); ok {
		claims.EmailVerified = verified
	}
	if iat, ok := mapClaims["iat"].(float64); ok {
		claims.IssuedAt = time.Unix(int64(iat), 0)
	}
	if exp, ok := mapClaims["exp"].(float64); ok {
		claims.ExpiresAt = time.Unix(int64(exp), 0)
	}

	return claims, nil
}

// VerifyIDTokenRequireEmail verifies token and ensures email is verified.
func (v *Verifier) VerifyIDTokenRequireEmail(ctx context.Context, idToken string) (*Claims, error) {
	claims, err := v.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	if !claims.EmailVerified {
		return nil, ErrEmailNotVerified
	}
	return claims, nil
}

func (v *Verifier) ensureCache(ctx context.Context) error {
	v.mu.RLock()
	if v.cache != nil {
		v.mu.RUnlock()
		return nil
	}
	v.mu.RUnlock()

	v.mu.Lock()
	defer v.mu.Unlock()

	if v.cache != nil {
		return nil
	}

	v.cacheCtx = context.Background()
	v.cache = jwk.NewCache(v.cacheCtx)
	v.cache.Register(v.jwksURL, jwk.WithMinRefreshInterval(15*time.Minute))
	v.cachedSet = jwk.NewCachedSet(v.cache, v.jwksURL)

	return nil
}

func (v *Verifier) getKeyByKid(ctx context.Context, kid string) (interface{}, error) {
	v.mu.RLock()
	cachedSet := v.cachedSet
	cache := v.cache
	v.mu.RUnlock()

	key, found := cachedSet.LookupKeyID(kid)
	if found {
		var rawKey interface{}
		if err := key.Raw(&rawKey); err != nil {
			return nil, fmt.Errorf("failed to get raw key: %w", err)
		}
		return rawKey, nil
	}

	_, err := cache.Refresh(ctx, v.jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh JWKS: %w", err)
	}

	key, found = cachedSet.LookupKeyID(kid)
	if !found {
		return nil, ErrKeyNotFound
	}

	var rawKey interface{}
	if err := key.Raw(&rawKey); err != nil {
		return nil, fmt.Errorf("failed to get raw key: %w", err)
	}

	return rawKey, nil
}
