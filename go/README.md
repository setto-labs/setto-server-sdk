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
        APIKey:      "sk_setto.your_key",
        Environment: setto.Production,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Check verification status
    status, err := client.GetVerificationStatus(context.Background(), "user_id")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Phone verified: %v\n", status.IsPhoneVerified)
}
```

## Client Configuration

### Basic

```go
client, err := setto.NewClient(setto.Config{
    APIKey:      "sk_setto....",   // Required. Must start with "sk_setto."
    Environment: setto.Production,   // Production or Development
})
```

### With Options

```go
client, err := setto.NewClient(
    setto.Config{
        APIKey:      "sk_setto....",
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

Production environment enforces HTTPS. API key format (`sk_setto.` prefix) is always validated.

---

## API Reference

### Integration

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
payment, err := client.GetPaymentStatus(ctx, "payment_id")
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
        APIKey:      "sk_setto.your_key",
        Environment: setto.Production,
    })
    if err != nil {
        log.Fatal(err)
    }

    // 1. Check verification status
    status, err := client.GetVerificationStatus(ctx, "user_id")
    if err != nil {
        handleError(err)
        return
    }
    fmt.Printf("Phone verified: %v\n", status.IsPhoneVerified)

    // 2. Check payment status
    payment, err := client.GetPaymentStatus(ctx, "payment_id")
    if err != nil {
        handleError(err)
        return
    }
    if payment.IsPaymentComplete() {
        fmt.Printf("Payment complete! TxHash: %s\n", payment.TxHash)
    }

    // 3. Verify JWT
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
