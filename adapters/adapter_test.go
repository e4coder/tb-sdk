package adapters_test

import (
	"math/big"
	"testing"

	"github.com/e4coder/tb-sdk/adapters"
	"github.com/ethereum/go-ethereum/common"
)

func TestNonceAdapter(t *testing.T) {
	val := adapters.BIG_INT_ADAPTER(big.NewInt(324))
	if val != "0x144" {
		t.Log(val)
		t.Fail()
	}

	val2 := adapters.BIG_INT_ADAPTER(nil)
	if val2 != "" {
		t.Log(val2)
		t.Fail()
	}
}

func TestAddressAdapter(t *testing.T) {
	zeroAddress := common.Address{}.Hex()
	val := adapters.ADDRESS_ADAPTER(nil)

	if val != zeroAddress {
		t.Log("FAILED NIL ADDRESS")
		t.Log(val)
		t.Fail()
	}

	val2 := adapters.ADDRESS_ADAPTER(&common.MaxAddress)

	if val2 != common.MaxAddress.Hex() {
		t.Log("FAILED MAX ADDRESS")
		t.Log(val)
		t.Fail()
	}
}

func TestPackedDataAdapter(t *testing.T) {
	val := adapters.PACKED_DATA_ADAPTER([]byte("Hello"))

	if val != "0x48656c6c6f" {
		t.Log(val)
		t.Fail()
	}

	val2 := adapters.PACKED_DATA_ADAPTER(nil)

	if val2 != "0x" {
		t.Log(val2)
		t.Fail()
	}

	val3 := adapters.PACKED_DATA_ADAPTER([]byte{})

	if val3 != "0x" {
		t.Log(val3)
		t.Fail()
	}
}

func TestAddressPackedDataAdapater(t *testing.T) {

	val := adapters.ADDRESS_PACKED_DATA_ADAPTER(&common.MaxAddress, []interface{}{[]byte("Hello")})

	if val != "0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF48656c6c6f" {
		t.Log(val)
		t.Fail()
	}

	val2 := adapters.ADDRESS_PACKED_DATA_ADAPTER(&common.MaxAddress, nil)

	if val2 != "0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF" {
		t.Log(val2)
		t.Fail()
	}

	val3 := adapters.ADDRESS_PACKED_DATA_ADAPTER(nil, nil)

	if val3 != "0x" {
		t.Log(val3)
		t.Fail()
	}

	val4 := adapters.ADDRESS_PACKED_DATA_ADAPTER(nil, []interface{}{})

	if val4 != "0x" {
		t.Log(val4)
		t.Fail()
	}

	val5 := adapters.ADDRESS_PACKED_DATA_ADAPTER(&common.MaxAddress, []interface{}{})

	if val5 != "0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF" {
		t.Log(val5)
		t.Fail()
	}
}
