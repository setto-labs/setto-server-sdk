use crate::errors::SettoError;
use crate::types::*;

const PRODUCTION_URL: &str = "https://wallet.settopay.com";
const DEVELOPMENT_URL: &str = "https://dev-wallet.settopay.com";

pub struct Client {
    api_key: String,
    base_url: String,
    http: reqwest::Client,
}

impl Client {
    pub fn new(config: Config) -> Result<Self, SettoError> {
        if !config.api_key.starts_with("sk_partner.") {
            panic!("setto: API key must start with 'sk_partner.'");
        }

        let base_url = config.base_url.unwrap_or_else(|| {
            match config.environment {
                Environment::Production => PRODUCTION_URL.to_string(),
                Environment::Development => DEVELOPMENT_URL.to_string(),
            }
        });

        let timeout = std::time::Duration::from_millis(config.timeout_ms.unwrap_or(30_000));
        let http = reqwest::Client::builder()
            .timeout(timeout)
            .build()
            .map_err(SettoError::Network)?;

        Ok(Self {
            api_key: config.api_key,
            base_url,
            http,
        })
    }

    pub async fn create_merchant(
        &self,
        _req: &CreateMerchantRequest,
    ) -> Result<CreateMerchantResponse, SettoError> {
        todo!("not implemented")
    }

    pub async fn get_merchant(
        &self,
        _merchant_id: &str,
    ) -> Result<GetMerchantResponse, SettoError> {
        todo!("not implemented")
    }

    pub async fn update_merchant(
        &self,
        _req: &UpdateMerchantRequest,
    ) -> Result<UpdateMerchantResponse, SettoError> {
        todo!("not implemented")
    }

    pub async fn get_verification_status(
        &self,
        _user_id: &str,
    ) -> Result<VerificationStatus, SettoError> {
        todo!("not implemented")
    }

    pub async fn exchange_account_link_token(
        &self,
        _link_token: &str,
    ) -> Result<AccountLinkInfo, SettoError> {
        todo!("not implemented")
    }

    pub async fn get_payment_status(
        &self,
        _payment_id: &str,
    ) -> Result<PaymentInfo, SettoError> {
        todo!("not implemented")
    }

    #[allow(dead_code)]
    fn api_key(&self) -> &str {
        &self.api_key
    }

    #[allow(dead_code)]
    fn base_url(&self) -> &str {
        &self.base_url
    }
}
