#include "setto/client.h"

#include <stdexcept>

namespace setto {

static const char* PRODUCTION_URL = "https://wallet.settopay.com";
static const char* DEVELOPMENT_URL = "https://dev-wallet.settopay.com";

struct Client::Impl {
    std::string api_key;
    std::string base_url;
    uint32_t timeout_ms;
};

Client::Client(const Config& config) : impl_(std::make_unique<Impl>()) {
    if (config.api_key.rfind("sk_partner.", 0) != 0) {
        throw std::invalid_argument("setto: API key must start with 'sk_partner.'");
    }

    impl_->api_key = config.api_key;
    impl_->base_url = config.base_url.value_or(
        config.environment == Environment::Production ? PRODUCTION_URL : DEVELOPMENT_URL
    );
    impl_->timeout_ms = config.timeout_ms;
}

Client::~Client() = default;
Client::Client(Client&&) noexcept = default;
Client& Client::operator=(Client&&) noexcept = default;

CreateMerchantResponse Client::create_merchant(const CreateMerchantRequest& /*req*/) {
    throw std::runtime_error("not implemented");
}

GetMerchantResponse Client::get_merchant(const std::string& /*merchant_id*/) {
    throw std::runtime_error("not implemented");
}

UpdateMerchantResponse Client::update_merchant(const UpdateMerchantRequest& /*req*/) {
    throw std::runtime_error("not implemented");
}

VerificationStatus Client::get_verification_status(const std::string& /*user_id*/) {
    throw std::runtime_error("not implemented");
}

AccountLinkInfo Client::exchange_account_link_token(const std::string& /*link_token*/) {
    throw std::runtime_error("not implemented");
}

PaymentInfo Client::get_payment_status(const std::string& /*payment_id*/) {
    throw std::runtime_error("not implemented");
}

} // namespace setto
