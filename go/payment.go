package setto

import (
	"context"
	"fmt"
)

// GetPaymentStatus retrieves the status of a payment.
func (c *Client) GetPaymentStatus(ctx context.Context, paymentID string) (*PaymentInfo, error) {
	var info PaymentInfo
	if err := c.do(ctx, "GET", "/api/external/payment/"+paymentID, nil, &info); err != nil {
		return nil, fmt.Errorf("get payment status: %w", err)
	}
	return &info, nil
}

// IsPaymentComplete returns true if the payment has been successfully completed.
func (p *PaymentInfo) IsPaymentComplete() bool {
	return p.Status == PaymentStatusIncluded
}

// IsPaymentFailed returns true if the payment has failed.
func (p *PaymentInfo) IsPaymentFailed() bool {
	return p.Status == PaymentStatusFailed || p.Status == PaymentStatusCancelled
}

// IsPaymentPending returns true if the payment is still pending.
func (p *PaymentInfo) IsPaymentPending() bool {
	return p.Status == PaymentStatusPending || p.Status == PaymentStatusSubmitted
}
