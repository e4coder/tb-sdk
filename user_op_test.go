package tbsdk_test

import (
	"math/big"
	"strings"
	"testing"

	tbsdk "github.com/e4coder/tb-sdk"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const ABI = `[{"constant":false,"inputs":[{"name":"_address","type":"address"},{"name":"_value","type":"uint256"}],"name":"Deposit","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
const ABI_FACTORY = `[{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"bytes32","name":"salt","type":"bytes32"}],"name":"createAccount","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"nonpayable","type":"function"}]`

func PackCallData(abi *abi.ABI) ([]byte, error) {
	data, err := abi.Pack("execute", common.MaxAddress, big.NewInt(100000000), []byte{})

	return data, err
}

func PackFactoryData(abi *abi.ABI, addr *common.Address, salt [32]byte) ([]byte, error) {
	data, err := abi.Pack("createAccount", addr, salt)

	return data, err
}

func TestBuilder(t *testing.T) {
	a, _ := abi.JSON(strings.NewReader(ABI))
	data, _ := PackCallData(&a)

	b, _ := abi.JSON(strings.NewReader(ABI_FACTORY))
	data2, _ := PackCallData(&b)

	addr23 := common.BigToAddress(big.NewInt(23))

	tbsdk.NewOperationBuilder().
		Sender(&addr23).
		Nonce(nil).
		CallData(data).
		FactoryAndData(&addr23, data2).
		PaymasterAndData(&addr23, data).
		Signature("0x25d0a5e9d0c038f9c7c0a094b6e2a591ed8a38e8f478b65c0f14c92b0e6b0d45024daab7f8bfad5b3079234dcb155ed3fbc2bf3a8d194f9b9e1b56f1c3f9e1b1c1").
		Build()

}
