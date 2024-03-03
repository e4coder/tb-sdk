package tbsdk_test

import (
	"fmt"
	"math/big"
	"testing"

	tbsdk "github.com/e4coder/tb-sdk"
	"github.com/ethereum/go-ethereum/common"
)

var Bundler tbsdk.Bundler = tbsdk.Bundler{
	RpcUri:      "http://localhost:3000",
	ChainId:     big.NewInt(0),
	EntryPoints: []common.Address{},
}

func TestChainId(t *testing.T) {
	res, err := Bundler.Eth_chainId()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	chainId := &tbsdk.Int64ChainId{}
	err = tbsdk.DecodeResult(res.Result, chainId)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	if chainId.Value != 80001 {
		t.Logf("expected 80001 got %d", chainId.Value)
		t.FailNow()
	}
}

func TestSupportedEntryPoints(t *testing.T) {
	res, err := Bundler.Eth_supportedEntryPoints()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	entryPoints := &tbsdk.AddressArrayEntryPoints{}
	err = tbsdk.DecodeResult(res.Result, entryPoints)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for _, v := range entryPoints.Value {
		if !common.IsHexAddress(v.Hex()) {
			t.Log(v)
			t.Fail()
		}
	}
}

func TestBundlerInit(t *testing.T) {
	err := Bundler.Init("http://localhost:3000")
	if err != nil {
		t.Fail()
	}
}

func TestEstimateUserOpGas(t *testing.T) {
	// TODO! counterfactual address requires bytecode

	// err := Bundler.Init()
	// if err != nil {
	// 	t.FailNow()
	// }

	// factoryAddress := common.HexToAddress("0x1B4eBFb69B8B9A085A0483981a46Ff5AE0a1a135")
	// a, _ := abi.JSON(strings.NewReader(ABI))
	// data, _ := PackABIData(&a)
	// sender := common.Address{}
	// b, _ := abi.JSON(strings.NewReader(ABI_FACTORY))
	// owner := common.HexToAddress("0xc1a2ae5c7695314cb6cab404efd2b404c8747bd2")
	// data2, err := PackFactoryData(&b, &owner)
	// if err != nil {
	// 	t.Log(err)
	// 	t.FailNow()
	// }

	// userOp := tbsdk.NewOperationBuilder().
	// 	Nonce(big.NewInt(123)).
	// 	Sender(&sender).
	// 	CallData(data).
	// 	FactoryAndData(&factoryAddress, data2).
	// 	Signature("0x25d0a5e9d0c038f9c7c0a094b6e2a591ed8a38e8f478b65c0f14c92b0e6b0d45024daab7f8bfad5b3079234dcb155ed3fbc2bf3a8d194f9b9e1b56f1c3f9e1b1c1").
	// 	Build()

	// _, err = Bundler.Eth_estimateUserOperationGas(userOp)
	// if err != nil {
	// 	uo, _ := json.MarshalIndent(userOp, "", "  ")
	// 	t.Log(string(uo))
	// 	t.FailNow()
	// }

}
