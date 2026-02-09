export class SettoError extends Error {
  constructor(
    public readonly systemError?: string,
    public readonly paymentError?: string,
    public readonly validationError?: string,
    public readonly httpStatus?: number,
  ) {
    super(`setto: ${systemError || paymentError || validationError || `HTTP ${httpStatus}`}`);
    this.name = "SettoError";
  }

  get code(): string {
    return this.systemError || this.paymentError || this.validationError || "";
  }
}

export function isSettoError(err: unknown): err is SettoError {
  return err instanceof SettoError;
}

export class NetworkError extends Error {
  constructor(
    message: string,
    public readonly cause?: Error,
  ) {
    super(`setto: network error: ${message}`);
    this.name = "NetworkError";
  }
}

// Error code constants
export const SystemError = {
  OK: "SYSTEM_OK",
  INTERNAL: "SYSTEM_INTERNAL",
  RPC_FAILED: "SYSTEM_RPC_FAILED",
  RATE_LIMITED: "SYSTEM_RATE_LIMITED",
} as const;

export const PaymentErrorCode = {
  OK: "PAYMENT_OK",
  NOT_FOUND: "PAYMENT_NOT_FOUND",
  MERCHANT_NOT_FOUND: "PAYMENT_MERCHANT_NOT_FOUND",
  OTT_REQUIRED: "PAYMENT_OTT_REQUIRED",
  OTT_INVALID: "PAYMENT_OTT_INVALID",
  OTT_EXPIRED: "PAYMENT_OTT_EXPIRED",
  OTT_ALREADY_USED: "PAYMENT_OTT_ALREADY_USED",
  STORE_LIMIT_EXCEEDED: "PAYMENT_STORE_LIMIT_EXCEEDED",
} as const;
