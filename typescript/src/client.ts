import type {
  SettoConfig,
  CreateMerchantRequest,
  CreateMerchantResponse,
  GetMerchantResponse,
  UpdateMerchantRequest,
  UpdateMerchantResponse,
  VerificationStatus,
  AccountLinkInfo,
  PaymentInfo,
} from "./types.js";

const PRODUCTION_URL = "https://wallet.settopay.com";
const DEVELOPMENT_URL = "https://dev-wallet.settopay.com";

export class SettoClient {
  private readonly apiKey: string;
  private readonly baseUrl: string;
  private readonly timeout: number;

  constructor(config: SettoConfig) {
    if (!config.apiKey?.startsWith("sk_partner_")) {
      throw new Error("setto: API key must start with 'sk_partner_'");
    }
    this.apiKey = config.apiKey;
    this.baseUrl =
      config.baseUrl ??
      (config.environment === "production" ? PRODUCTION_URL : DEVELOPMENT_URL);
    this.timeout = config.timeout ?? 30000;
  }

  async createMerchant(
    _req: CreateMerchantRequest,
  ): Promise<CreateMerchantResponse> {
    throw new Error("not implemented");
  }

  async getMerchant(_merchantId: string): Promise<GetMerchantResponse> {
    throw new Error("not implemented");
  }

  async updateMerchant(
    _req: UpdateMerchantRequest,
  ): Promise<UpdateMerchantResponse> {
    throw new Error("not implemented");
  }

  async getVerificationStatus(
    _userId: string,
  ): Promise<VerificationStatus> {
    throw new Error("not implemented");
  }

  async exchangeAccountLinkToken(
    _linkToken: string,
  ): Promise<AccountLinkInfo> {
    throw new Error("not implemented");
  }

  async getPaymentStatus(_paymentId: string): Promise<PaymentInfo> {
    throw new Error("not implemented");
  }
}
