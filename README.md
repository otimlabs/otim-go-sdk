# Otim Go SDK

Go SDK for interacting with the Otim orchestration platform. This SDK provides a client for building and submitting settlement orchestrations, with secure signing powered by Turnkey.

## Installation

```bash
go get github.com/otimlabs/otim-go-sdk
```

## Quick Start

```go
package main

import (
	"context"
	"log"

	"github.com/otimlabs/otim-go-sdk/client"
	"github.com/otimlabs/otim-go-sdk/signer"
)

func main() {
	ctx := context.Background()

	// Initialize the signer with your private key
	ethSigner, err := signer.NewEthSigner("private-key-hex-string")
	if err != nil {
		log.Fatal(err)
	}

	// Create the client (delegate address is fetched automatically)
	otimClient, err := client.NewClient(
		ethSigner,
		"https://api.otim.com",
		"your-api-key",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Build a settlement orchestration
	buildResp, err := otimClient.BuildSettlementOrchestration(ctx, &client.BuildSettlementOrchestrationRequest{
		Params: &client.SettlementParams{
			AcceptedTokens: map[client.ChainID][]common.Address{
				1: {common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")}, // USDC on Ethereum
			},
			SettlementChainId: 1, // Ethereum mainnet
			SettlementToken:   common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
			SettlementAmount:  hexutil.Big(*big.NewInt(1000000)), // 1 USDC (6 decimals)
			RecipientAddress:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Sign the orchestration (EIP-7702 authorization + EIP-712 instructions)
	newReq, err := otimClient.NewOrchestrationFromBuild(ctx, buildResp)
	if err != nil {
		log.Fatal(err)
	}

	// Submit the signed orchestration
	err = otimClient.NewOrchestration(ctx, newReq)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Orchestration Types

The SDK supports multiple orchestration types through the `OrchestrationParams` interface:

### Standard Settlement

Standard multi-chain settlement orchestration:

```go
import (
	"math/big"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/otimlabs/otim-go-sdk/client"
)

settlementAmount := big.NewInt(1000000) // 1 USDC (6 decimals)
usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

buildReq := &client.BuildSettlementOrchestrationRequest{
	Params: &client.SettlementParams{
		AcceptedTokens: map[client.ChainID][]common.Address{
			1:  {usdcAddress}, // Ethereum
			10: {common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607")}, // Optimism USDC
		},
		SettlementChainId: 1,
		SettlementToken:   usdcAddress,
		SettlementAmount:  hexutil.Big(*settlementAmount),
		RecipientAddress:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0"),
	},
}
```

### Vault Withdrawal Settlement

Vault withdrawal settlement orchestration for ERC4626 vaults:

```go
withdrawAmount := big.NewInt(5000000000) // 5000 tokens (assuming 6 decimals)

buildReq := &client.BuildSettlementOrchestrationRequest{
	Params: &client.VaultWithdrawSettlementParams{
		VaultAddress:         common.HexToAddress("0x1234..."), // ERC4626 vault address
		VaultUnderlyingToken: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
		VaultChainId:         1, // Ethereum
		SettlementChainId:    1,
		SettlementToken:      common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		RecipientAddress:     common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0"),
		WithdrawAmount:       hexutil.Big(*withdrawAmount),
	},
}
```

## Packages

### `client`

The main client package for interacting with the Otim API.

- `NewClient` - Creates a new Otim client
- `BuildSettlementOrchestration` - Builds a settlement orchestration request
- `NewOrchestrationFromBuild` - Signs the orchestration using Turnkey
- `NewOrchestration` - Submits the signed orchestration

### `signer`

Signing utilities with Turnkey integration.

- `EthSigner` - Implements the `Signer` interface using ECDSA keys
- `NewEthSigner` - Creates a signer from a hex-encoded private key

The signer supports observability through custom Turnkey SDK options:

```go
import (
	"net/http"

	"github.com/otimlabs/otim-go-sdk/signer"
	sdk "github.com/tkhq/go-sdk"
)

// Use a custom HTTP client with OpenTelemetry instrumentation
httpClient := &http.Client{
	Transport: otelhttp.NewTransport(http.DefaultTransport),
}

ethSigner, err := signer.NewEthSigner(
	privateKeyHex,
	sdk.WithHTTPClient(httpClient),
)
```

### `abi`

Generated Go bindings for Otim action contracts, including:

- `SweepAction`, `SweepERC20Action`, `SweepCCTPAction`
- `TransferAction`, `TransferERC20Action`, `TransferCCTPAction`
- `UniswapV3ExactInputAction`
- `DepositERC4626Action`, `WithdrawERC4626Action`
- And more...

## Configuration

The SDK requires the following configuration:

| Parameter             | Description                                        |
| --------------------- | -------------------------------------------------- |
| `apiUrl`              | Otim API endpoint URL                              |
| `apiKey`              | Your Otim API key                                  |
| `developerPrivateKey` | Hex-encoded private key for Turnkey authentication |

## Examples

For a complete integration example, see [`client/integration_test.go`](client/integration_test.go).

The integration test demonstrates the full orchestration flow:

1. Loading configuration from environment variables
2. Initializing the `EthSigner` with Turnkey credentials
3. Creating the Otim client
4. Building a settlement orchestration request
5. Signing the EIP-7702 authorization and EIP-712 instructions via Turnkey
6. Submitting the signed orchestration to the API

### Running the Integration Test

**WARNING**: if you run this integration test with your production credentials, all wallets will be deleted and orchestrations canceled due to test cleanup. Only run with test credentials.

```bash
export OTIM_API_URL="https://api.otim.com"
export OTIM_API_KEY="your-api-key"
export OTIM_PRIVATE_KEY="your-private-key-hex"

go test -v ./client -run TestSettlementOrchestrationIntegration
```

## License

This project is dual-licensed under the Apache 2.0 and MIT licenses. See [LICENSE-APACHE](LICENSE-APACHE) and [LICENSE-MIT](LICENSE-MIT) for details.
