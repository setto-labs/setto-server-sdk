#pragma once

#include <cstdint>
#include <optional>
#include <stdexcept>
#include <string>

namespace setto {

class SettoError : public std::runtime_error {
public:
    SettoError(
        std::optional<std::string> system_error,
        std::optional<std::string> payment_error,
        std::optional<std::string> validation_error,
        uint16_t http_status
    );

    const std::optional<std::string>& system_error() const { return system_error_; }
    const std::optional<std::string>& payment_error() const { return payment_error_; }
    const std::optional<std::string>& validation_error() const { return validation_error_; }
    uint16_t http_status() const { return http_status_; }
    std::string code() const;

private:
    std::optional<std::string> system_error_;
    std::optional<std::string> payment_error_;
    std::optional<std::string> validation_error_;
    uint16_t http_status_;
};

class NetworkError : public std::runtime_error {
public:
    explicit NetworkError(const std::string& message)
        : std::runtime_error("setto: network error: " + message) {}
};

namespace SystemError {
    inline constexpr const char* OK = "SYSTEM_OK";
    inline constexpr const char* INTERNAL = "SYSTEM_INTERNAL";
    inline constexpr const char* RPC_FAILED = "SYSTEM_RPC_FAILED";
    inline constexpr const char* RATE_LIMITED = "SYSTEM_RATE_LIMITED";
}

namespace PaymentErrorCode {
    inline constexpr const char* OK = "PAYMENT_OK";
    inline constexpr const char* NOT_FOUND = "PAYMENT_NOT_FOUND";
    inline constexpr const char* MERCHANT_NOT_FOUND = "PAYMENT_MERCHANT_NOT_FOUND";
    inline constexpr const char* OTT_REQUIRED = "PAYMENT_OTT_REQUIRED";
    inline constexpr const char* OTT_INVALID = "PAYMENT_OTT_INVALID";
    inline constexpr const char* OTT_EXPIRED = "PAYMENT_OTT_EXPIRED";
    inline constexpr const char* OTT_ALREADY_USED = "PAYMENT_OTT_ALREADY_USED";
    inline constexpr const char* STORE_LIMIT_EXCEEDED = "PAYMENT_STORE_LIMIT_EXCEEDED";
}

} // namespace setto
