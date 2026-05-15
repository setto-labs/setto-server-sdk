export enum Environment {
  Production = "production",
  Development = "development",
}

export interface SettoConfig {
  apiKey: string;
  environment: Environment;
  timeout?: number;
  baseUrl?: string;
}

/** Payment status values matching proto PaymentStatus enum (gRPC-Gateway UPPER_SNAKE_CASE). */
export type PaymentStatus =
  | "PAYMENT_STATUS_UNSPECIFIED"
  | "PAYMENT_STATUS_PENDING"
  | "PAYMENT_STATUS_PROCESSING"
  | "PAYMENT_STATUS_SUBMITTED"
  | "PAYMENT_STATUS_INCLUDED"
  | "PAYMENT_STATUS_CONFIRMED"
  | "PAYMENT_STATUS_FINALIZED"
  | "PAYMENT_STATUS_FAILED"
  | "PAYMENT_STATUS_REFUND_PENDING"
  | "PAYMENT_STATUS_CANCELLED";

/** Wallet type values matching proto WalletType enum (gRPC-Gateway UPPER_SNAKE_CASE). */
export type WalletType =
  | "WALLET_TYPE_UNSPECIFIED"
  | "WALLET_TYPE_SETTO"
  | "WALLET_TYPE_METAMASK"
  | "WALLET_TYPE_OKX"
  | "WALLET_TYPE_COINBASE"
  | "WALLET_TYPE_PHANTOM";

// Merchant types
export interface CreateMerchantRequest {
  name: string;
  photoUrl?: string;
  payoutEvmAddress: string;
  payoutSvmAddress?: string;
  feeRate?: string;
}

export interface CreateMerchantResponse {
  merchantId: string;
}

export interface GetMerchantResponse {
  merchantId: string;
  name: string;
  photoUrl: string;
  payoutEvmAddress: string;
  payoutSvmAddress: string;
}

/** Update merchant wallet addresses. */
export interface UpdateMerchantRequest {
  merchantId: string;
  name?: string;
  photoUrl?: string;
  payoutEvmAddress?: string;
  payoutSvmAddress?: string;
}

export interface UpdateMerchantResponse {
  merchantId: string;
  name: string;
  photoUrl: string;
  payoutEvmAddress: string;
  payoutSvmAddress: string;
}

// Verification types
export interface VerificationStatus {
  isPhoneVerified: boolean;
  verifiedAt: number;
}

export interface AccountLinkInfo {
  userId: string;
  email: string;
  isPhoneVerified: boolean;
}

// Payment types (fields match proto GetExternalPaymentStatusResponse, gRPC-Gateway snake_case JSON)
export interface PaymentInfo {
  payment_id: string;
  status: PaymentStatus;
  tx_hash?: string;
  amount: string;
  currency: string;
  created_at: number;
  completed_at?: number;
  decimals?: number;
  sender_address?: string;
  pool_address?: string;
  chain_id?: number;
  contract_address?: string;
}

// JWT Claims
export interface Claims {
  userId: string;
  email: string;
  emailVerified: boolean;
  issuedAt: Date;
  expiresAt: Date;
}
