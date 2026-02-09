# Setto Server SDK for Go

Production-ready Go SDK for the Setto Wallet Server REST API.

## Requirements

- Go 1.22 or later

## Installation

```bash
go get github.com/setto-labs/setto-server-sdk/go
```

## Quick Start

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
        APIKey:      "sk_partner.your_key",
        Environment: setto.Production,
    })
    if err != nil {
        log.Fatal(err)
    }

    merchant, err := client.CreateMerchant(context.Background(), &setto.CreateMerchantRequest{
        Email:            "merchant@example.com",
        Name:             "My Store",
        PayoutEVMAddress: "0x1234...abcd",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Merchant ID: %s\n", merchant.MerchantID)
}
```

## Client Configuration

### Basic

```go
client, err := setto.NewClient(setto.Config{
    APIKey:      "sk_partner....",   // Required. Must start with "sk_partner."
    Environment: setto.Production,   // Production or Development
})
```

### With Options

```go
client, err := setto.NewClient(
    setto.Config{
        APIKey:      "sk_partner....",
        Environment: setto.Development,
    },
    setto.WithTimeout(10 * time.Second),     // Default: 30s
    setto.WithHTTPClient(customHTTPClient),   // Custom *http.Client
    setto.WithBaseURL("https://custom.url"),  // Override base URL
)
```

### Environments

| Environment | Base URL | HTTPS Required |
|-------------|----------|----------------|
| `setto.Production` | `https://wallet.settopay.com` | Yes |
| `setto.Development` | `https://dev-wallet.settopay.com` | No |

Production environment enforces HTTPS. API key format (`sk_partner.` prefix) is always validated.

---

## API Reference

### Merchant

#### CreateMerchant

Creates a new merchant account in the Setto Wallet Server. The server automatically looks up or creates a user by email, then creates a merchant linked to that user.

```go
resp, err := client.CreateMerchant(ctx, &setto.CreateMerchantRequest{
    Email:            "merchant@example.com",       // Required
    Name:             "My Store",                   // Required
    PayoutEVMAddress: "0xabc...",                   // EVM payout address
    PayoutSVMAddress: "So1ana...",                  // Solana payout address (optional)
})
// resp.MerchantID — the created merchant's ID
```

**Request:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Email` | `string` | Yes | Merchant owner's email |
| `Name` | `string` | Yes | Merchant display name |
| `PayoutEVMAddress` | `string` | No | EVM chain payout address |
| `PayoutSVMAddress` | `string` | No | Solana payout address |

**Response:**

| Field | Type | Description |
|-------|------|-------------|
| `MerchantID` | `string` | Created merchant ID |

#### GetMerchant

Retrieves merchant details by ID.

```go
resp, err := client.GetMerchant(ctx, "merchant_id_here")
// resp.MerchantID, resp.PayoutEVMAddress, resp.PayoutSVMAddress
```

**Response:**

| Field | Type | Description |
|-------|------|-------------|
| `MerchantID` | `string` | Merchant ID |
| `PayoutEVMAddress` | `string` | EVM payout address |
| `PayoutSVMAddress` | `string` | Solana payout address |

#### UpdateMerchant

Updates merchant payout addresses. Requires a One-Time Token (OTT) for security.

```go
resp, err := client.UpdateMerchant(ctx, &setto.UpdateMerchantRequest{
    MerchantID:       "merchant_id",
    OneTimeToken:     "ott_token",      // Required — OTT from Setto
    PayoutEVMAddress: "0xnew...",
    PayoutSVMAddress: "NewSo1...",
})
```

**Request:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `MerchantID` | `string` | Yes | Target merchant ID |
| `OneTimeToken` | `string` | Yes | One-Time Token for authorization |
| `PayoutEVMAddress` | `string` | No | New EVM payout address |
| `PayoutSVMAddress` | `string` | No | New Solana payout address |

---

### Partner

#### GetVerificationStatus

Checks whether a user has completed phone verification.

```go
status, err := client.GetVerificationStatus(ctx, "setto_user_id")
if status.IsPhoneVerified {
    // User is phone-verified
}
```

**Response:**

| Field | Type | Description |
|-------|------|-------------|
| `IsPhoneVerified` | `bool` | Whether the user's phone is verified |

#### ExchangeAccountLinkToken

Exchanges a short-lived link token (from the Setto app) for the user's account information.

```go
info, err := client.ExchangeAccountLinkToken(ctx, "link_token_from_app")
// info.UserID   — Setto user ID
// info.Email    — User's email
```

**Response:**

| Field | Type | Description |
|-------|------|-------------|
| `UserID` | `string` | Setto user ID |
| `Email` | `string` | User's email address |

---

### Payment

#### GetPaymentStatus

Queries the real-time status of a payment.

```go
payment, err := client.GetPaymentStatus(ctx, "payment_id")
fmt.Printf("Status: %s\n", payment.Status)
```

**Response (`PaymentInfo`):**

| Field | Type | Description |
|-------|------|-------------|
| `PaymentID` | `string` | Payment ID |
| `MerchantID` | `string` | Merchant ID |
| `Status` | `PaymentStatus` | Current status |
| `Amount` | `string` | Payment amount |
| `Currency` | `string` | Currency code |
| `TxHash` | `string` | On-chain transaction hash (if submitted) |
| `CreatedAt` | `int64` | Creation timestamp (Unix) |
| `UpdatedAt` | `int64` | Last update timestamp (Unix) |

**PaymentStatus values:**

| Constant | Value | Description |
|----------|-------|-------------|
| `PaymentStatusPending` | `"pending"` | Payment initiated, awaiting processing |
| `PaymentStatusSubmitted` | `"submitted"` | Transaction submitted to blockchain |
| `PaymentStatusIncluded` | `"included"` | Transaction confirmed on-chain |
| `PaymentStatusFailed` | `"failed"` | Payment failed |
| `PaymentStatusCancelled` | `"cancelled"` | Payment cancelled |

**Helper methods on `PaymentInfo`:**

```go
payment.IsPaymentComplete() // true if status is "included"
payment.IsPaymentFailed()   // true if status is "failed" or "cancelled"
payment.IsPaymentPending()  // true if status is "pending" or "submitted"
```

---

### JWT Verification

Verify Setto-issued ID tokens using JWKS (RS256 with automatic key caching).

#### Standalone Verifier

```go
verifier := setto.NewVerifier(
    "https://wallet.settopay.com/.well-known/jwks.json",
    "https://wallet.settopay.com",
)

claims, err := verifier.VerifyIDToken(ctx, idTokenString)
fmt.Printf("User: %s, Email: %s\n", claims.UserID, claims.Email)
```

#### Client-based Verifier

When you already have a `Client`, create a verifier that auto-configures the JWKS URL:

```go
verifier := client.NewVerifier()
claims, err := verifier.VerifyIDToken(ctx, idTokenString)
```

#### VerifyIDTokenRequireEmail

Same as `VerifyIDToken`, but also validates that the token contains a verified email:

```go
claims, err := verifier.VerifyIDTokenRequireEmail(ctx, idTokenString)
// Returns ErrEmailNotVerified if email_verified is false or email is empty
```

**Claims:**

| Field | Type | Description |
|-------|------|-------------|
| `UserID` | `string` | Setto user ID (`sub` claim) |
| `Email` | `string` | User's email |
| `EmailVerified` | `bool` | Whether the email is verified |
| `IssuedAt` | `time.Time` | Token issued at |
| `ExpiresAt` | `time.Time` | Token expiration |

**JWKS behavior:**
- Keys are lazily fetched on first verification call
- Cached for 15 minutes, then refreshed automatically
- If a token has an unknown `kid`, a one-time re-fetch is attempted before failing

---

## Error Handling

### WalletError

Errors from the Setto Wallet Server are returned as `*WalletError`:

```go
resp, err := client.CreateMerchant(ctx, req)
if err != nil {
    if walletErr, ok := setto.IsWalletError(err); ok {
        fmt.Printf("Code: %s\n", walletErr.Code)
        fmt.Printf("Message: %s\n", walletErr.Message)
        fmt.Printf("HTTP Status: %d\n", walletErr.HTTPStatus)

        // Check specific error categories
        if walletErr.PaymentError != "" {
            fmt.Printf("Payment error: %s\n", walletErr.PaymentError)
        }
        return
    }
    // Network or other error
    log.Fatal(err)
}
```

**WalletError fields:**

| Field | Type | Description |
|-------|------|-------------|
| `SystemError` | `string` | System-level error code |
| `PaymentError` | `string` | Payment-related error code |
| `ValidationError` | `string` | Validation error code |
| `Code` | `string` | Primary error code (first non-empty of above) |
| `Message` | `string` | Human-readable error message |
| `HTTPStatus` | `int` | HTTP status code from the server |

### NetworkError

Network-level errors (DNS, timeout, connection refused) are wrapped as `*NetworkError`:

```go
if netErr, ok := err.(*setto.NetworkError); ok {
    fmt.Printf("Network error: %s\n", netErr.Err)
}
```

### Error Code Constants

```go
// System errors
setto.SystemOK
setto.SystemInternal
setto.SystemRPCFailed

// Payment errors
setto.PaymentNotFound
setto.PaymentOTTRequired
setto.PaymentOTTExpired
setto.PaymentOTTAlreadyUsed
setto.PaymentMerchantNotFound
setto.PaymentStoreLimitExceeded

// Validation errors
setto.ValidationInvalidRequest
setto.ValidationInvalidFormat
setto.ValidationInvalidID

// JWT errors
setto.ErrTokenInvalid
setto.ErrTokenExpired
setto.ErrIssuerMismatch
setto.ErrKeyNotFound
setto.ErrEmailNotVerified
```

---

## Full Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    setto "github.com/setto-labs/setto-server-sdk/go"
)

func main() {
    ctx := context.Background()

    // Initialize client
    client, err := setto.NewClient(setto.Config{
        APIKey:      "sk_partner.your_key",
        Environment: setto.Production,
    })
    if err != nil {
        log.Fatal(err)
    }

    // 1. Create merchant
    merchant, err := client.CreateMerchant(ctx, &setto.CreateMerchantRequest{
        Email:            "store@example.com",
        Name:             "Example Store",
        PayoutEVMAddress: "0x742d35Cc6634C0532925a3b844Bc9e7595f2bD18",
    })
    if err != nil {
        handleError(err)
        return
    }
    fmt.Printf("Created merchant: %s\n", merchant.MerchantID)

    // 2. Check verification status
    status, err := client.GetVerificationStatus(ctx, "user_id")
    if err != nil {
        handleError(err)
        return
    }
    fmt.Printf("Phone verified: %v\n", status.IsPhoneVerified)

    // 3. Check payment status
    payment, err := client.GetPaymentStatus(ctx, "payment_id")
    if err != nil {
        handleError(err)
        return
    }
    if payment.IsPaymentComplete() {
        fmt.Printf("Payment complete! TxHash: %s\n", payment.TxHash)
    }

    // 4. Verify JWT
    verifier := client.NewVerifier()
    claims, err := verifier.VerifyIDToken(ctx, "eyJhbGci...")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Authenticated user: %s (%s)\n", claims.UserID, claims.Email)
}

func handleError(err error) {
    if walletErr, ok := setto.IsWalletError(err); ok {
        fmt.Printf("API error [%s]: %s\n", walletErr.Code, walletErr.Message)
        return
    }
    if netErr, ok := err.(*setto.NetworkError); ok {
        fmt.Printf("Network error: %v\n", netErr.Err)
        return
    }
    log.Fatal(err)
}
```

## License

[MIT](../LICENSE)
