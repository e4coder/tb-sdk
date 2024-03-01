package tbsdk

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Decoder interface {
	Decode(interface{}) error
}

type Int64ChainId struct {
	Value int64
}

func (i *Int64ChainId) Decode(rawResult interface{}) error {
	if rawResult == nil {
		return fmt.Errorf("rawResult not found")
	}

	hexVal, ok := rawResult.(string)
	if !ok {
		return fmt.Errorf("invalid type, expected string for rawResult")
	}
	result, err := strconv.ParseInt(strings.TrimPrefix(hexVal, "0x"), 16, 64)
	if err != nil {
		return err
	}
	i.Value = result
	return nil
}

type AddressArrayEntryPoints struct {
	Value []common.Address
}

func (i *AddressArrayEntryPoints) Decode(rawResult interface{}) error {
	if rawResult == nil {
		return fmt.Errorf("rawResult not found")
	}

	resSlice, ok := rawResult.([]interface{})
	if !ok {
		return fmt.Errorf("invalid type, expected []interface{} for rawResult")
	}

	for _, v := range resSlice {
		strValue, ok := v.(string)
		if !ok {
			return fmt.Errorf("invalid type, expected string for strValue")
		}
		i.Value = append(i.Value, common.HexToAddress(strValue))
	}
	return nil
}

func DecodeResult[T Decoder](r interface{}, dest T) error {
	err := dest.Decode(r)
	if err != nil {
		return err
	}
	return nil
}
