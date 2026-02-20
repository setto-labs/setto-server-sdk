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

// LinkAccountDirect performs S2S direct account linking via IdP token.
// The IdP token is verified by setto-server which matches/creates the user.
// No OTT intermediary required.
func (c *Client) LinkAccountDirect(ctx context.Context, idToken string) (*AccountLinkDirectResult, error) {
	reqBody := &linkAccountDirectRequest{IDToken: idToken}

	var raw linkAccountDirectResponse
	if err := c.do(ctx, "POST", "/api/integration/link-account-direct", reqBody, &raw); err != nil {
		return nil, fmt.Errorf("link account direct: %w", err)
	}

	return &AccountLinkDirectResult{
		UserID:          raw.UserID,
		Email:           raw.Email,
		IsPhoneVerified: raw.IsPhoneVerified,
		IsNewUser:       raw.IsNewUser,
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
