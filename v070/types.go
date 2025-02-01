package v070

import (
	"encoding/json"
	"math/big"

	"github.com/e4coder/tb-sdk/adapters"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

const DUMMY_SIGNATURE = "0xfffffffffffffffffffffffffffffff0000000000000000000000000000000007aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1c"

func (op *PackedUserOperation) ToUnpackedUserOperation() UnpackedUserOperation {
	// decode initCode    ->   initCode[:20] = facotry, initCode[20:] = factoryData
	var factory *common.Address = nil
	var factoryData []byte = []byte{}
	if len(op.InitCode) != 0 {
		if len(op.InitCode) >= 20 {
			_factory := common.BytesToAddress(op.InitCode[:20])
			factory = &_factory
			factoryData = op.InitCode[20:]
		}
	}

	// decode accountGasLimits   -> accountGasLimits[:16] = verificationGasLimit, accountGasLimits[16:] = callGasLimit
	var verificationGasLimit *big.Int = nil
	var callGasLimit *big.Int = nil
	if len(op.AccountGasLimits) == 32 {
		verificationGasLimit = big.NewInt(0).SetBytes(op.AccountGasLimits[:16])
		callGasLimit = big.NewInt(0).SetBytes(op.AccountGasLimits[16:])
	}

	// decode GasFees    ->   GasFees[:16] = maxPriorityFeePerGas, GasFees[16:] = maxFeePerGas
	var maxPriorityFeePerGas *big.Int = nil
	var maxFeePerGas *big.Int = nil
	if len(op.GasFees) == 32 {
		maxPriorityFeePerGas = big.NewInt(0).SetBytes(op.GasFees[:16])
		maxFeePerGas = big.NewInt(0).SetBytes(op.GasFees[16:])
	}

	// decode PaymasterAndData
	// PaymasterAndData[:20] = paymaster
	// PaymasterAndData[20:36] = validationGasLimit
	// PaymasterAndData[36:52] = postOpLimit
	// PaymasterAndData[52:] = paymasterData
	var paymaster *common.Address = nil
	var paymasterVerificationGasLimit *big.Int = nil
	var paymasterPostOpGasLimit *big.Int = nil
	var paymasterData []byte = []byte{}
	if len(op.PaymasterAndData) != 0 {
		if len(op.PaymasterAndData) >= 52 {
			_paymaster := common.BytesToAddress(op.PaymasterAndData[:20])
			paymaster = &_paymaster
			paymasterVerificationGasLimit = big.NewInt(0).SetBytes(op.PaymasterAndData[20:36])
			paymasterPostOpGasLimit = big.NewInt(0).SetBytes(op.PaymasterAndData[36:52])
			paymasterData = op.PaymasterAndData[52:]
		}
	}

	return UnpackedUserOperation{
		Sender:                        op.Sender,
		Nonce:                         op.Nonce,
		CallData:                      op.CallData,
		CallGasLimit:                  callGasLimit,
		VerificationGasLimit:          verificationGasLimit,
		PreVerificationGas:            op.PreVerificationGas,
		MaxPriorityFeePerGas:          maxPriorityFeePerGas,
		MaxFeePerGas:                  maxFeePerGas,
		Factory:                       factory,
		FactoryData:                   factoryData,
		Paymaster:                     paymaster,
		PaymasterVerificationGasLimit: paymasterVerificationGasLimit,
		PaymasterPostOpGasLimit:       paymasterPostOpGasLimit,
		PaymasterData:                 paymasterData,
		SignatureType:                 "0x00",
		Signature:                     DUMMY_SIGNATURE,
	}
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

func (pOp *PackedUserOperationString) ToPackedUserOperation() (PackedUserOperation, error) {
	sender := common.HexToAddress(pOp.Sender)
	nonce, err := hexutil.DecodeBig(pOp.Nonce)
	initCode, err := hexutil.Decode(pOp.InitCode)
	callData, err := hexutil.Decode(pOp.CallData)
	accountGasLimits, err := hexutil.Decode(pOp.AccountGasLimits)
	preVerificationGas, err := hexutil.DecodeBig(pOp.PreVerificationGas)
	paymasterAndData, err := hexutil.Decode(pOp.PaymasterAndData)
	signature := pOp.Signature

	if err != nil {
		return PackedUserOperation{}, err
	}

	return PackedUserOperation{
		Sender:             &sender,
		Nonce:              nonce,
		InitCode:           initCode,
		CallData:           callData,
		AccountGasLimits:   accountGasLimits,
		PreVerificationGas: preVerificationGas,
		PaymasterAndData:   paymasterAndData,
		Signature:          signature,
	}, nil
}
