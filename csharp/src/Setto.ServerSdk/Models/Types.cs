namespace Setto.ServerSdk.Models;

using System.Text.Json.Serialization;

public enum Environment
{
    Production,
    Development,
}

public sealed class SettoConfig
{
    public required string ApiKey { get; init; }
    public required Environment Environment { get; init; }
    public int? TimeoutMs { get; init; }
    public string? BaseUrl { get; init; }
}

// Merchant types

public sealed class CreateMerchantRequest
{
    public required string Name { get; init; }
    public required string PayoutEvmAddress { get; init; }
    public string? PhotoUrl { get; init; }
    public string? PayoutSvmAddress { get; init; }
    public string? FeeRate { get; init; }
}

public sealed class CreateMerchantResponse
{
    [JsonPropertyName("merchant_id")]
    public required string MerchantId { get; init; }
}

public sealed class GetMerchantResponse
{
    [JsonPropertyName("merchant_id")]
    public required string MerchantId { get; init; }
    [JsonPropertyName("name")]
    public required string Name { get; init; }
    [JsonPropertyName("photo_url")]
    public required string PhotoUrl { get; init; }
    [JsonPropertyName("payout_evm_address")]
    public required string PayoutEvmAddress { get; init; }
    [JsonPropertyName("payout_svm_address")]
    public required string PayoutSvmAddress { get; init; }
}

/// <summary>Update merchant wallet addresses.</summary>
public sealed class UpdateMerchantRequest
{
    public required string MerchantId { get; init; }
    public string? Name { get; init; }
    public string? PhotoUrl { get; init; }
    public string? PayoutEvmAddress { get; init; }
    public string? PayoutSvmAddress { get; init; }
}

public sealed class UpdateMerchantResponse
{
    [JsonPropertyName("merchant_id")]
    public required string MerchantId { get; init; }
    [JsonPropertyName("name")]
    public required string Name { get; init; }
    [JsonPropertyName("photo_url")]
    public required string PhotoUrl { get; init; }
    [JsonPropertyName("payout_evm_address")]
    public required string PayoutEvmAddress { get; init; }
    [JsonPropertyName("payout_svm_address")]
    public required string PayoutSvmAddress { get; init; }
}

// Verification types

public sealed class VerificationStatus
{
    [JsonPropertyName("is_phone_verified")]
    public required bool IsPhoneVerified { get; init; }
    [JsonPropertyName("verified_at")]
    public required long VerifiedAt { get; init; }
}

public sealed class AccountLinkInfo
{
    [JsonPropertyName("user_id")]
    public required string UserId { get; init; }
    [JsonPropertyName("email")]
    public required string Email { get; init; }
    [JsonPropertyName("is_phone_verified")]
    public required bool IsPhoneVerified { get; init; }
}

// Payment types

/// <summary>Payment status values matching proto PaymentStatus enum (gRPC-Gateway UPPER_SNAKE_CASE).</summary>
[JsonConverter(typeof(JsonStringEnumConverter<PaymentStatus>))]
public enum PaymentStatus
{
    [JsonPropertyName("PAYMENT_STATUS_UNSPECIFIED")] Unspecified,
    [JsonPropertyName("PAYMENT_STATUS_PENDING")] Pending,
    [JsonPropertyName("PAYMENT_STATUS_PROCESSING")] Processing,
    [JsonPropertyName("PAYMENT_STATUS_SUBMITTED")] Submitted,
    [JsonPropertyName("PAYMENT_STATUS_INCLUDED")] Included,
    [JsonPropertyName("PAYMENT_STATUS_CONFIRMED")] Confirmed,
    [JsonPropertyName("PAYMENT_STATUS_FINALIZED")] Finalized,
    [JsonPropertyName("PAYMENT_STATUS_FAILED")] Failed,
    [JsonPropertyName("PAYMENT_STATUS_REFUND_PENDING")] RefundPending,
    [JsonPropertyName("PAYMENT_STATUS_CANCELLED")] Cancelled,
}

