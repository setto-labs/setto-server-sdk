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
    pub photo_url: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub payout_svm_address: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub fee_rate: Option<String>,
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

/// Update merchant wallet addresses.
#[derive(Debug, Serialize)]
pub struct UpdateMerchantRequest {
    pub merchant_id: String,
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

/// Payment status values matching proto PaymentStatus enum (gRPC-Gateway UPPER_SNAKE_CASE).
#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
pub enum PaymentStatus {
    #[serde(rename = "PAYMENT_STATUS_UNSPECIFIED")]
    Unspecified,
    #[serde(rename = "PAYMENT_STATUS_PENDING")]
    Pending,
    #[serde(rename = "PAYMENT_STATUS_PROCESSING")]
    Processing,
    #[serde(rename = "PAYMENT_STATUS_SUBMITTED")]
    Submitted,
    #[serde(rename = "PAYMENT_STATUS_INCLUDED")]
    Included,
    #[serde(rename = "PAYMENT_STATUS_CONFIRMED")]
    Confirmed,
    #[serde(rename = "PAYMENT_STATUS_FINALIZED")]
    Finalized,
    #[serde(rename = "PAYMENT_STATUS_FAILED")]
    Failed,
    #[serde(rename = "PAYMENT_STATUS_REFUND_PENDING")]
    RefundPending,
    #[serde(rename = "PAYMENT_STATUS_CANCELLED")]
    Cancelled,
}

/// Wallet type values matching proto WalletType enum (gRPC-Gateway UPPER_SNAKE_CASE).
#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
pub enum WalletType {
    #[serde(rename = "WALLET_TYPE_UNSPECIFIED")]
    Unspecified,
    #[serde(rename = "WALLET_TYPE_SETTO")]
    Setto,
    #[serde(rename = "WALLET_TYPE_METAMASK")]
    Metamask,
    #[serde(rename = "WALLET_TYPE_OKX")]
    OKX,
    #[serde(rename = "WALLET_TYPE_COINBASE")]
    Coinbase,
    #[serde(rename = "WALLET_TYPE_PHANTOM")]
    Phantom,
}

/// Request for initiating a payment.
#[derive(Debug, Serialize)]
pub struct InitiatePaymentRequest {
    pub merchant_id: String,
    pub amount: String,
    pub chain_id: i32,
    pub contract_address: String,
    pub wallet_type: WalletType,
    pub setto_user_id: String,
}

/// Response from payment initiation.
/// Fields match proto InitiatePaymentResponse (gRPC-Gateway snake_case JSON).
#[derive(Debug, Deserialize)]
pub struct InitiatePaymentResponse {
    pub payment_id: String,
    pub merchant_id: String,
    pub pool_address: String,
    pub amount: String,
    pub chain_id: i32,
    pub contract_address: String,
    pub expires_at: i64,
    pub created_at: i64,
    pub fee_amount: String,
    #[serde(default)]
    pub merchant_address: Option<String>,
    #[serde(default)]
    pub decimals: Option<i32>,
}

/// Payment information matching proto GetExternalPaymentStatusResponse (gRPC-Gateway snake_case JSON).
#[derive(Debug, Deserialize)]
pub struct PaymentInfo {
    pub payment_id: String,
    pub status: PaymentStatus,
    #[serde(default)]
    pub tx_hash: Option<String>,
    pub amount: String,
    pub currency: String,
    pub created_at: i64,
    #[serde(default)]
    pub completed_at: Option<i64>,
    #[serde(default)]
    pub decimals: Option<u32>,
    #[serde(default)]
    pub sender_address: Option<String>,
    #[serde(default)]
    pub pool_address: Option<String>,
    #[serde(default)]
    pub chain_id: Option<i32>,
    #[serde(default)]
    pub contract_address: Option<String>,
}

impl PaymentInfo {
    pub fn is_complete(&self) -> bool {
        self.status == PaymentStatus::Included
    }

    pub fn is_failed(&self) -> bool {
        matches!(self.status, PaymentStatus::Failed | PaymentStatus::Cancelled)
    }

    pub fn is_pending(&self) -> bool {
        matches!(
            self.status,
            PaymentStatus::Pending | PaymentStatus::Processing | PaymentStatus::Submitted
        )
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
