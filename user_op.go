package tbsdk

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OperationBuilder struct {
	Operation UserOp
}

func NewOperationBuilder() *OperationBuilder {
	return &OperationBuilder{
		Operation: UserOp{},
	}
}

func (ob *OperationBuilder) Sender(sender *common.Address) *OperationBuilder {
	ob.Operation.Sender = sender
	return ob
}

func (ob *OperationBuilder) Nonce(nonce *big.Int) *OperationBuilder {
	ob.Operation.Nonce = nonce
	return ob
}

func (ob *OperationBuilder) Factory(factory *common.Address) *OperationBuilder {
	ob.Operation.Factory = factory
	return ob
}

func (ob *OperationBuilder) FactoryData(factoryCallData []byte) *OperationBuilder {
	ob.Operation.FactoryData = factoryCallData
	return ob
}

func (ob *OperationBuilder) FactoryAndData(factory *common.Address, factoryCallData []byte) *OperationBuilder {
	return ob.Factory(factory).FactoryData(factoryCallData)
}

func (ob *OperationBuilder) CallData(callData []byte) *OperationBuilder {
	ob.Operation.CallData = callData
	return ob
}

func (ob *OperationBuilder) CallGasLimit(callGasLimit *big.Int) *OperationBuilder {
	ob.Operation.CallGasLimit = callGasLimit
	return ob
}

func (ob *OperationBuilder) VerificationGasLimit(verificationGasLimit *big.Int) *OperationBuilder {
	ob.Operation.VerificationGasLimit = verificationGasLimit
	return ob
}

func (ob *OperationBuilder) PreVerificationGas(preVerificationGas *big.Int) *OperationBuilder {
	ob.Operation.PreVerificationGas = preVerificationGas
	return ob
}

func (ob *OperationBuilder) MaxFeePerGas(maxFeePerGas *big.Int) *OperationBuilder {
	ob.Operation.MaxFeePerGas = maxFeePerGas
	return ob
}

func (ob *OperationBuilder) MaxPriorityFeePerGas(maxPriorityFeePerGas *big.Int) *OperationBuilder {
	ob.Operation.MaxPriorityFeePerGas = maxPriorityFeePerGas
	return ob
}

func (ob *OperationBuilder) Paymaster(paymaster *common.Address) *OperationBuilder {
	ob.Operation.Paymaster = paymaster
	return ob
}

func (ob *OperationBuilder) PaymasterVerificationGasLimit(paymasterVerificationGasLimit *big.Int) *OperationBuilder {
	ob.Operation.PaymasterVerificationGasLimit = paymasterVerificationGasLimit
	return ob
}

func (ob *OperationBuilder) PaymasterPostOpGasLimit(paymasterPostOpGasLimit *big.Int) *OperationBuilder {
	ob.Operation.PaymasterPostOpGasLimit = paymasterPostOpGasLimit
	return ob
}

func (ob *OperationBuilder) PaymasterData(paymasterCallData []byte) *OperationBuilder {
	ob.Operation.PaymasterData = paymasterCallData
	return ob
}

func (ob *OperationBuilder) PaymasterAndData(paymaster *common.Address, paymasterCallData []byte) *OperationBuilder {
	return ob.Paymaster(paymaster).PaymasterData(paymasterCallData)
}

func (ob *OperationBuilder) Signature(sig string) *OperationBuilder {
	ob.Operation.Signature = sig
	return ob
}

func (ob *OperationBuilder) Adapt(adapter Adapter, val interface{}, args ...interface{}) string {
	return adapter(val, args)
}

func (ob *OperationBuilder) Build() *PackedUserOp {
	initCode := ob.Adapt(ADDRESS_PACKED_DATA_ADAPTER, ob.Operation.Factory, ob.Operation.FactoryData)
	paymasterAndData := ob.Adapt(ADDRESS_PACKED_DATA_ADAPTER, ob.Operation.Paymaster, ob.Operation.PaymasterData)

	packedOp := &PackedUserOp{
		Sender:               ob.Adapt(ADDRESS_ADAPTER, ob.Operation.Sender),
		Nonce:                ob.Adapt(BIG_INT_ADAPTER, ob.Operation.Nonce), // fixme
		InitCode:             initCode,
		CallData:             ob.Adapt(PACKED_DATA_ADAPTER, ob.Operation.CallData),
		CallGasLimit:         ob.Adapt(BIG_INT_ADAPTER, ob.Operation.CallGasLimit),
		VerificationGasLimit: ob.Adapt(BIG_INT_ADAPTER, ob.Operation.VerificationGasLimit),
		PreVerificationGas:   ob.Adapt(BIG_INT_ADAPTER, ob.Operation.PreVerificationGas),
		MaxFeePerGas:         ob.Adapt(BIG_INT_ADAPTER, ob.Operation.MaxFeePerGas),
		MaxPriorityFeePerGas: ob.Adapt(BIG_INT_ADAPTER, ob.Operation.MaxPriorityFeePerGas),
		PaymasterAndData:     paymasterAndData,
		Signature:            ob.Operation.Signature,
	}

	return packedOp
}

func (ob *OperationBuilder) BuildWithOperationHash(chainId *big.Int, entryPoint common.Address) (*PackedUserOp, []byte, error) {
	packedOp := ob.Build()

	uoData, err := GetUserOperationHash(packedOp, chainId.Int64(), entryPoint)
	if err != nil {
		return nil, nil, fmt.Errorf("err operationHash : %v", err)
	}

	return packedOp, uoData, nil
}

func (ob *OperationBuilder) BuildWithSignature(chainId *big.Int, entryPoint common.Address, privKey *ecdsa.PrivateKey) (*PackedUserOp, error) {
	if chainId == nil {
		return nil, fmt.Errorf("chainId <nil>")
	}

	if privKey == nil {
		return nil, fmt.Errorf("privKey <nil>")
	}

	packedOp, uoData, err := ob.BuildWithOperationHash(chainId, entryPoint)
	if err != nil {
		return nil, err
	}

	sig, err := SignDataWithEthereumPrivateKey(uoData, privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %v", err)
	}
	packedOp.Signature = "0x" + hex.EncodeToString(sig)

	return packedOp, nil
}

func SignUserOp() {

}
