package setto

import "time"

// PaymentStatus represents the status of a payment.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSubmitted PaymentStatus = "submitted"
	PaymentStatusIncluded  PaymentStatus = "included"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// ---- Merchant types ----

// CreateMerchantRequest is the request for creating a merchant.
type CreateMerchantRequest struct {
	Email            string `json:"email,omitempty"`
	Name             string `json:"name"`
	PayoutEVMAddress string `json:"payout_evm_address"`
	PayoutSVMAddress string `json:"payout_svm_address,omitempty"`
	FeeRate          string `json:"fee_rate,omitempty"`
	OneTimeToken     string `json:"one_time_token,omitempty"`
}

// CreateMerchantResponse is the response from creating a merchant.
type CreateMerchantResponse struct {
	MerchantID string
}

// GetMerchantResponse is the response from getting merchant info.
type GetMerchantResponse struct {
	MerchantID       string
	PayoutEVMAddress string
	PayoutSVMAddress string
}

// UpdateMerchantRequest is the request for updating a merchant.
type UpdateMerchantRequest struct {
	MerchantID       string `json:"-"`
	OneTimeToken     string `json:"one_time_token,omitempty"`
	PayoutEVMAddress string `json:"payout_evm_address,omitempty"`
	PayoutSVMAddress string `json:"payout_svm_address,omitempty"`
}

// UpdateMerchantResponse is the response from updating a merchant.
type UpdateMerchantResponse struct {
	MerchantID       string
	PayoutEVMAddress string
	PayoutSVMAddress string
}

// ---- Partner / Verification types ----

// VerificationStatus holds the result of a verification status query.
type VerificationStatus struct {
	IsPhoneVerified bool
	VerifiedAt      int64 // Unix ms, 0 if not verified
}

// AccountLinkInfo holds the result of a link token exchange.
type AccountLinkInfo struct {
	UserID          string
	Email           string
	IsPhoneVerified bool
}

// ---- Payment types ----

// PaymentInfo represents the payment information.
type PaymentInfo struct {
	PaymentID   string        `json:"paymentId"`
	Status      PaymentStatus `json:"status"`
	TxHash      string        `json:"txHash,omitempty"`
	Amount      string        `json:"amount"`
	Currency    string        `json:"currency"`
	CreatedAt   int64         `json:"createdAt"`
	CompletedAt int64         `json:"completedAt,omitempty"`
}

// ---- JWT Claims ----

// Claims represents the verified claims from a Setto Wallet ID Token.
type Claims struct {
	UserID        string
	Email         string
	EmailVerified bool
	IssuedAt      time.Time
	ExpiresAt     time.Time
}

// ---- Internal wire types (gRPC-Gateway JSON format) ----

type createMerchantResponse struct {
	SystemError     string `json:"system_error"`
	PaymentError    string `json:"payment_error"`
	ValidationError string `json:"validation_error"`
	MerchantID      string `json:"merchant_id"`
}

type getMerchantResponse struct {
	SystemError      string `json:"system_error"`
	PaymentError     string `json:"payment_error"`
	MerchantID       string `json:"merchant_id"`
	PayoutEVMAddress string `json:"payout_evm_address"`
	PayoutSVMAddress string `json:"payout_svm_address"`
}

type updateMerchantResponse struct {
	SystemError      string `json:"system_error"`
	PaymentError     string `json:"payment_error"`
	ValidationError  string `json:"validation_error"`
	MerchantID       string `json:"merchant_id"`
	PayoutEVMAddress string `json:"payout_evm_address"`
	PayoutSVMAddress string `json:"payout_svm_address"`
}

type getVerificationStatusResponse struct {
	IsPhoneVerified bool  `json:"is_phone_verified"`
	VerifiedAt      int64 `json:"verified_at"`
}

type exchangeAccountLinkTokenRequest struct {
	LinkToken string `json:"link_token"`
}

type exchangeAccountLinkTokenResponse struct {
	UserID          string `json:"user_id"`
	Email           string `json:"email"`
	IsPhoneVerified bool   `json:"is_phone_verified"`
}
