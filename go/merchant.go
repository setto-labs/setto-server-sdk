package setto

import (
	"context"
	"fmt"
)

// CreateMerchant creates a new merchant in Setto Wallet.
// The merchant owner is determined by the partner's authenticated identity.
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

