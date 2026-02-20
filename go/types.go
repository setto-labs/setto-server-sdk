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

// ---- Integration / Verification types ----

// VerificationStatus holds the result of a verification status query.
type VerificationStatus struct {
	IsPhoneVerified bool
	VerifiedAt      int64 // Unix ms, 0 if not verified
}

// AccountLinkDirectResult holds the result of a S2S direct account link.
type AccountLinkDirectResult struct {
	UserID          string
	Email           string
	IsPhoneVerified bool
	IsNewUser       bool
}

// ---- Profile types ----

// PayerProfile holds the payer's profile for a payment.
type PayerProfile struct {
	SettoID     string
	DisplayName string
	PhotoURL    string
	ETag        string
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

type getVerificationStatusResponse struct {
	IsPhoneVerified bool  `json:"is_phone_verified"`
	VerifiedAt      int64 `json:"verified_at"`
}

type linkAccountDirectRequest struct {
	IDToken string `json:"id_token"`
}

type linkAccountDirectResponse struct {
	UserID          string `json:"user_id"`
	Email           string `json:"email"`
	IsPhoneVerified bool   `json:"is_phone_verified"`
	IsNewUser       bool   `json:"is_new_user"`
}

type getPayerProfileResponse struct {
	SettoID     string `json:"setto_id"`
	DisplayName string `json:"display_name"`
	PhotoURL    string `json:"photo_url"`
	ETag        string `json:"etag"`
}

// ---- InitiatePayment types ----

// InitiatePaymentRequest is the request for initiating a payment.
type InitiatePaymentRequest struct {
	MerchantID      string `json:"merchant_id"`
	Amount          string `json:"amount"`
	ChainID         int32  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	WalletType      string `json:"wallet_type"`
	SettoUserID     string `json:"setto_user_id"`
}

// InitiatePaymentResponse is the response from payment initiation.
type InitiatePaymentResponse struct {
	PaymentID       string `json:"payment_id"`
	MerchantID      string `json:"merchant_id"`
	PoolAddress     string `json:"pool_address"`
	Amount          string `json:"amount"`
	ChainID         int32  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	ExpiresAt       int64  `json:"expires_at"`
	CreatedAt       int64  `json:"created_at"`
	FeeAmount       string `json:"fee_amount"`
	MerchantAddress string `json:"merchant_address,omitempty"`
	Deadline        int64  `json:"deadline,omitempty"`
}

type initiatePaymentWireRequest struct {
	MerchantID      string `json:"merchant_id"`
	Amount          string `json:"amount"`
	ChainID         int32  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	WalletType      string `json:"wallet_type"`
	SettoUserID     string `json:"setto_user_id"`
}

type initiatePaymentWireResponse struct {
	PaymentID       string `json:"payment_id"`
	MerchantID      string `json:"merchant_id"`
	PoolAddress     string `json:"pool_address"`
	Amount          string `json:"amount"`
	ChainID         int32  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	ExpiresAt       int64  `json:"expires_at"`
	CreatedAt       int64  `json:"created_at"`
	FeeAmount       string `json:"fee_amount"`
	MerchantAddress string `json:"merchant_address,omitempty"`
	Deadline        int64  `json:"deadline,omitempty"`
}
