package tbsdk

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	v070 "github.com/e4coder/tb-sdk/v070"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

type Bundler struct {
	RpcUri      string
	ChainId     *big.Int
	EntryPoints []common.Address
	EthClient   *ethclient.Client
}

var Client *http.Client = &http.Client{}

func NewBundler() *Bundler {
	return &Bundler{}
}

func (b *Bundler) Init(rpcUri string, client *ethclient.Client) error {
	b.RpcUri = rpcUri
	b.EthClient = client

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

// map[callGasLimit:0x78fb paymasterPostOpGasLimit:0x0 paymasterVerificationGasLimit:0x0 preVerificationGas:0xc870 verificationGasLimit:0x10cd9]
type Eth_estimateUserOperationGasResponse struct {
	CallGasLimit                  *big.Int `json:"callGasLimit`
	PaymasterPostOpGasLimit       *big.Int `json:"paymasterPostOpGasLimit"`
	PaymasterVerificationGasLimit *big.Int `json:"paymasterVerificationGasLimit"`
	PreVerificationGas            *big.Int `json:"preVerificationGas"`
	VerificationGasLimit          *big.Int `json:"preVerificationGas"`
}

func (b *Bundler) Eth_estimateUserOperationGas_v070(userOp *v070.RpcUnpackedUserOperation, entrypoint *common.Address, stateOverrideSet map[common.Address]gethclient.OverrideAccount) (*RpcResponse, *Eth_estimateUserOperationGasResponse, error) {
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
		return nil, nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, nil, errors.New("rpc_error")
	}

	result, ok := response.Result.(map[string]interface{})
	if !ok {
		return response, nil, errors.New("result not ok")
	}

	fmt.Println(result)
	estimations := &Eth_estimateUserOperationGasResponse{
		CallGasLimit:                  _gasToBigInt(result["callGasLimit"].(string)),
		PaymasterPostOpGasLimit:       _gasToBigInt(result["paymasterPostOpGasLimit"].(string)),
		PaymasterVerificationGasLimit: _gasToBigInt(result["paymasterVerificationGasLimit"].(string)),
		PreVerificationGas:            _gasToBigInt(result["preVerificationGas"].(string)),
		VerificationGasLimit:          _gasToBigInt(result["verificationGasLimit"].(string)),
		// verificationGasLimit
	}

	return response, estimations, nil
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

func (b *Bundler) Eth_sendUserOperation_v070(userOp *v070.RpcUnpackedUserOperation, entrypoint *common.Address) (*RpcResponse, error) {
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

type Pimlico_getUserOperationGasPriceValues struct {
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
}

// map[fast:map[maxFeePerGas:0x3c063a21 maxPriorityFeePerGas:0x3b9aca00] slow:map[maxFeePerGas:0x3c063a21 maxPriorityFeePerGas:0x3b9aca00] standard:map[maxFeePerGas:0x3c063a21 maxPriorityFeePerGas:0x3b9aca00
type Pimlico_getUserOperationGasPriceResponse struct {
	Slow     Pimlico_getUserOperationGasPriceValues `json:"slow"`
	Standard Pimlico_getUserOperationGasPriceValues `json:"standard"`
	Fast     Pimlico_getUserOperationGasPriceValues `json:"fast"`
}

func (b *Bundler) Pimlico_getUserOperationGasPrice() (*RpcResponse, *Pimlico_getUserOperationGasPriceResponse, error) {
	params := []interface{}{}

	response, err := b._call("pimlico_getUserOperationGasPrice", params)

	if err != nil {
		return nil, nil, err
	}

	if response.Error != nil {
		fmt.Println(response.Error)

		return response, nil, errors.New("rpc_error")
	}

	result, ok := response.Result.(map[string]interface{})
	if !ok {
		return response, nil, errors.New("result not ok")
	}

	fast := result["fast"].(map[string]interface{})
	slow := result["slow"].(map[string]interface{})
	standard := result["standard"].(map[string]interface{})

	opGasPrice := &Pimlico_getUserOperationGasPriceResponse{
		Standard: Pimlico_getUserOperationGasPriceValues{
			MaxFeePerGas:         _gasToBigInt(standard["maxFeePerGas"].(string)),
			MaxPriorityFeePerGas: _gasToBigInt(standard["maxPriorityFeePerGas"].(string)),
		},
		Slow: Pimlico_getUserOperationGasPriceValues{
			MaxFeePerGas:         _gasToBigInt(slow["maxFeePerGas"].(string)),
			MaxPriorityFeePerGas: _gasToBigInt(slow["maxPriorityFeePerGas"].(string)),
		},
		Fast: Pimlico_getUserOperationGasPriceValues{
			MaxFeePerGas:         _gasToBigInt(fast["maxFeePerGas"].(string)),
			MaxPriorityFeePerGas: _gasToBigInt(fast["maxPriorityFeePerGas"].(string)),
		},
	}

	return response, opGasPrice, nil

}

func _gasToBigInt(val string) *big.Int {
	value, success := big.NewInt(0).SetString(strings.TrimPrefix(val, "0x"), 16)
	if !success {
		return big.NewInt(0)
	}

	return value
}

// TODO Debug_ implement namespace

func (b *Bundler) Debug_clearState() {}

func (b *Bundler) Debug_dumpMempool() {}

func (b *Bundler) Debug_sendBundleNow() {}

func (b *Bundler) Debug_setBundlingMode() {}

func (b *Bundler) Debug_setReputation() {}

func (b *Bundler) Debug_dumpReputation() {}
