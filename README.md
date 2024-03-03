# TB-SDK: THE BUNDLER SDK

This golang SDK is designed to aid in the integration with the
ERC-4773 Account Abstractions Bundler component. The TB-SDK provides a
comprehensive suite of functions to interact directly with a bundler
instance, streamlining the process of preparing, executing, and managing
user operations on the blockchain.

## Supported Methods

### `eth_` Namespace

Methods defined by the [ERC-4337 spec](https://github.com/eth-infinitism/account-abstraction/blob/develop/erc/ERCS/erc-4337.md#rpc-methods-eth-namespace).

| Method | Supported |
| ------ | :-----------: |
| `eth_chainId` | ✅ |
| `eth_supportedEntryPoints` | ✅ |
| `eth_estimateUserOperationGas` | ✅ |
| `eth_sendUserOperation` | 🚧 |
| `eth_getUserOperationByHash` | 🚧 |
| `eth_getUserOperationReceipt` | 🚧 |

### `debug_` Namespace

Method defined by the [ERC-4337 spec](https://github.com/eth-infinitism/account-abstraction/blob/develop/erc/ERCS/erc-4337.md#rpc-methods-debug-namespace). Used only for debugging/testing and should be disabled on production APIs.

| Method | Supported |
| ------ | :-----------: |
| `debug_bundler_clearState` | 🚧 |
| `debug_bundler_dumpMempool` | 🚧 |
| `debug_bundler_sendBundleNow` | 🚧 |
| `debug_bundler_setBundlingMode` | 🚧 |
| `debug_bundler_setReputation` | 🚧 |
| `debug_bundler_dumpReputation` | 🚧 |
| `debug_bundler_addUserOps` | 🚧 | |

## Sample Usage
```go
package main

import (
    tbsdk "github.com/e4coder/tb-sdk"
)

func main() {
    bundler := tbsdk.NewBundler()
	err := bundler.Init("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	userOp := tbsdk.NewOperationBuilder().
		Nonce(big.NewInt(123)).
		Sender(sender).
		CallData(data).
		FactoryAndData(&factoryAddress, factoryData).
		Signature(signature).
		Build()

    // RPC -> eth_estimateUserOperationGas
	rpcResponse, err := bundler.Eth_estimateUserOperationGas(userOp)
	if err != nil {
		panic(err)
	}

	fmt.Println(rpcResponse.Result)
}

```

## Contributing

Feel free to open pull requests and issues, the project is in active development any and all help is very much appreciated