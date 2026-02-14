namespace Setto.ServerSdk;

using Setto.ServerSdk.Models;

public sealed class SettoClient : IDisposable
{
    private const string ProductionUrl = "https://wallet.settopay.com";
    private const string DevelopmentUrl = "https://dev-wallet.settopay.com";

    private readonly string _apiKey;
    private readonly string _baseUrl;
    private readonly HttpClient _httpClient;

    public SettoClient(SettoConfig config)
    {
        if (string.IsNullOrEmpty(config.ApiKey) || !config.ApiKey.StartsWith("sk_setto."))
            throw new ArgumentException("setto: API key must start with 'sk_setto.'");

        _apiKey = config.ApiKey;
        _baseUrl = config.BaseUrl ?? (config.Environment == Environment.Production ? ProductionUrl : DevelopmentUrl);

        _httpClient = new HttpClient
        {
            Timeout = TimeSpan.FromMilliseconds(config.TimeoutMs ?? 30_000),
        };
    }

    public Task<CreateMerchantResponse> CreateMerchantAsync(CreateMerchantRequest request, CancellationToken ct = default)
        => throw new NotImplementedException();

    public Task<GetMerchantResponse> GetMerchantAsync(string merchantId, CancellationToken ct = default)
        => throw new NotImplementedException();

    public Task<UpdateMerchantResponse> UpdateMerchantAsync(UpdateMerchantRequest request, CancellationToken ct = default)
        => throw new NotImplementedException();

    public Task<VerificationStatus> GetVerificationStatusAsync(string userId, CancellationToken ct = default)
        => throw new NotImplementedException();

    public Task<AccountLinkInfo> ExchangeAccountLinkTokenAsync(string linkToken, CancellationToken ct = default)
        => throw new NotImplementedException();

    public Task<PaymentInfo> GetPaymentStatusAsync(string paymentId, CancellationToken ct = default)
        => throw new NotImplementedException();

    public void Dispose()
    {
        _httpClient.Dispose();
    }
}
