namespace Setto.ServerSdk;

public class SettoException : Exception
{
    public string? SystemError { get; }
    public string? PaymentError { get; }
    public string? ValidationError { get; }
    public int? HttpStatus { get; }

    public SettoException(string? systemError = null, string? paymentError = null, string? validationError = null, int? httpStatus = null)
        : base($"setto: {systemError ?? paymentError ?? validationError ?? $"HTTP {httpStatus}"}")
    {
        SystemError = systemError;
        PaymentError = paymentError;
        ValidationError = validationError;
        HttpStatus = httpStatus;
    }

    public string Code => SystemError ?? PaymentError ?? ValidationError ?? "";
}

public class NetworkException : Exception
{
    public NetworkException(string message, Exception? innerException = null)
        : base($"setto: network error: {message}", innerException) { }
}

public static class SystemErrorCode
{
    public const string OK = "SYSTEM_OK";
    public const string Internal = "SYSTEM_INTERNAL";
    public const string RpcFailed = "SYSTEM_RPC_FAILED";
    public const string RateLimited = "SYSTEM_RATE_LIMITED";
}

public static class PaymentErrorCode
{
    public const string OK = "PAYMENT_OK";
    public const string NotFound = "PAYMENT_NOT_FOUND";
    public const string MerchantNotFound = "PAYMENT_MERCHANT_NOT_FOUND";
    public const string OttRequired = "PAYMENT_OTT_REQUIRED";
    public const string OttInvalid = "PAYMENT_OTT_INVALID";
    public const string OttExpired = "PAYMENT_OTT_EXPIRED";
    public const string OttAlreadyUsed = "PAYMENT_OTT_ALREADY_USED";
    public const string StoreLimitExceeded = "PAYMENT_STORE_LIMIT_EXCEEDED";
}
