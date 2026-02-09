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
    public string? Email { get; init; }
    public string? PhotoUrl { get; init; }
    public string? PayoutSvmAddress { get; init; }
    public string? FeeRate { get; init; }
    public string? OneTimeToken { get; init; }
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

/// <summary>Update merchant wallet addresses. Requires OTT from frontend SDK.</summary>
public sealed class UpdateMerchantRequest
{
    public required string MerchantId { get; init; }
    public required string OneTimeToken { get; init; }
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

/// <summary>Update merchant display info only (name, photo_url). No OTT required.</summary>
public sealed class UpdateMerchantProfileRequest
{
    public required string MerchantId { get; init; }
    public string? Name { get; init; }
    public string? PhotoUrl { get; init; }
}

public sealed class UpdateMerchantProfileResponse
{
    [JsonPropertyName("merchant_id")]
    public required string MerchantId { get; init; }
    [JsonPropertyName("name")]
    public required string Name { get; init; }
    [JsonPropertyName("photo_url")]
    public required string PhotoUrl { get; init; }
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

[JsonConverter(typeof(JsonStringEnumConverter))]
public enum PaymentStatus
{
    [JsonPropertyName("pending")] Pending,
    [JsonPropertyName("submitted")] Submitted,
    [JsonPropertyName("included")] Included,
    [JsonPropertyName("failed")] Failed,
    [JsonPropertyName("cancelled")] Cancelled,
}

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

    public bool IsComplete => Status == PaymentStatus.Included;
    public bool IsFailed => Status is PaymentStatus.Failed or PaymentStatus.Cancelled;
    public bool IsPending => Status is PaymentStatus.Pending or PaymentStatus.Submitted;
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
