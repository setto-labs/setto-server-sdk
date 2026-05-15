package setto

import "time"

// PaymentStatus represents the status of a payment.
// Values match proto enum PaymentStatus JSON serialization (gRPC-Gateway UPPER_SNAKE_CASE).
type PaymentStatus string

const (
	PaymentStatusUnspecified  PaymentStatus = "PAYMENT_STATUS_UNSPECIFIED"
	PaymentStatusPending      PaymentStatus = "PAYMENT_STATUS_PENDING"
	PaymentStatusProcessing   PaymentStatus = "PAYMENT_STATUS_PROCESSING"
	PaymentStatusSubmitted    PaymentStatus = "PAYMENT_STATUS_SUBMITTED"
	PaymentStatusIncluded     PaymentStatus = "PAYMENT_STATUS_INCLUDED"
	PaymentStatusConfirmed    PaymentStatus = "PAYMENT_STATUS_CONFIRMED"
	PaymentStatusFinalized    PaymentStatus = "PAYMENT_STATUS_FINALIZED"
	PaymentStatusFailed       PaymentStatus = "PAYMENT_STATUS_FAILED"
	PaymentStatusRefundPending PaymentStatus = "PAYMENT_STATUS_REFUND_PENDING"
	PaymentStatusCancelled    PaymentStatus = "PAYMENT_STATUS_CANCELLED"
)

// WalletType represents the type of wallet used for payment.
// Values match proto enum WalletType JSON serialization (gRPC-Gateway UPPER_SNAKE_CASE).
type WalletType string

const (
	WalletTypeUnspecified WalletType = "WALLET_TYPE_UNSPECIFIED"
	WalletTypeSetto       WalletType = "WALLET_TYPE_SETTO"
	WalletTypeMetamask    WalletType = "WALLET_TYPE_METAMASK"
	WalletTypeOKX         WalletType = "WALLET_TYPE_OKX"
	WalletTypeCoinbase    WalletType = "WALLET_TYPE_COINBASE"
	WalletTypePhantom     WalletType = "WALLET_TYPE_PHANTOM"
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
// Fields match proto GetExternalPaymentStatusResponse (gRPC-Gateway snake_case JSON).
type PaymentInfo struct {
	PaymentID       string        `json:"payment_id"`
	Status          PaymentStatus `json:"status"`
	TxHash          string        `json:"tx_hash,omitempty"`
	Amount          string        `json:"amount"`
	Currency        string        `json:"currency"`
	CreatedAt       int64         `json:"created_at"`
	CompletedAt     int64         `json:"completed_at,omitempty"`
	Decimals        uint32        `json:"decimals,omitempty"`
	SenderAddress   string        `json:"sender_address,omitempty"`
	PoolAddress     string        `json:"pool_address,omitempty"`
	ChainID         int32         `json:"chain_id,omitempty"`
	ContractAddress string        `json:"contract_address,omitempty"`
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
	MerchantID      string     `json:"merchant_id"`
	Amount          string     `json:"amount"`
	ChainID         int32      `json:"chain_id"`
	ContractAddress string     `json:"contract_address"`
	WalletType      WalletType `json:"wallet_type"`
	SettoUserID     string     `json:"setto_user_id"`
}

// InitiatePaymentResponse is the response from payment initiation.
// Fields match proto InitiatePaymentResponse (gRPC-Gateway snake_case JSON).
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
	Decimals        int32  `json:"decimals,omitempty"`
}

type initiatePaymentWireRequest struct {
	MerchantID      string     `json:"merchant_id"`
	Amount          string     `json:"amount"`
	ChainID         int32      `json:"chain_id"`
	ContractAddress string     `json:"contract_address"`
	WalletType      WalletType `json:"wallet_type"`
	SettoUserID     string     `json:"setto_user_id"`
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
	Decimals        int32  `json:"decimals,omitempty"`
}