/// <summary>Wallet type values matching proto WalletType enum (gRPC-Gateway UPPER_SNAKE_CASE).</summary>
[JsonConverter(typeof(JsonStringEnumConverter<WalletType>))]
public enum WalletType
{
    [JsonPropertyName("WALLET_TYPE_UNSPECIFIED")] Unspecified,
    [JsonPropertyName("WALLET_TYPE_SETTO")] Setto,
    [JsonPropertyName("WALLET_TYPE_METAMASK")] Metamask,
    [JsonPropertyName("WALLET_TYPE_OKX")] OKX,
    [JsonPropertyName("WALLET_TYPE_COINBASE")] Coinbase,
    [JsonPropertyName("WALLET_TYPE_PHANTOM")] Phantom,
}

/// <summary>Request for initiating a payment.</summary>
public sealed class InitiatePaymentRequest
{
    [JsonPropertyName("merchant_id")]
    public required string MerchantId { get; init; }
    [JsonPropertyName("amount")]
    public required string Amount { get; init; }
    [JsonPropertyName("chain_id")]
    public required int ChainId { get; init; }
    [JsonPropertyName("contract_address")]
    public required string ContractAddress { get; init; }
    [JsonPropertyName("wallet_type")]
    public required WalletType WalletType { get; init; }
    [JsonPropertyName("setto_user_id")]
    public required string SettoUserId { get; init; }
}

/// <summary>Response from payment initiation. Fields match proto InitiatePaymentResponse (gRPC-Gateway snake_case JSON).</summary>
public sealed class InitiatePaymentResponse
{
    [JsonPropertyName("payment_id")]
    public required string PaymentId { get; init; }
    [JsonPropertyName("merchant_id")]
    public required string MerchantId { get; init; }
    [JsonPropertyName("pool_address")]
    public required string PoolAddress { get; init; }
    [JsonPropertyName("amount")]
    public required string Amount { get; init; }
    [JsonPropertyName("chain_id")]
    public required int ChainId { get; init; }
    [JsonPropertyName("contract_address")]
    public required string ContractAddress { get; init; }
    [JsonPropertyName("expires_at")]
    public required long ExpiresAt { get; init; }
    [JsonPropertyName("created_at")]
    public required long CreatedAt { get; init; }
    [JsonPropertyName("fee_amount")]
    public required string FeeAmount { get; init; }
    [JsonPropertyName("merchant_address")]
    public string? MerchantAddress { get; init; }
    [JsonPropertyName("decimals")]
    public int? Decimals { get; init; }
}

/// <summary>Payment information matching proto GetExternalPaymentStatusResponse (gRPC-Gateway snake_case JSON).</summary>
public sealed class PaymentInfo
{
    [JsonPropertyName("payment_id")]
    public required string PaymentId { get; init; }
    [JsonPropertyName("status")]
    public required PaymentStatus Status { get; init; }
    [JsonPropertyName("tx_hash")]
    public string? TxHash { get; init; }
    [JsonPropertyName("amount")]
    public required string Amount { get; init; }
    [JsonPropertyName("currency")]
    public required string Currency { get; init; }
    [JsonPropertyName("created_at")]
    public required long CreatedAt { get; init; }
    [JsonPropertyName("completed_at")]
    public long? CompletedAt { get; init; }
    [JsonPropertyName("decimals")]
    public uint? Decimals { get; init; }
    [JsonPropertyName("sender_address")]
    public string? SenderAddress { get; init; }
    [JsonPropertyName("pool_address")]
    public string? PoolAddress { get; init; }
    [JsonPropertyName("chain_id")]
    public int? ChainId { get; init; }
    [JsonPropertyName("contract_address")]
    public string? ContractAddress { get; init; }

    public bool IsComplete => Status == PaymentStatus.Included;
    public bool IsFailed => Status is PaymentStatus.Failed or PaymentStatus.Cancelled;
    public bool IsPending => Status is PaymentStatus.Pending or PaymentStatus.Processing or PaymentStatus.Submitted;
}

// JWT Claims

public sealed class Claims
{
    public required string UserId { get; init; }
    public required string Email { get; init; }
    public required bool EmailVerified { get; init; }
    public required DateTimeOffset IssuedAt { get; init; }
    public required DateTimeOffset ExpiresAt { get; init; }
}
