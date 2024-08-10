package tbsdk

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
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

func (b *Bundler) _call(method string, params interface{}) (*RpcResponse, error) {
	request, err := PrepareRPCCall(b.RpcUri, method, params)
	if err != nil {
		return nil, err
	}

	return HandleRpcRequest(request, Client)
}

func (b *Bundler) Eth_chainId() (*RpcResponse, error) {
	params := []interface{}{}
	return b._call("eth_chainId", params)
}

func (b *Bundler) Eth_supportedEntryPoints() (*RpcResponse, error) {
	params := []interface{}{}
	return b._call("eth_supportedEntryPoints", params)
}

func (b *Bundler) Eth_estimateUserOperationGas(userOp *PackedUserOp, entrypoint *common.Address, stateOverrideSet map[common.Address]gethclient.OverrideAccount) (*RpcResponse, error) {
	if entrypoint == nil {
		entrypoint = &b.EntryPoints[0]
	}

	var params []interface{}
	if stateOverrideSet != nil {
		params = []interface{}{userOp, *entrypoint, stateOverrideSet}
	} else {
		params = []interface{}{userOp, *entrypoint}
	}

	response, err := b._call("eth_estimateUserOperationGas", params)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, errors.New("rpc_error")
	}

	return response, nil
}

func (b *Bundler) Eth_sendUserOperation(userOp *PackedUserOp, entrypoint *common.Address) (*RpcResponse, error) {
	if entrypoint == nil {
		entrypoint = &b.EntryPoints[0]
	}
	params := []interface{}{userOp, *entrypoint}

	response, err := b._call("eth_sendUserOperation", params)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, errors.New("rpc_error")
	}

	return response, nil
}

func (b *Bundler) Eth_getUserOperationByHash(userOpHash string) (*RpcResponse, error) {
	params := []interface{}{userOpHash}

	response, err := b._call("eth_getUserOperationByHash", params)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, errors.New("rpc_error")
	}

	return response, nil
}

func (b *Bundler) Eth_getUserOperationReceipt(userOpHash string) (*RpcResponse, error) {
	params := []interface{}{userOpHash}

	response, err := b._call("eth_getUserOperationReceipt", params)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, errors.New("rpc_error")
	}

	return response, nil
}

func (b *Bundler) Rundler_maxPriorityFeePerGas() (*RpcResponse, error) {
	params := []interface{}{}

	response, err := b._call("rundler_maxPriorityFeePerGas", params)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, errors.New("rpc_error")
	}

	return response, nil
}

// TODO Debug_ implement namespace

func (b *Bundler) Debug_clearState() {}

func (b *Bundler) Debug_dumpMempool() {}

func (b *Bundler) Debug_sendBundleNow() {}

func (b *Bundler) Debug_setBundlingMode() {}

func (b *Bundler) Debug_setReputation() {}

func (b *Bundler) Debug_dumpReputation() {}
