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

	fmt.Println(chainId.Value)

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
	err := Bundler.Init()
	if err != nil {
		t.Fail()
	}
}
