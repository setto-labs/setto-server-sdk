package setto

import (
	"context"
	"fmt"
)

// GetVerificationStatus checks if a user has completed phone verification.
// Used before store/merchant creation to enforce verification requirement.
func (c *Client) GetVerificationStatus(ctx context.Context, userID string) (*VerificationStatus, error) {
	var raw getVerificationStatusResponse
	if err := c.do(ctx, "GET", "/api/integration/user/"+userID+"/verification", nil, &raw); err != nil {
		return nil, fmt.Errorf("get verification status: %w", err)
	}

	return &VerificationStatus{
		IsPhoneVerified: raw.IsPhoneVerified,
		VerifiedAt:      raw.VerifiedAt,
	}, nil
}

// ExchangeAccountLinkToken exchanges a one-time link token for user info.
// The link token is atomically consumed; replay is impossible.
func (c *Client) ExchangeAccountLinkToken(ctx context.Context, linkToken string) (*AccountLinkInfo, error) {
	reqBody := &exchangeAccountLinkTokenRequest{LinkToken: linkToken}

	var raw exchangeAccountLinkTokenResponse
	if err := c.do(ctx, "POST", "/api/integration/exchange-link-token", reqBody, &raw); err != nil {
		return nil, fmt.Errorf("exchange account link token: %w", err)
	}

	return &AccountLinkInfo{
		UserID:          raw.UserID,
		Email:           raw.Email,
		IsPhoneVerified: raw.IsPhoneVerified,
	}, nil
}

// GetPayerProfile returns the payer's profile for a given payment.
// Used by external integrations to display payer info (name, photo) without exposing email.
func (c *Client) GetPayerProfile(ctx context.Context, paymentID string) (*PayerProfile, error) {
	var raw getPayerProfileResponse
	if err := c.do(ctx, "GET", "/api/external/payment/"+paymentID+"/payer", nil, &raw); err != nil {
		return nil, fmt.Errorf("get payer profile: %w", err)
	}

	return &PayerProfile{
		SettoID:     raw.SettoID,
		DisplayName: raw.DisplayName,
		PhotoURL:    raw.PhotoURL,
		ETag:        raw.ETag,
	}, nil
}

// InitiatePayment creates a new payment session and returns payment information.
// The server generates a payment_id (SSoT) and the SDK/client uses it to execute the payment.
// Auth: X-API-Key (external integration)
func (c *Client) InitiatePayment(ctx context.Context, req *InitiatePaymentRequest) (*InitiatePaymentResponse, error) {
	wireReq := &initiatePaymentWireRequest{
		MerchantID:      req.MerchantID,
		Amount:          req.Amount,
		ChainID:         req.ChainID,
		ContractAddress: req.ContractAddress,
		WalletType:      req.WalletType,
		SettoUserID:     req.SettoUserID,
	}

	var raw initiatePaymentWireResponse
	if err := c.do(ctx, "POST", "/api/integration/payment/initiate", wireReq, &raw); err != nil {
		return nil, fmt.Errorf("initiate payment: %w", err)
	}

	return &InitiatePaymentResponse{
		PaymentID:       raw.PaymentID,
		MerchantID:      raw.MerchantID,
		PoolAddress:     raw.PoolAddress,
		Amount:          raw.Amount,
		ChainID:         raw.ChainID,
		ContractAddress: raw.ContractAddress,
		ExpiresAt:       raw.ExpiresAt,
		CreatedAt:       raw.CreatedAt,
		FeeAmount:       raw.FeeAmount,
		MerchantAddress: raw.MerchantAddress,
		Deadline:        raw.Deadline,
	}, nil
}
