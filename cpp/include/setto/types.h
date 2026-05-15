#pragma once

#include <cstdint>
#include <optional>
#include <string>

namespace setto {

enum class Environment {
    Production,
    Development,
};

struct Config {
    std::string api_key;
    Environment environment;
    uint32_t timeout_ms = 30000;
    std::optional<std::string> base_url;
};

// Merchant types

struct CreateMerchantRequest {
    std::string name;
    std::string payout_evm_address;
    std::optional<std::string> photo_url;
    std::optional<std::string> payout_svm_address;
    std::optional<std::string> fee_rate;
};

struct CreateMerchantResponse {
    std::string merchant_id;
};

struct GetMerchantResponse {
    std::string merchant_id;
    std::string name;
    std::string photo_url;
    std::string payout_evm_address;
    std::string payout_svm_address;
};

/// Update merchant wallet addresses.
struct UpdateMerchantRequest {
    std::string merchant_id;
    std::optional<std::string> name;
    std::optional<std::string> photo_url;
    std::optional<std::string> payout_evm_address;
    std::optional<std::string> payout_svm_address;
};

struct UpdateMerchantResponse {
    std::string merchant_id;
    std::string name;
    std::string photo_url;
    std::string payout_evm_address;
    std::string payout_svm_address;
};

// Verification types

struct VerificationStatus {
    bool is_phone_verified;
    int64_t verified_at;
};

struct AccountLinkInfo {
    std::string user_id;
    std::string email;
    bool is_phone_verified;
};

// Payment types

/// Payment status values matching proto PaymentStatus enum (gRPC-Gateway UPPER_SNAKE_CASE).
enum class PaymentStatus {
    Unspecified,
    Pending,
    Processing,
    Submitted,
    Included,
    Confirmed,
    Finalized,
    Failed,
    RefundPending,
    Cancelled,
};

/// Converts PaymentStatus to proto-compatible UPPER_SNAKE_CASE string.
inline const char* payment_status_to_string(PaymentStatus s) {
    switch (s) {
    case PaymentStatus::Unspecified:   return "PAYMENT_STATUS_UNSPECIFIED";
    case PaymentStatus::Pending:       return "PAYMENT_STATUS_PENDING";
    case PaymentStatus::Processing:    return "PAYMENT_STATUS_PROCESSING";
    case PaymentStatus::Submitted:     return "PAYMENT_STATUS_SUBMITTED";
    case PaymentStatus::Included:      return "PAYMENT_STATUS_INCLUDED";
    case PaymentStatus::Confirmed:     return "PAYMENT_STATUS_CONFIRMED";
    case PaymentStatus::Finalized:     return "PAYMENT_STATUS_FINALIZED";
    case PaymentStatus::Failed:        return "PAYMENT_STATUS_FAILED";
    case PaymentStatus::RefundPending: return "PAYMENT_STATUS_REFUND_PENDING";
    case PaymentStatus::Cancelled:     return "PAYMENT_STATUS_CANCELLED";
    }
    return "PAYMENT_STATUS_UNSPECIFIED";
}

/// Parses proto-compatible UPPER_SNAKE_CASE string to PaymentStatus.
inline PaymentStatus payment_status_from_string(const std::string& s) {
    if (s == "PAYMENT_STATUS_PENDING")        return PaymentStatus::Pending;
    if (s == "PAYMENT_STATUS_PROCESSING")     return PaymentStatus::Processing;
    if (s == "PAYMENT_STATUS_SUBMITTED")      return PaymentStatus::Submitted;
    if (s == "PAYMENT_STATUS_INCLUDED")       return PaymentStatus::Included;
    if (s == "PAYMENT_STATUS_CONFIRMED")      return PaymentStatus::Confirmed;
    if (s == "PAYMENT_STATUS_FINALIZED")      return PaymentStatus::Finalized;
    if (s == "PAYMENT_STATUS_FAILED")         return PaymentStatus::Failed;
    if (s == "PAYMENT_STATUS_REFUND_PENDING") return PaymentStatus::RefundPending;
    if (s == "PAYMENT_STATUS_CANCELLED")      return PaymentStatus::Cancelled;
    return PaymentStatus::Unspecified;
}

/// Wallet type values matching proto WalletType enum (gRPC-Gateway UPPER_SNAKE_CASE).
enum class WalletType {
    Unspecified,
    Setto,
    Metamask,
    OKX,
    Coinbase,
    Phantom,
};

/// Converts WalletType to proto-compatible UPPER_SNAKE_CASE string.
inline const char* wallet_type_to_string(WalletType w) {
    switch (w) {
    case WalletType::Unspecified: return "WALLET_TYPE_UNSPECIFIED";
    case WalletType::Setto:      return "WALLET_TYPE_SETTO";
    case WalletType::Metamask:   return "WALLET_TYPE_METAMASK";
    case WalletType::OKX:        return "WALLET_TYPE_OKX";
    case WalletType::Coinbase:   return "WALLET_TYPE_COINBASE";
    case WalletType::Phantom:    return "WALLET_TYPE_PHANTOM";
    }
    return "WALLET_TYPE_UNSPECIFIED";
}

/// Parses proto-compatible UPPER_SNAKE_CASE string to WalletType.
inline WalletType wallet_type_from_string(const std::string& s) {
    if (s == "WALLET_TYPE_SETTO")    return WalletType::Setto;
    if (s == "WALLET_TYPE_METAMASK") return WalletType::Metamask;
    if (s == "WALLET_TYPE_OKX")      return WalletType::OKX;
    if (s == "WALLET_TYPE_COINBASE") return WalletType::Coinbase;
    if (s == "WALLET_TYPE_PHANTOM")  return WalletType::Phantom;
    return WalletType::Unspecified;
}

/// Request for initiating a payment.
struct InitiatePaymentRequest {
    std::string merchant_id;
    std::string amount;
    int32_t chain_id;
    std::string contract_address;
    WalletType wallet_type;
    std::string setto_user_id;
};

/// Response from payment initiation.
/// Fields match proto InitiatePaymentResponse (gRPC-Gateway snake_case JSON).
struct InitiatePaymentResponse {
    std::string payment_id;
    std::string merchant_id;
    std::string pool_address;
    std::string amount;
    int32_t chain_id;
    std::string contract_address;
    int64_t expires_at;
    int64_t created_at;
    std::string fee_amount;
    std::optional<std::string> merchant_address;
    std::optional<int32_t> decimals;
};

/// Payment information matching proto GetExternalPaymentStatusResponse (gRPC-Gateway snake_case JSON).
struct PaymentInfo {
    std::string payment_id;
    PaymentStatus status;
    std::optional<std::string> tx_hash;
    std::string amount;
    std::string currency;
    int64_t created_at;
    std::optional<int64_t> completed_at;
    std::optional<uint32_t> decimals;
    std::optional<std::string> sender_address;
    std::optional<std::string> pool_address;
    std::optional<int32_t> chain_id;
    std::optional<std::string> contract_address;

    bool is_complete() const { return status == PaymentStatus::Included; }
    bool is_failed() const { return status == PaymentStatus::Failed || status == PaymentStatus::Cancelled; }
    bool is_pending() const {
        return status == PaymentStatus::Pending
            || status == PaymentStatus::Processing
            || status == PaymentStatus::Submitted;
    }
};

// JWT Claims

struct Claims {
    std::string user_id;
    std::string email;
    bool email_verified;
    int64_t issued_at;
    int64_t expires_at;
};

} // namespace setto
