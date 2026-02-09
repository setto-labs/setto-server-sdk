# Setto Server SDK for Rust

> **Status: Skeleton** — Type definitions and API structure are defined, but the implementation is not yet complete. Contributions are welcome.

## Requirements

- Rust 2021 edition or later
- Tokio async runtime

## Installation (Coming Soon)

```toml
# Cargo.toml
[dependencies]
setto-server-sdk = "0.1"
```

## Planned Usage

```rust
use setto_server_sdk::{Client, Config, Environment};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::new(Config {
        api_key: "sk_partner_your_key".to_string(),
        environment: Environment::Production,
        base_url: None,
        timeout_ms: None,
    })?;

    // Create a merchant
    let merchant = client.create_merchant(&CreateMerchantRequest {
        email: "merchant@example.com".to_string(),
        name: "My Store".to_string(),
        payout_evm_address: Some("0x1234...abcd".to_string()),
        payout_svm_address: None,
    }).await?;
    println!("Merchant ID: {}", merchant.merchant_id);

    // Check verification status
    let status = client.get_verification_status("setto_user_id").await?;
    println!("Phone verified: {}", status.is_phone_verified);

    // Get payment status
    let payment = client.get_payment_status("payment_id").await?;
    println!("Payment status: {:?}", payment.status);

    Ok(())
}
```

## API Surface

### Client

```rust
let client = Client::new(Config {
    api_key: String,                  // Required. Must start with "sk_partner_"
    environment: Environment,         // Production or Development
    base_url: Option<String>,         // Override base URL
    timeout_ms: Option<u64>,          // Request timeout (default: 30000)
})?;
```

### Methods

| Method | Description |
|--------|-------------|
| `create_merchant(&req)` | Create a new merchant |
| `get_merchant(merchant_id)` | Get merchant details |
| `update_merchant(&req)` | Update merchant payout addresses (requires OTT) |
| `get_verification_status(user_id)` | Check user's phone verification status |
| `exchange_account_link_token(link_token)` | Exchange link token for account info |
| `get_payment_status(payment_id)` | Get real-time payment status |

### Error Handling

```rust
use setto_server_sdk::errors::SettoError;

match client.create_merchant(&req).await {
    Ok(resp) => println!("Created: {}", resp.merchant_id),
    Err(SettoError::Wallet(e)) => println!("API error [{}]: {}", e.code, e.message),
    Err(SettoError::Network(e)) => println!("Network error: {}", e),
}
```

## Dependencies

| Crate | Purpose |
|-------|---------|
| `reqwest` | HTTP client |
| `serde` / `serde_json` | JSON serialization |
| `thiserror` | Error type derivation |
| `jsonwebtoken` | JWT verification |
| `tokio` | Async runtime |

## Project Structure

```
rust/
├── Cargo.toml
└── src/
    ├── lib.rs        # Public exports
    ├── client.rs     # Client struct (skeleton)
    ├── types.rs      # Request/response types with serde
    └── errors.rs     # Error types with thiserror
```

## Contributing

This SDK skeleton defines the complete API surface matching the [Go SDK](../go/). To implement:

1. Add HTTP request logic in `client.rs` (using `reqwest`)
2. Add JWKS JWT verification (using `jsonwebtoken`)
3. Add error parsing to match `SettoError::Wallet` format
4. Add tests

See the [Go SDK source](../go/) for reference implementation details.

## License

[MIT](../LICENSE)
