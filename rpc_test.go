package tbsdk_test

import (
	"fmt"
	"math/big"
	"testing"

	tbsdk "github.com/e4coder/tb-sdk"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
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
	Bundler.Eth_estimateUserOperationGas(tbsdk.NewOperationBuilder().Build(), &common.MaxAddress, map[common.Address]gethclient.OverrideAccount{})

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

const ABI_ACCOUNT = `[
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "dest",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "func",
                "type": "bytes"
            }
        ],
        "name": "execute",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`

func TestEstimateUserOpGasNewAccount(t *testing.T) {
	// sign := "0xfffffffffffffffffffffffffffffff0000000000000000000000000000000007aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1c"

	// factoryABI, err := abi.JSON(strings.NewReader(ABI_FACTORY))
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Log(err)
	// 	t.FailNow()
	// }
	// accountFACTORY, err := abi.JSON(strings.NewReader(ABI_ACCOUNT))
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Log(err)
	// 	t.FailNow()
	// }

	// err = Bundler.Init("http://localhost:3000")
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.FailNow()
	// }
	// factoryAddress := common.HexToAddress("0x4BBa2E1c4856228c0572f7b64f14916E2F091391")
	// // accountAddress := common.HexToAddress("0xd40aeab1d9e7c57523c2f5381f79c9738a73fe2d")
	// owner := common.HexToAddress("0xB77e8533ce316AC7355191bf775F1Be1967650eD")

	// callData, err := PackABIData(&accountFACTORY)
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Log(err)
	// 	t.FailNow()
	// }
	// t.Log("call data")
	// t.Log(callData)
	// var salt32 [32]byte
	// salt, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000002")
	// copy(salt32[:], salt[:32])

	// factory := factoryAddress.Bytes()
	// ar, _ := hex.DecodeString("F235B58DC3b2169136A857B06aaedcE1aEC4c667")
	// ep, _ := hex.DecodeString("8024A70A99d35FF24Cacc861e946945530ee96A3")
	// creationCode, _ := hex.DecodeString(TestAccountContractCreationCode)

	// initCode := tbsdk.AbiEncodePacked(creationCode, tbsdk.AbiEncode(
	// 	owner.Bytes(),
	// 	ar,
	// 	ep,
	// 	factory,
	// ))

	// account := tbsdk.ComputeCounterfactualAddress(owner.Bytes(), salt, initCode)
	// accountAddress := common.HexToAddress(account)

	// factoryData, err := PackFactoryData(&factoryABI, &accountAddress, salt32)
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Log(err)
	// 	t.FailNow()
	// }

	// userOp := tbsdk.NewOperationBuilder().
	// 	Nonce(big.NewInt(123)).
	// 	Sender(&accountAddress).
	// 	CallData(callData).
	// 	FactoryAndData(&factoryAddress, factoryData).
	// 	Signature(sign).
	// 	Build()

	// _, err = Bundler.Eth_estimateUserOperationGas(userOp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	uo, _ := json.MarshalIndent(userOp, "", "  ")
	// 	t.Log(string(uo))
	// 	t.FailNow()
	// }

}

func TestEstimateUserOpGasExistingAccount(t *testing.T) {
	// sign := "0xfffffffffffffffffffffffffffffff0000000000000000000000000000000007aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1c"
	// accountFACTORY, err := abi.JSON(strings.NewReader(ABI_ACCOUNT))
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Log(err)
	// 	t.FailNow()
	// }

	// err = Bundler.Init("http://localhost:3000")
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.FailNow()
	// }
	// accountAddress := common.HexToAddress("0xd40aeab1d9e7c57523c2f5381f79c9738a73fe2d")

	// callData, err := PackABIData(&accountFACTORY)
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Log(err)
	// 	t.FailNow()
	// }
	// t.Log("call data")
	// t.Log(callData)

	// userOp := tbsdk.NewOperationBuilder().
	// 	Nonce(big.NewInt(123)).
	// 	Sender(&accountAddress).
	// 	CallData(callData).
	// 	Signature(sign).
	// 	Build()

	// _, err = Bundler.Eth_estimateUserOperationGas(userOp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	uo, _ := json.MarshalIndent(userOp, "", "  ")
	// 	t.Log(string(uo))
	// 	t.FailNow()
	// }
}
