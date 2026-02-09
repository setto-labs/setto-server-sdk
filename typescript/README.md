# Setto Server SDK for TypeScript

> **Status: Skeleton** — Type definitions and API structure are defined, but the implementation is not yet complete. Contributions are welcome.

## Requirements

- Node.js 18 or later
- TypeScript 5.0 or later

## Installation (Coming Soon)

```bash
npm install @setto/server-sdk
```

## Planned Usage

```typescript
import { SettoClient } from "@setto/server-sdk";

const client = new SettoClient({
  apiKey: "sk_partner_your_key",
  environment: "production",
});

// Create a merchant
const merchant = await client.createMerchant({
  email: "merchant@example.com",
  name: "My Store",
  payoutEvmAddress: "0x1234...abcd",
});
console.log(`Merchant ID: ${merchant.merchantId}`);

// Check verification status
const status = await client.getVerificationStatus("setto_user_id");
console.log(`Phone verified: ${status.isPhoneVerified}`);

// Get payment status
const payment = await client.getPaymentStatus("payment_id");
console.log(`Payment status: ${payment.status}`);
```

## API Surface

### Client

```typescript
const client = new SettoClient({
  apiKey: string;         // Required. Must start with "sk_partner_"
  environment?: string;   // "production" (default) or "development"
  baseUrl?: string;       // Override base URL
  timeout?: number;       // Request timeout in ms (default: 30000)
});
```

### Methods

| Method | Description |
|--------|-------------|
| `createMerchant(req)` | Create a new merchant |
| `getMerchant(merchantId)` | Get merchant details |
| `updateMerchant(req)` | Update merchant payout addresses (requires OTT) |
| `getVerificationStatus(userId)` | Check user's phone verification status |
| `exchangeAccountLinkToken(linkToken)` | Exchange link token for account info |
| `getPaymentStatus(paymentId)` | Get real-time payment status |

### Types

All request/response types are fully defined in [`src/types.ts`](src/types.ts).

## Project Structure

```
typescript/
├── package.json
├── tsconfig.json
└── src/
    ├── index.ts      # Public exports
    ├── client.ts     # SettoClient class (skeleton)
    ├── types.ts      # Request/response type definitions
    └── errors.ts     # Error types
```

## Contributing

This SDK skeleton defines the complete API surface matching the [Go SDK](../go/). To implement:

1. Add HTTP request logic in `client.ts` (using `fetch`)
2. Add JWKS JWT verification (using `jose` library)
3. Add error parsing to match `WalletError` format
4. Add tests

See the [Go SDK source](../go/) for reference implementation details.

## License

[MIT](../LICENSE)
