# Setto Server SDK for C++

> **Status: Skeleton** — Type definitions and API structure are defined, but the implementation is not yet complete. Contributions are welcome.

## Requirements

- C++20 or later
- CMake 3.20+
- libcurl
- nlohmann/json (optional, recommended)

## Installation (Coming Soon)

### CMake FetchContent

```cmake
include(FetchContent)
FetchContent_Declare(
    setto-server-sdk
    GIT_REPOSITORY https://github.com/setto-labs/setto-server-sdk.git
    GIT_TAG go/v0.1.0
    SOURCE_SUBDIR cpp
)
FetchContent_MakeAvailable(setto-server-sdk)

target_link_libraries(your_target PRIVATE setto-server-sdk)
```

### Manual Build

```bash
cd cpp
mkdir build && cd build
cmake ..
make
```

## Planned Usage

```cpp
#include <setto/client.h>
#include <iostream>

int main() {
    setto::Config config;
    config.api_key = "sk_partner.your_key";
    config.environment = setto::Environment::Production;

    setto::Client client(config);

    // Create a merchant
    setto::CreateMerchantRequest req;
    req.email = "merchant@example.com";
    req.name = "My Store";
    req.payout_evm_address = "0x1234...abcd";

    auto merchant = client.create_merchant(req);
    std::cout << "Merchant ID: " << merchant.merchant_id << std::endl;

    // Check verification status
    auto status = client.get_verification_status("setto_user_id");
    std::cout << "Phone verified: " << status.is_phone_verified << std::endl;

    // Get payment status
    auto payment = client.get_payment_status("payment_id");
    std::cout << "Payment status: " << payment.status << std::endl;

    return 0;
}
```

## API Surface

### Client

```cpp
setto::Config config;
config.api_key = "sk_partner....";                       // Required
config.environment = setto::Environment::Production;     // Production or Development
config.base_url = "https://custom.url";                  // Optional override
config.timeout_ms = 30000;                               // Optional (default: 30000)

setto::Client client(config);
```

### Methods

| Method | Description |
|--------|-------------|
| `create_merchant(req)` | Create a new merchant |
| `get_merchant(merchant_id)` | Get merchant details |
| `update_merchant(req)` | Update merchant payout addresses (requires OTT) |
| `get_verification_status(user_id)` | Check user's phone verification status |
| `exchange_account_link_token(link_token)` | Exchange link token for account info |
| `get_payment_status(payment_id)` | Get real-time payment status |

### Error Handling

```cpp
#include <setto/errors.h>

try {
    auto merchant = client.create_merchant(req);
} catch (const setto::WalletError& e) {
    std::cerr << "API error [" << e.code() << "]: " << e.what() << std::endl;
} catch (const setto::NetworkError& e) {
    std::cerr << "Network error: " << e.what() << std::endl;
}
```

## Dependencies

| Library | Purpose |
|---------|---------|
| libcurl | HTTP client |
| nlohmann/json | JSON parsing (optional) |

## Project Structure

```
cpp/
├── CMakeLists.txt
├── include/setto/
│   ├── client.h      # Client class declaration
│   ├── types.h       # Request/response structs
│   └── errors.h      # Exception classes
└── src/
    ├── client.cpp     # Client implementation (skeleton)
    └── errors.cpp     # Error implementation
```

## Contributing

This SDK skeleton defines the complete API surface matching the [Go SDK](../go/). To implement:

1. Add HTTP request logic in `src/client.cpp` (using libcurl)
2. Add JSON parsing (using nlohmann/json or similar)
3. Add JWKS JWT verification
4. Add error parsing to match `WalletError` format
5. Add tests

See the [Go SDK source](../go/) for reference implementation details.

## License

[MIT](../LICENSE)
