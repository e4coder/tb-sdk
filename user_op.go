package tbsdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OperationBuilder struct {
	Operation UserOp
}

func NewOperationBuilder() *OperationBuilder {
	return &OperationBuilder{}
}

func (ob *OperationBuilder) Sender(sender common.Address) *OperationBuilder {
	ob.Operation.Sender = sender.Hex()
	return ob
}

func (ob *OperationBuilder) Nonce() *OperationBuilder {
	return ob
}

func (ob *OperationBuilder) Factory(factory common.Address) *OperationBuilder {
	ob.Operation.Factory = factory
	return ob
}

func (ob *OperationBuilder) FactoryData(factoryCallData []byte) *OperationBuilder {
	ob.Operation.FactoryData = factoryCallData
	return ob
}

func (ob *OperationBuilder) FactoryAndData(factory common.Address, factoryCallData []byte) *OperationBuilder {
	return ob.Factory(factory).FactoryData(factoryCallData)
}

func (ob *OperationBuilder) CallData(callData []byte) *OperationBuilder {
	ob.Operation.CallData = callData
	return ob
}

func (ob *OperationBuilder) CallGasLimit(callGasLimit big.Int) *OperationBuilder {
	ob.Operation.CallGasLimit = callGasLimit
	return ob
}

func (ob *OperationBuilder) VerificationGasLimit(verificationGasLimit big.Int) *OperationBuilder {
	ob.Operation.VerificationGasLimit = verificationGasLimit
	return ob
}

func (ob *OperationBuilder) PreVerificationGas(preVerificationGas big.Int) *OperationBuilder {
	ob.Operation.PreVerificationGas = preVerificationGas
	return ob
}

func (ob *OperationBuilder) MaxFeePerGas(maxFeePerGas big.Int) *OperationBuilder {
	ob.Operation.MaxFeePerGas = maxFeePerGas
	return ob
}

func (ob *OperationBuilder) MaxPriorityFeePerGas(maxPriorityFeePerGas big.Int) *OperationBuilder {
	ob.Operation.MaxPriorityFeePerGas = maxPriorityFeePerGas
	return ob
}

func (ob *OperationBuilder) Paymaster(paymaster common.Address) *OperationBuilder {
	ob.Operation.Paymaster = paymaster
	return ob
}

func (ob *OperationBuilder) PaymasterVerificationGasLimit(paymasterVerificationGasLimit big.Int) *OperationBuilder {
	ob.Operation.PaymasterVerificationGasLimit = paymasterVerificationGasLimit
	return ob
}

func (ob *OperationBuilder) PaymasterPostOpGasLimit(paymasterPostOpGasLimit big.Int) *OperationBuilder {
	ob.Operation.PaymasterPostOpGasLimit = paymasterPostOpGasLimit
	return ob
}

func (ob *OperationBuilder) PaymasterData(paymasterCallData []byte) *OperationBuilder {
	ob.Operation.PaymasterData = paymasterCallData
	return ob
}

func (ob *OperationBuilder) PaymasterAndData(paymaster common.Address, paymasterCallData []byte) *OperationBuilder {
	return ob.Paymaster(paymaster).PaymasterData(paymasterCallData)
}

func (ob *OperationBuilder) Signature(sig string) *OperationBuilder {
	ob.Operation.Signature = sig
	return ob
}

func (ob *OperationBuilder) Build() *PackedUserOp {
	return &PackedUserOp{}
}

func PackABIData() {

}

func SignUserOp() {

}
