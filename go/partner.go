package setto

import (
	"context"
	"fmt"
)

// GetVerificationStatus checks if a user has completed phone verification.
// Used before store/merchant creation to enforce verification requirement.
func (c *Client) GetVerificationStatus(ctx context.Context, userID string) (*VerificationStatus, error) {
	var raw getVerificationStatusResponse
	if err := c.do(ctx, "GET", "/api/partner/user/"+userID+"/verification", nil, &raw); err != nil {
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
	if err := c.do(ctx, "POST", "/api/partner/exchange-link-token", reqBody, &raw); err != nil {
		return nil, fmt.Errorf("exchange account link token: %w", err)
	}

	return &AccountLinkInfo{
		UserID:          raw.UserID,
		Email:           raw.Email,
		IsPhoneVerified: raw.IsPhoneVerified,
	}, nil
}
