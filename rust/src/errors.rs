use thiserror::Error;

#[derive(Debug, Error)]
pub enum SettoError {
    #[error("setto: system error: {code}")]
    System { code: String, http_status: u16 },

    #[error("setto: payment error: {code}")]
    Payment { code: String, http_status: u16 },

    #[error("setto: validation error: {message}")]
    Validation { message: String, http_status: u16 },

    #[error("setto: HTTP {status}")]
    Http { status: u16, body: String },

    #[error("setto: network error: {0}")]
    Network(#[from] reqwest::Error),
}

pub struct SystemError;

impl SystemError {
    pub const OK: &str = "SYSTEM_OK";
    pub const INTERNAL: &str = "SYSTEM_INTERNAL";
    pub const RPC_FAILED: &str = "SYSTEM_RPC_FAILED";
    pub const RATE_LIMITED: &str = "SYSTEM_RATE_LIMITED";
}

pub struct PaymentErrorCode;

impl PaymentErrorCode {
    pub const OK: &str = "PAYMENT_OK";
    pub const NOT_FOUND: &str = "PAYMENT_NOT_FOUND";
    pub const MERCHANT_NOT_FOUND: &str = "PAYMENT_MERCHANT_NOT_FOUND";
    pub const OTT_REQUIRED: &str = "PAYMENT_OTT_REQUIRED";
    pub const OTT_INVALID: &str = "PAYMENT_OTT_INVALID";
    pub const OTT_EXPIRED: &str = "PAYMENT_OTT_EXPIRED";
    pub const OTT_ALREADY_USED: &str = "PAYMENT_OTT_ALREADY_USED";
    pub const STORE_LIMIT_EXCEEDED: &str = "PAYMENT_STORE_LIMIT_EXCEEDED";
}
