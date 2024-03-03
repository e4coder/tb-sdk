package tbsdk

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type Bundler struct {
	RpcUri      string
	ChainId     *big.Int
	EntryPoints []common.Address
}

var Client *http.Client = &http.Client{}

func NewBundler() *Bundler {
	return &Bundler{}
}

func (b *Bundler) Init(rpcUri string) error {
	b.RpcUri = rpcUri

	res, err := b.Eth_chainId()
	if err != nil {
		return err
	}

	if res.Result == nil {
		fmt.Println(res.Error)
		return err
	}

	chainId := &Int64ChainId{}
	err = DecodeResult(res.Result, chainId)
	if err != nil {
		return err
	}

	res2, err := b.Eth_supportedEntryPoints()
	if err != nil {
		return err
	}

	entryPoints := &AddressArrayEntryPoints{}
	err = DecodeResult(res2.Result, entryPoints)
	if err != nil {
		return err
	}

	b.ChainId = big.NewInt(chainId.Value)
	b.EntryPoints = entryPoints.Value

	return nil
}

func (b *Bundler) Eth_chainId() (*RpcResponse, error) {
	params := []interface{}{}

	request, err := PrepareRPCCall(b.RpcUri, "eth_chainId", params)
	if err != nil {
		return nil, err
	}

	return HandleRpcRequest(request, Client)
}

func (b *Bundler) Eth_supportedEntryPoints() (*RpcResponse, error) {
	params := []interface{}{}

	request, err := PrepareRPCCall(b.RpcUri, "eth_supportedEntryPoints", params)
	if err != nil {
		return nil, err
	}

	return HandleRpcRequest(request, Client)
}

func (b *Bundler) Eth_estimateUserOperationGas(userOp *PackedUserOp) (*RpcResponse, error) {
	params := []interface{}{userOp, b.EntryPoints[0]}

	request, err := PrepareRPCCall(b.RpcUri, "eth_estimateUserOperationGas", params)
	if err != nil {
		return nil, err
	}

	response, err := HandleRpcRequest(request, Client)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, errors.New("rpc_error")
	}

	return response, nil
}

func (b *Bundler) Eth_sendUserOperation(userOp UserOp) {}

func (b *Bundler) Eth_getUserOperationByHash() {}

func (b *Bundler) Eth_getUserOperationReceipt() {}

func (b *Bundler) Debug_clearState() {}

func (b *Bundler) Debug_dumpMempool() {}

func (b *Bundler) Debug_sendBundleNow() {}

func (b *Bundler) Debug_setBundlingMode() {}

func (b *Bundler) Debug_setReputation() {}

func (b *Bundler) Debug_dumpReputation() {}
