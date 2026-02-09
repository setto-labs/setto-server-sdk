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

export type PaymentStatus =
  | "pending"
  | "submitted"
  | "included"
  | "failed"
  | "cancelled";

// Merchant types
export interface CreateMerchantRequest {
  email?: string;
  name: string;
  photoUrl?: string;
  payoutEvmAddress: string;
  payoutSvmAddress?: string;
  feeRate?: string;
  oneTimeToken?: string;
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

/** Update merchant wallet addresses. Requires OTT from frontend SDK. */
export interface UpdateMerchantRequest {
  merchantId: string;
  oneTimeToken: string;
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

/** Update merchant display info only (name, photoUrl). No OTT required. */
export interface UpdateMerchantProfileRequest {
  merchantId: string;
  name?: string;
  photoUrl?: string;
}

export interface UpdateMerchantProfileResponse {
  merchantId: string;
  name: string;
  photoUrl: string;
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

// Payment types
export interface PaymentInfo {
  paymentId: string;
  status: PaymentStatus;
  txHash?: string;
  amount: string;
  currency: string;
  createdAt: number;
  completedAt?: number;
}

// JWT Claims
export interface Claims {
  userId: string;
  email: string;
  emailVerified: boolean;
  issuedAt: Date;
  expiresAt: Date;
}
