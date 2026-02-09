#include "setto/errors.h"

namespace setto {

SettoError::SettoError(
    std::optional<std::string> system_error,
    std::optional<std::string> payment_error,
    std::optional<std::string> validation_error,
    uint16_t http_status
)
    : std::runtime_error(
          "setto: " +
          (system_error.value_or(
               payment_error.value_or(
                   validation_error.value_or(
                       "HTTP " + std::to_string(http_status))))))
    , system_error_(std::move(system_error))
    , payment_error_(std::move(payment_error))
    , validation_error_(std::move(validation_error))
    , http_status_(http_status)
{
}

std::string SettoError::code() const {
    if (system_error_) return *system_error_;
    if (payment_error_) return *payment_error_;
    if (validation_error_) return *validation_error_;
    return "";
}

} // namespace setto
