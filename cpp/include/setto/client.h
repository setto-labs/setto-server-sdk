#pragma once

#include "setto/errors.h"
#include "setto/types.h"

#include <memory>
#include <string>

namespace setto {

class Client {
public:
    explicit Client(const Config& config);
    ~Client();

    Client(const Client&) = delete;
    Client& operator=(const Client&) = delete;
    Client(Client&&) noexcept;
    Client& operator=(Client&&) noexcept;

    CreateMerchantResponse create_merchant(const CreateMerchantRequest& req);
    GetMerchantResponse get_merchant(const std::string& merchant_id);
    UpdateMerchantResponse update_merchant(const UpdateMerchantRequest& req);
    VerificationStatus get_verification_status(const std::string& user_id);
    AccountLinkInfo exchange_account_link_token(const std::string& link_token);
    PaymentInfo get_payment_status(const std::string& payment_id);

private:
    struct Impl;
    std::unique_ptr<Impl> impl_;
};

} // namespace setto
