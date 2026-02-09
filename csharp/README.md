# Setto Server SDK for C#

> **Status: Skeleton** — Type definitions and API structure are defined, but the implementation is not yet complete. Contributions are welcome.

## Requirements

- .NET 8.0 or later

## Installation (Coming Soon)

```bash
dotnet add package Setto.ServerSdk
```

## Planned Usage

```csharp
using Setto.ServerSdk;
using Setto.ServerSdk.Models;

var client = new SettoClient(new SettoConfig
{
    ApiKey = "sk_partner.your_key",
    Environment = Environment.Production,
});

// Create a merchant
var merchant = await client.CreateMerchantAsync(new CreateMerchantRequest
{
    Email = "merchant@example.com",
    Name = "My Store",
    PayoutEvmAddress = "0x1234...abcd",
});
Console.WriteLine($"Merchant ID: {merchant.MerchantId}");

// Check verification status
var status = await client.GetVerificationStatusAsync("setto_user_id");
Console.WriteLine($"Phone verified: {status.IsPhoneVerified}");

// Get payment status
var payment = await client.GetPaymentStatusAsync("payment_id");
Console.WriteLine($"Payment status: {payment.Status}");
```

## API Surface

### Client

```csharp
var client = new SettoClient(new SettoConfig
{
    ApiKey = "sk_partner....",           // Required. Must start with "sk_partner."
    Environment = Environment.Production, // Production or Development
    BaseUrl = null,                       // Override base URL (optional)
    TimeoutMs = 30_000,                  // Request timeout in ms (optional)
});
```

### Methods

| Method | Description |
|--------|-------------|
| `CreateMerchantAsync(req, ct)` | Create a new merchant |
| `GetMerchantAsync(merchantId, ct)` | Get merchant details |
| `UpdateMerchantAsync(req, ct)` | Update merchant payout addresses (requires OTT) |
| `GetVerificationStatusAsync(userId, ct)` | Check user's phone verification status |
| `ExchangeAccountLinkTokenAsync(linkToken, ct)` | Exchange link token for account info |
| `GetPaymentStatusAsync(paymentId, ct)` | Get real-time payment status |

All methods accept an optional `CancellationToken` parameter.

### Error Handling

```csharp
try
{
    var merchant = await client.CreateMerchantAsync(req);
}
catch (SettoWalletException ex)
{
    Console.WriteLine($"API error [{ex.Code}]: {ex.Message}");
    Console.WriteLine($"HTTP Status: {ex.HttpStatus}");
}
catch (SettoNetworkException ex)
{
    Console.WriteLine($"Network error: {ex.Message}");
}
```

## Dependencies

| Package | Purpose |
|---------|---------|
| `System.IdentityModel.Tokens.Jwt` | JWT verification |
| `Microsoft.IdentityModel.Protocols.OpenIdConnect` | OIDC/JWKS support |

## Project Structure

```
csharp/
├── Setto.ServerSdk.sln
└── src/Setto.ServerSdk/
    ├── Setto.ServerSdk.csproj
    ├── SettoClient.cs         # Client class (skeleton)
    ├── SettoError.cs          # Exception types
    └── Models/
        └── Types.cs           # Request/response models
```

## Contributing

This SDK skeleton defines the complete API surface matching the [Go SDK](../go/). To implement:

1. Add HTTP request logic in `SettoClient.cs` (using `HttpClient`)
2. Add JWKS JWT verification (using `System.IdentityModel.Tokens.Jwt`)
3. Add error parsing to match `SettoWalletException` format
4. Add tests

See the [Go SDK source](../go/) for reference implementation details.

## License

[MIT](../LICENSE)
