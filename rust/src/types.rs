use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Environment {
    Production,
    Development,
}

#[derive(Debug, Clone)]
pub struct Config {
    pub api_key: String,
    pub environment: Environment,
    pub timeout_ms: Option<u64>,
    pub base_url: Option<String>,
}

// Merchant types

#[derive(Debug, Serialize)]
pub struct CreateMerchantRequest {
    pub name: String,
    pub payout_evm_address: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub email: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub photo_url: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub payout_svm_address: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub fee_rate: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub one_time_token: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct CreateMerchantResponse {
    pub merchant_id: String,
}

#[derive(Debug, Deserialize)]
pub struct GetMerchantResponse {
    pub merchant_id: String,
    pub name: String,
    pub photo_url: String,
    pub payout_evm_address: String,
    pub payout_svm_address: String,
}

/// Update merchant wallet addresses. Requires OTT from frontend SDK.
#[derive(Debug, Serialize)]
pub struct UpdateMerchantRequest {
    pub merchant_id: String,
    pub one_time_token: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub photo_url: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub payout_evm_address: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub payout_svm_address: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct UpdateMerchantResponse {
    pub merchant_id: String,
    pub name: String,
    pub photo_url: String,
    pub payout_evm_address: String,
    pub payout_svm_address: String,
}

/// Update merchant display info only (name, photo_url). No OTT required.
#[derive(Debug, Serialize)]
pub struct UpdateMerchantProfileRequest {
    pub merchant_id: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub photo_url: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct UpdateMerchantProfileResponse {
    pub merchant_id: String,
    pub name: String,
    pub photo_url: String,
}

// Verification types

#[derive(Debug, Deserialize)]
pub struct VerificationStatus {
    pub is_phone_verified: bool,
    pub verified_at: i64,
}

#[derive(Debug, Deserialize)]
pub struct AccountLinkInfo {
    pub user_id: String,
    pub email: String,
    pub is_phone_verified: bool,
}

// Payment types

#[derive(Debug, Clone, Copy, PartialEq, Eq, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum PaymentStatus {
    Pending,
    Submitted,
    Included,
    Failed,
    Cancelled,
}

#[derive(Debug, Deserialize)]
pub struct PaymentInfo {
    pub payment_id: String,
    pub status: PaymentStatus,
    pub tx_hash: Option<String>,
    pub amount: String,
    pub currency: String,
    pub created_at: i64,
    pub completed_at: Option<i64>,
}

impl PaymentInfo {
    pub fn is_complete(&self) -> bool {
        self.status == PaymentStatus::Included
    }

    pub fn is_failed(&self) -> bool {
        matches!(self.status, PaymentStatus::Failed | PaymentStatus::Cancelled)
    }

    pub fn is_pending(&self) -> bool {
        matches!(self.status, PaymentStatus::Pending | PaymentStatus::Submitted)
    }
}

// JWT Claims

#[derive(Debug, Deserialize)]
pub struct Claims {
    pub user_id: String,
    pub email: String,
    pub email_verified: bool,
    pub iat: i64,
    pub exp: i64,
}
