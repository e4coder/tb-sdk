package tbsdk_test

import (
	"testing"

	tbsdk "github.com/e4coder/tb-sdk"
	"github.com/ethereum/go-ethereum/common"
)

func TestBuilder(t *testing.T) {
	// zeroAddress := common.BigToAddress(common.Big0).Hex()
	// oneAddress := common.BigToAddress(common.Big1).Hex()
	ob := tbsdk.NewOperationBuilder()
	ob.
		Sender(common.Address{}).
		Nonce().
		Factory(common.BigToAddress(common.Big1)).
		FactoryData([]byte("Hello World")).
		CallData([]byte("Call Data")).
		CallGasLimit(*common.Big2).
		VerificationGasLimit(*common.Big1).
		PreVerificationGas(*common.Big2).
		MaxFeePerGas(*common.Big1).
		MaxPriorityFeePerGas(*common.Big2).
		Paymaster(common.BigToAddress(common.Big2)).
		PaymasterData([]byte("PaymasterData")).
		PaymasterVerificationGasLimit(*common.Big0).
		PaymasterPostOpGasLimit(*common.Big2).
		Signature("signature")

	ob.Build()

	// if (userOp.Sender == zeroAddress) &&
	// 	(userOp.Factory == oneAddress) &&
	// 	(hex.EncodeToString(userOp.CallData) == hex.EncodeToString([]byte("Call Data"))) {
	// 	return
	// }
}
