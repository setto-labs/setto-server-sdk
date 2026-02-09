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
    std::optional<std::string> email;
    std::optional<std::string> photo_url;
    std::optional<std::string> payout_svm_address;
    std::optional<std::string> fee_rate;
    std::optional<std::string> one_time_token;
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

/// Update merchant wallet addresses. Requires OTT from frontend SDK.
struct UpdateMerchantRequest {
    std::string merchant_id;
    std::string one_time_token;
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

/// Update merchant display info only (name, photo_url). No OTT required.
struct UpdateMerchantProfileRequest {
    std::string merchant_id;
    std::optional<std::string> name;
    std::optional<std::string> photo_url;
};

struct UpdateMerchantProfileResponse {
    std::string merchant_id;
    std::string name;
    std::string photo_url;
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

enum class PaymentStatus {
    Pending,
    Submitted,
    Included,
    Failed,
    Cancelled,
};

struct PaymentInfo {
    std::string payment_id;
    PaymentStatus status;
    std::optional<std::string> tx_hash;
    std::string amount;
    std::string currency;
    int64_t created_at;
    std::optional<int64_t> completed_at;

    bool is_complete() const { return status == PaymentStatus::Included; }
    bool is_failed() const { return status == PaymentStatus::Failed || status == PaymentStatus::Cancelled; }
    bool is_pending() const { return status == PaymentStatus::Pending || status == PaymentStatus::Submitted; }
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
