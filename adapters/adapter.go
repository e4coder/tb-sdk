package adapters

import (
	"encoding/hex"
	"fmt"
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

var TYPED_SIGNATURE_ADAPTER Adapter = func(aa ...AdapterArgs) string {
	fmt.Printf("IN ADAPTER")
	fmt.Printf("%v", aa)
	sigType, ok := aa[0].(string) // 00 | 01 | 02
	if !ok {
		return fmt.Sprintf("0xf1%v", sigType)
	}

	signature, ok := aa[1].([]interface{})[0].(string)
	if !ok {
		return fmt.Sprintf("0xff%v", signature)
	}

	signature = strings.TrimPrefix(signature, "0x")

	return fmt.Sprintf("0x%s%s", sigType, signature)
}

func Adapt(adapter Adapter, val interface{}, args ...interface{}) string {
	return adapter(val, args)
}
