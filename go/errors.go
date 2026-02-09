package setto

import (
	"errors"
	"fmt"
)

// System error codes.
const (
	SystemOK          = "SYSTEM_OK"
	SystemInternal    = "SYSTEM_INTERNAL"
	SystemRPCFailed   = "SYSTEM_RPC_FAILED"
	SystemRateLimited = "SYSTEM_RATE_LIMITED"
)

// Payment error codes.
const (
	PaymentOK                       = "PAYMENT_OK"
	PaymentNotFound                 = "PAYMENT_NOT_FOUND"
	PaymentMerchantNotFound         = "PAYMENT_MERCHANT_NOT_FOUND"
	PaymentMerchantNameRequired     = "PAYMENT_MERCHANT_NAME_REQUIRED"
	PaymentPayoutAddressRequired    = "PAYMENT_PAYOUT_ADDRESS_REQUIRED"
	PaymentInvalidEVMAddress        = "PAYMENT_INVALID_EVM_ADDRESS"
	PaymentInvalidSVMAddress        = "PAYMENT_INVALID_SVM_ADDRESS"
	PaymentOTTRequired              = "PAYMENT_OTT_REQUIRED"
	PaymentOTTInvalid               = "PAYMENT_OTT_INVALID"
	PaymentOTTExpired               = "PAYMENT_OTT_EXPIRED"
	PaymentOTTAlreadyUsed           = "PAYMENT_OTT_ALREADY_USED"
	PaymentOTTScopeMismatch         = "PAYMENT_OTT_SCOPE_MISMATCH"
	PaymentStoreLimitExceeded       = "PAYMENT_STORE_LIMIT_EXCEEDED"
	PaymentAmountRequired           = "PAYMENT_AMOUNT_REQUIRED"
	PaymentAmountTooLow             = "PAYMENT_AMOUNT_TOO_LOW"
	PaymentAmountTooHigh            = "PAYMENT_AMOUNT_TOO_HIGH"
	PaymentAmountInvalidFormat      = "PAYMENT_AMOUNT_INVALID_FORMAT"
	PaymentProductNameRequired      = "PAYMENT_PRODUCT_NAME_REQUIRED"
	PaymentProductNameTooShort      = "PAYMENT_PRODUCT_NAME_TOO_SHORT"
	PaymentProductNameTooLong       = "PAYMENT_PRODUCT_NAME_TOO_LONG"
	PaymentProductDescTooLong       = "PAYMENT_PRODUCT_DESC_TOO_LONG"
	PaymentInvalidStock             = "PAYMENT_INVALID_STOCK"
	PaymentProductTagTooLong        = "PAYMENT_PRODUCT_TAG_TOO_LONG"
	PaymentProductMainImagesTooMany = "PAYMENT_PRODUCT_MAIN_IMAGES_TOO_MANY"
	PaymentProductDetailImagesTooMany = "PAYMENT_PRODUCT_DETAIL_IMAGES_TOO_MANY"
	PaymentProductLimitExceeded     = "PAYMENT_PRODUCT_LIMIT_EXCEEDED"
)

// Validation error codes.
const (
	ValidationOK             = "VALIDATION_OK"
	ValidationRequiredField  = "VALIDATION_REQUIRED_FIELD"
	ValidationInvalidFormat  = "VALIDATION_INVALID_FORMAT"
	ValidationInvalidAddress = "VALIDATION_INVALID_ADDRESS"
	ValidationInvalidAmount  = "VALIDATION_INVALID_AMOUNT"
	ValidationInvalidChainID = "VALIDATION_INVALID_CHAIN_ID"
	ValidationInvalidID      = "VALIDATION_INVALID_ID"
	ValidationInvalidRequest = "VALIDATION_INVALID_REQUEST"
)

// WalletError represents a structured error from Setto Wallet Server.
type WalletError struct {
	SystemError     string `json:"system_error,omitempty"`
	PaymentError    string `json:"payment_error,omitempty"`
	ValidationError string `json:"validation_error,omitempty"`
	Code            string `json:"-"`
	Message         string `json:"-"`
	HTTPStatus      int    `json:"-"`
}

func (e *WalletError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Code != "" {
		return fmt.Sprintf("setto: %s", e.Code)
	}
	return fmt.Sprintf("setto: HTTP %d", e.HTTPStatus)
}

// IsWalletError checks if the error is a WalletError and returns it.
func IsWalletError(err error) (*WalletError, bool) {
	var walletErr *WalletError
	if errors.As(err, &walletErr) {
		return walletErr, true
	}
	return nil, false
}

// NetworkError represents a network-level error.
type NetworkError struct {
	Cause error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("setto: network error: %v", e.Cause)
}

func (e *NetworkError) Unwrap() error {
	return e.Cause
}

// JWT verification errors.
var (
	ErrTokenInvalid     = errors.New("setto: token is invalid")
	ErrTokenExpired     = errors.New("setto: token has expired")
	ErrIssuerMismatch   = errors.New("setto: issuer mismatch")
	ErrKeyNotFound      = errors.New("setto: signing key not found")
	ErrEmailNotVerified = errors.New("setto: email not verified")
)

// checkResponseErrors converts embedded gRPC-Gateway error fields to WalletError.
func checkResponseErrors(sysErr, payErr, valErr string) error {
	if sysErr != "" && sysErr != SystemOK {
		return &WalletError{SystemError: sysErr, Code: sysErr, Message: fmt.Sprintf("system error: %s", sysErr)}
	}
	if payErr != "" && payErr != PaymentOK {
		return &WalletError{PaymentError: payErr, Code: payErr, Message: fmt.Sprintf("payment error: %s", payErr)}
	}
	if valErr != "" && valErr != ValidationOK {
		return &WalletError{ValidationError: valErr, Code: valErr, Message: fmt.Sprintf("validation error: %s", valErr)}
	}
	return nil
}
