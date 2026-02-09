package setto

import (
	"context"
	"fmt"
)

// CreateMerchant creates a new merchant in Setto Wallet.
//
// User identification (one of):
//   - Email: Platform partner creates on behalf of user (user auto-created if not found)
//   - OneTimeToken: Individual partner creates directly (user_id extracted from OTT)
func (c *Client) CreateMerchant(ctx context.Context, req *CreateMerchantRequest) (*CreateMerchantResponse, error) {
	var raw createMerchantResponse
	if err := c.do(ctx, "POST", "/api/merchant", req, &raw); err != nil {
		return nil, fmt.Errorf("create merchant: %w", err)
	}

	if err := checkResponseErrors(raw.SystemError, raw.PaymentError, raw.ValidationError); err != nil {
		return nil, err
	}

	return &CreateMerchantResponse{
		MerchantID: raw.MerchantID,
	}, nil
}

// GetMerchant retrieves merchant information.
func (c *Client) GetMerchant(ctx context.Context, merchantID string) (*GetMerchantResponse, error) {
	var raw getMerchantResponse
	if err := c.do(ctx, "GET", "/api/merchant/"+merchantID, nil, &raw); err != nil {
		return nil, fmt.Errorf("get merchant: %w", err)
	}

	if err := checkResponseErrors(raw.SystemError, raw.PaymentError, ""); err != nil {
		return nil, err
	}

	return &GetMerchantResponse{
		MerchantID:       raw.MerchantID,
		Name:             raw.Name,
		PhotoURL:         raw.PhotoURL,
		PayoutEVMAddress: raw.PayoutEVMAddress,
		PayoutSVMAddress: raw.PayoutSVMAddress,
	}, nil
}

// UpdateMerchant updates merchant wallet addresses.
// Requires a One-Time Token with scope UPDATE_MERCHANT.
func (c *Client) UpdateMerchant(ctx context.Context, req *UpdateMerchantRequest) (*UpdateMerchantResponse, error) {
	var raw updateMerchantResponse
	if err := c.do(ctx, "PUT", "/api/merchant/"+req.MerchantID, req, &raw); err != nil {
		return nil, fmt.Errorf("update merchant: %w", err)
	}

	if err := checkResponseErrors(raw.SystemError, raw.PaymentError, raw.ValidationError); err != nil {
		return nil, err
	}

	return &UpdateMerchantResponse{
		MerchantID:       raw.MerchantID,
		Name:             raw.Name,
		PhotoURL:         raw.PhotoURL,
		PayoutEVMAddress: raw.PayoutEVMAddress,
		PayoutSVMAddress: raw.PayoutSVMAddress,
	}, nil
}

// UpdateMerchantProfile updates only merchant display info (name, photo_url).
// No OTT required — uses Partner API Key authentication only.
// Used by Commerce Server for automatic sync when store info is updated.
func (c *Client) UpdateMerchantProfile(ctx context.Context, req *UpdateMerchantProfileRequest) (*UpdateMerchantProfileResponse, error) {
	var raw updateMerchantProfileResponse
	if err := c.do(ctx, "PATCH", "/api/merchant/"+req.MerchantID+"/profile", req, &raw); err != nil {
		return nil, fmt.Errorf("update merchant profile: %w", err)
	}

	if err := checkResponseErrors(raw.SystemError, raw.PaymentError, raw.ValidationError); err != nil {
		return nil, err
	}

	return &UpdateMerchantProfileResponse{
		MerchantID: raw.MerchantID,
		Name:       raw.Name,
		PhotoURL:   raw.PhotoURL,
	}, nil
}
