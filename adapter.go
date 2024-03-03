package tbsdk

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Adapter func(...AdapterArgs) string
type AdapterArgs interface{}

var BIG_INT_ADAPTER Adapter = func(aa ...AdapterArgs) string {
	nonce, ok := aa[0].(*big.Int)
	if !ok || nonce == nil {
		return ""
	}

	return "0x" + nonce.Text(16)
}

var ADDRESS_ADAPTER Adapter = func(aa ...AdapterArgs) string {
	address, ok := aa[0].(*common.Address)
	if !ok || address == nil {
		return common.Address{}.Hex()
	}

	return address.Hex()
}

var ADDRESS_PACKED_DATA_ADAPTER Adapter = func(aa ...AdapterArgs) string {
	factoryAddress, ok := aa[0].(*common.Address)
	if !ok || factoryAddress == nil {
		return "0x"
	}

	factoryHex := factoryAddress.Hex()

	if len(aa) != 2 {
		return factoryHex
	}

	args, ok := aa[1].([]interface{})
	if !ok || len(args) != 1 {
		return factoryHex
	}

	factoryAddressData, ok := args[0].([]byte)
	if !ok || len(factoryAddressData) == 0 {
		return factoryHex
	}

	return strings.Join([]string{factoryHex, hex.EncodeToString(factoryAddressData)}, "")
}

var PACKED_DATA_ADAPTER Adapter = func(aa ...AdapterArgs) string {
	callData, ok := aa[0].([]byte)
	if !ok || len(callData) == 0 {
		return "0x"
	}

	return strings.Join([]string{"0x", hex.EncodeToString(callData)}, "")
}
