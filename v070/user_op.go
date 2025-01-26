package v070

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/e4coder/tb-sdk/abi"
	"github.com/e4coder/tb-sdk/adapters"
	"github.com/e4coder/tb-sdk/contracts/entrypoint"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OperationBuilder struct {
	Operation UnpackedUserOperation
}

func NewOperationBuilder() *OperationBuilder {
	return &OperationBuilder{
		Operation: UnpackedUserOperation{},
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

func (ob *OperationBuilder) MaxPriorityFeePerGas(maxPriorityFeePerGas *big.Int) *OperationBuilder {
	ob.Operation.MaxPriorityFeePerGas = maxPriorityFeePerGas
	return ob
}

func (ob *OperationBuilder) MaxFeePerGas(maxFeePerGas *big.Int) *OperationBuilder {
	ob.Operation.MaxFeePerGas = maxFeePerGas
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

func (ob *OperationBuilder) Signature(sig string) *OperationBuilder {
	ob.Operation.Signature = sig
	return ob
}

func (ob *OperationBuilder) SignatureType(sigType string) *OperationBuilder {
	ob.Operation.SignatureType = sigType
	return ob
}

func (ob *OperationBuilder) Build() *RpcUnpackedUserOperation {

	unpackedOp := &RpcUnpackedUserOperation{
		Sender:               adapters.Adapt(adapters.ADDRESS_ADAPTER, ob.Operation.Sender),
		Nonce:                adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.Nonce),
		CallData:             adapters.Adapt(adapters.PACKED_DATA_ADAPTER, ob.Operation.CallData),
		CallGasLimit:         adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.CallGasLimit),
		VerificationGasLimit: adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.VerificationGasLimit),
		PreVerificationGas:   adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.PreVerificationGas),
		MaxPriorityFeePerGas: adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.MaxPriorityFeePerGas),
		MaxFeePerGas:         adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.MaxFeePerGas),
		Signature:            adapters.Adapt(adapters.TYPED_SIGNATURE_ADAPTER, ob.Operation.SignatureType, ob.Operation.Signature),
	}

	if ob.Operation.Factory != nil {
		factory := adapters.Adapt(adapters.ADDRESS_ADAPTER, ob.Operation.Factory)
		unpackedOp.Factory = &factory

		factoryData := adapters.Adapt(adapters.PACKED_DATA_ADAPTER, ob.Operation.FactoryData)
		unpackedOp.FactoryData = &factoryData

	}

	if ob.Operation.Paymaster != nil {
		paymaster := adapters.Adapt(adapters.ADDRESS_ADAPTER, ob.Operation.Paymaster)
		unpackedOp.Paymaster = &paymaster

		paymasterData := adapters.Adapt(adapters.PACKED_DATA_ADAPTER, ob.Operation.PaymasterData)
		unpackedOp.PaymasterData = &paymasterData

		paymasterPostOpGasLimit := adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.PaymasterPostOpGasLimit)
		unpackedOp.PaymasterPostOpGasLimit = &paymasterPostOpGasLimit

		paymasterVerificationGasLimit := adapters.Adapt(adapters.BIG_INT_ADAPTER, ob.Operation.PaymasterVerificationGasLimit)
		unpackedOp.PaymasterVerificationGasLimit = &paymasterVerificationGasLimit
	}

	return unpackedOp
}

func (ob *OperationBuilder) BuildWithOperationHash(chainId *big.Int, entryPoint common.Address, client *ethclient.Client) (unpackedOp *RpcUnpackedUserOperation, packedOp *PackedUserOperation, opHash []byte, err error) {
	unpackedOp = ob.Build()
	packedOp, err = ob.UnpackedToPacked()

	// TODO: fix the the GetUserOperationHash
	// opHash, err = GetUserOperationHash(packedOp, chainId, &entryPoint)

	ep, err := entrypoint.NewEntryPointV07(entryPoint, client)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("err NewEntryPointV07 : %v", err)
	}

	accountGasLimit := *(*[32]byte)(packedOp.AccountGasLimits[:32])
	gasFees := *(*[32]byte)(packedOp.GasFees[:32])

	epOp := entrypoint.PackedUserOperation{
		Sender:             *packedOp.Sender,
		Nonce:              packedOp.Nonce,
		InitCode:           packedOp.InitCode,
		CallData:           packedOp.CallData,
		AccountGasLimits:   accountGasLimit,
		PreVerificationGas: packedOp.PreVerificationGas,
		GasFees:            gasFees,
		PaymasterAndData:   packedOp.PaymasterAndData,
	}

	opHash2, err := ep.GetUserOpHash(&bind.CallOpts{}, epOp)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("err operationHash : %v", err)
	}

	opHash = opHash2[:]
	fmt.Printf("\n\n\n %x \n\n\n", opHash)
	fmt.Println(packedOp)

	return unpackedOp, packedOp, opHash, nil
}

