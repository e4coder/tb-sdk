package tbsdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type RpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

type RpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
	Id      int         `json:"id"`
}

type RpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserOp struct {
	Sender                        string
	Nonce                         big.Int
	Factory                       common.Address
	FactoryData                   []byte
	CallData                      []byte
	CallGasLimit                  big.Int
	VerificationGasLimit          big.Int
	PreVerificationGas            big.Int
	MaxFeePerGas                  big.Int
	MaxPriorityFeePerGas          big.Int
	Paymaster                     common.Address
	PaymasterVerificationGasLimit big.Int
	PaymasterPostOpGasLimit       big.Int
	PaymasterData                 []byte
	Signature                     string
}

type PackedUserOp struct {
	Sender               string `json:"sender,omitempty"`
	Nonce                string `json:"nonce,omitempty"`
	InitCode             string `json:"initCode,omitempty"`
	CallData             string `json:"callData,omitempty"`
	Signature            string `json:"signature,omitempty"`
	CallGasLimit         string `json:"callGasLimit,omitempty"`
	VerificationGasLimit string `json:"verificationGasLimit,omitempty"`
	PreVerificationGas   string `json:"preVerificationGas,omitempty"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	PaymasterAndData     string `json:"paymasterAndData,omitempty"`
}
