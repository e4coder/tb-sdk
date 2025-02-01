package v070

import (
	"encoding/json"
	"math/big"

	"github.com/e4coder/tb-sdk/adapters"
	"github.com/ethereum/go-ethereum/common"
)

type UnpackedUserOperation struct {
	Sender                        *common.Address
	Nonce                         *big.Int
	CallData                      []byte
	CallGasLimit                  *big.Int
	VerificationGasLimit          *big.Int
	PreVerificationGas            *big.Int
	MaxPriorityFeePerGas          *big.Int
	MaxFeePerGas                  *big.Int
	Factory                       *common.Address
	FactoryData                   []byte
	Paymaster                     *common.Address
	PaymasterVerificationGasLimit *big.Int
	PaymasterPostOpGasLimit       *big.Int
	PaymasterData                 []byte
	Signature                     string
	SignatureType                 string
}

type RpcUnpackedUserOperation struct {
	Sender                        string  `json:"sender"`
	Nonce                         string  `json:"nonce"`
	CallData                      string  `json:"callData"`
	CallGasLimit                  string  `json:"callGasLimit,omitempty"`
	VerificationGasLimit          string  `json:"verificationGasLimit,omitempty"`
	PreVerificationGas            string  `json:"preVerificationGas,omitempty"`
	MaxPriorityFeePerGas          string  `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas                  string  `json:"maxFeePerGas,omitempty"`
	Factory                       *string `json:"factory,omitempty"`
	FactoryData                   *string `json:"factoryData"`
	Paymaster                     *string `json:"paymaster,omitempty"`
	PaymasterVerificationGasLimit *string `json:"paymasterVerificationGasLimit,omitempty"`
	PaymasterPostOpGasLimit       *string `json:"paymasterPostOpGasLimit,omitempty"`
	PaymasterData                 *string `json:"paymasterData,omitempty"`
	Signature                     string  `json:"signature"`
}

type PackedUserOperation struct {
	Sender             *common.Address `json:"address"`
	Nonce              *big.Int        `json:"nonce"`
	InitCode           []byte          `json:"initCode"`
	CallData           []byte          `json:"callData"`
	AccountGasLimits   []byte          `json:"accountGasLimits"`
	PreVerificationGas *big.Int        `json:"preVerificationGas"`
	GasFees            []byte          `json:"gasFees"`
	PaymasterAndData   []byte          `json:"paymasterAndData"`
	Signature          string          `json:"signature"`
}

func (op *PackedUserOperation) ToPackedUserOperationString() PackedUserOperationString {
	pHex := PackedUserOperationString{
		Sender:             adapters.Adapt(adapters.ADDRESS_ADAPTER, op.Sender),
		Nonce:              adapters.Adapt(adapters.BIG_INT_ADAPTER, op.Nonce),
		InitCode:           adapters.Adapt(adapters.PACKED_DATA_ADAPTER, op.InitCode),
		CallData:           adapters.Adapt(adapters.PACKED_DATA_ADAPTER, op.CallData),
		AccountGasLimits:   adapters.Adapt(adapters.PACKED_DATA_ADAPTER, op.AccountGasLimits),
		PreVerificationGas: adapters.Adapt(adapters.BIG_INT_ADAPTER, op.PreVerificationGas),
		GasFees:            adapters.Adapt(adapters.PACKED_DATA_ADAPTER, op.GasFees),
		PaymasterAndData:   adapters.Adapt(adapters.PACKED_DATA_ADAPTER, op.PaymasterAndData),
		Signature:          op.Signature,
	}

	return pHex
}

func (op PackedUserOperation) String() string {
	pHex := op.ToPackedUserOperationString()

	jOp, err := json.MarshalIndent(pHex, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(jOp)
}

type PackedUserOperationString struct {
	Sender             string `json:"sender"`
	Nonce              string `json:"nonce"`
	InitCode           string `json:"initCode"`
	CallData           string `json:"callData"`
	AccountGasLimits   string `json:"accountGasLimits"`
	PreVerificationGas string `json:"preVerificationGas"`
	GasFees            string `json:"gasFees"`
	PaymasterAndData   string `json:"paymasterAndData"`
	Signature          string `json:"signature"`
}