func (ob *OperationBuilder) BuildWithSignature(chainId *big.Int, entryPoint common.Address, privKey *ecdsa.PrivateKey, client *ethclient.Client) (unpackedOp *RpcUnpackedUserOperation, packedOp *PackedUserOperation, opHash []byte, err error) {
	if chainId == nil {
		return nil, nil, nil, fmt.Errorf("chainId <nil>")
	}

	if privKey == nil {
		return nil, nil, nil, fmt.Errorf("privKey <nil>")
	}

	unpackedOp, packedOp, opHash, err = ob.BuildWithOperationHash(chainId, entryPoint, client)
	if err != nil {
		return nil, nil, nil, err
	}

	sig, err := EIP191Message(opHash, privKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to sign: %v", err)
	}
	unpackedOp.Signature = "0x" + ob.Operation.SignatureType + hex.EncodeToString(sig)

	return unpackedOp, packedOp, opHash, nil
}

func normalizeUserOperation(op *UnpackedUserOperation) *UnpackedUserOperation {
	if op.Sender == nil {
		op.Sender = &common.Address{}
	}

	if op.Nonce == nil {
		op.Nonce = common.Big0
	}

	if op.CallData == nil {
		op.CallData = []byte{}
	}

	if op.CallGasLimit == nil {
		op.CallGasLimit = common.Big0
	}

	if op.VerificationGasLimit == nil {
		op.VerificationGasLimit = common.Big0
	}

	if op.PreVerificationGas == nil {
		op.PreVerificationGas = common.Big0
	}

	if op.MaxPriorityFeePerGas == nil {
		op.MaxPriorityFeePerGas = common.Big0
	}

	if op.MaxFeePerGas == nil {
		op.MaxFeePerGas = common.Big0
	}

	if op.Factory == nil {
		op.Factory = &common.Address{}
	}

	if op.FactoryData == nil {
		op.FactoryData = []byte{}
	}

	if op.Paymaster == nil {
		op.Paymaster = &common.Address{}
	}

	if op.PaymasterVerificationGasLimit == nil {
		op.PaymasterVerificationGasLimit = common.Big0
	}

	if op.PaymasterPostOpGasLimit == nil {
		op.PaymasterPostOpGasLimit = common.Big0
	}

	if op.PaymasterData == nil {
		op.PaymasterData = []byte{}
	}

	return op
}

func (ob *OperationBuilder) UnpackedToPacked() (*PackedUserOperation, error) {
	op := &ob.Operation

	// accountGasLimits   ->  [verificationGasLimit, callGasLimit]   // 32 bytes  16,16
	var accountGasLimits []byte
	if op.VerificationGasLimit != nil && op.CallGasLimit != nil {
		accountGasLimits = abi.EncodeCustomPad(16, op.VerificationGasLimit.Bytes(), op.CallGasLimit.Bytes())
	} else {
		accountGasLimits = []byte{}
	}
	// gasFees            ->  [maxPriorityFeePerGas, maxFeePerGas]   // 32 bytes  16,16
	var gasFees []byte
	if op.MaxFeePerGas != nil && op.MaxPriorityFeePerGas != nil {
		gasFees = abi.EncodeCustomPad(16, op.MaxPriorityFeePerGas.Bytes(), op.MaxFeePerGas.Bytes())
	} else {
		gasFees = []byte{}
	}
	// initCode           ->  [factory[:20], factoryData]
	var initCode []byte
	if op.Factory != nil && op.FactoryData != nil {
		initCode = abi.EncodePacked(op.Factory.Bytes(), op.FactoryData)
	} else {
		initCode = []byte{}
	}

	// paymasterAndData   ->  [paymaster[:20], validationGasLimit[20:36], postOpLimit[36:52], ....] //20,16,16,...
	var paymasterAndData []byte
	if op.Paymaster != nil {
		paymasterAndData = abi.EncodePacked(
			op.Paymaster.Bytes(),
			abi.EncodeCustomPad(
				16,
				op.PaymasterVerificationGasLimit.Bytes(),
				op.PaymasterPostOpGasLimit.Bytes(),
			),
			op.PaymasterData,
		)
	} else {
		paymasterAndData = []byte{}
	}

	pOp := &PackedUserOperation{
		Sender:             op.Sender,
		Nonce:              op.Nonce,
		InitCode:           initCode,
		CallData:           op.CallData,
		AccountGasLimits:   accountGasLimits,
		PreVerificationGas: op.PreVerificationGas,
		GasFees:            gasFees,
		PaymasterAndData:   paymasterAndData,
	}

	return pOp, nil
}
