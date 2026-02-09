# Setto Server SDK

Official server-side SDK for integrating with the [Setto Wallet Server](https://settopay.com). This SDK enables partner servers to manage merchants, verify users, process payments, and authenticate Setto users via JWT.

## Language Support

| Language | Status | Directory | Install |
|----------|--------|-----------|---------|
| **Go** | Production-ready | [`go/`](go/) | `go get github.com/setto-labs/setto-server-sdk/go` |
| **TypeScript** | Skeleton | [`typescript/`](typescript/) | Coming soon |
| **Rust** | Skeleton | [`rust/`](rust/) | Coming soon |
| **C#** | Skeleton | [`csharp/`](csharp/) | Coming soon |
| **C++** | Skeleton | [`cpp/`](cpp/) | Coming soon |

> **Note:** Only the Go SDK is production-ready. Other language SDKs provide the type definitions and API structure, but are not yet implemented. Contributions are welcome.

## Features

- **Merchant Management** — Create, retrieve, and update merchant accounts with payout wallet addresses
- **User Verification** — Check phone verification status for partner-linked users
- **Account Linking** — Exchange link tokens to associate Setto accounts with partner systems
- **Payment Status** — Query real-time payment status with convenient helper methods
- **JWT Verification** — Verify Setto-issued ID tokens using JWKS (RS256, auto-caching)

## API Overview

All SDKs communicate with the Setto Wallet Server via REST (HTTPS + JSON). Authentication uses an API key with the `X-API-Key` header.

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/merchant` | Create a new merchant |
| `GET` | `/api/merchant/{id}` | Get merchant details |
| `PUT` | `/api/merchant/{id}` | Update merchant (requires OTT) |
| `GET` | `/api/partner/user/{id}/verification` | Get user verification status |
| `POST` | `/api/partner/exchange-link-token` | Exchange account link token |
| `GET` | `/api/external/payment/{id}` | Get payment status |
| `GET` | `/.well-known/jwks.json` | JWKS endpoint (for JWT verification) |

### Environments

| Environment | Base URL |
|-------------|----------|
| Production | `https://wallet.settopay.com` |
| Development | `https://dev-wallet.settopay.com` |

## Authentication

All API requests require a partner API key. Keys follow the format `sk_partner_*`.

```
X-API-Key: sk_partner_your_key_here
```

Contact [Setto](https://settopay.com) to obtain your partner API key.

## Quick Start (Go)

```go
package main

import (
    "context"
    "fmt"
    "log"

    setto "github.com/setto-labs/setto-server-sdk/go"
)

func main() {
    client, err := setto.NewClient(setto.Config{
        APIKey:      "sk_partner_your_key",
        Environment: setto.Production,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create a merchant
    merchant, err := client.CreateMerchant(context.Background(), &setto.CreateMerchantRequest{
        Email:            "merchant@example.com",
        Name:             "My Store",
        PayoutEVMAddress: "0x1234...abcd",
    })
    if err != nil {
        if walletErr, ok := setto.IsWalletError(err); ok {
            fmt.Printf("Wallet error: %s - %s\n", walletErr.Code, walletErr.Message)
            return
        }
        log.Fatal(err)
    }

    fmt.Printf("Merchant created: %s\n", merchant.MerchantID)
}
```

See the [Go SDK README](go/README.md) for the full API reference.

## Error Handling

All SDKs use structured error types. Errors from the Setto Wallet Server include a machine-readable error code and a human-readable message.

**Error categories:**

| Category | Example Codes |
|----------|---------------|
| System | `SYSTEM_INTERNAL`, `SYSTEM_RPC_FAILED` |
| Payment | `PAYMENT_NOT_FOUND`, `PAYMENT_OTT_EXPIRED`, `PAYMENT_STORE_LIMIT_EXCEEDED` |
| Validation | `VALIDATION_INVALID_REQUEST`, `VALIDATION_INVALID_FORMAT` |

## License

[MIT](LICENSE)
