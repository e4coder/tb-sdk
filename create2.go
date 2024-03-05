package tbsdk

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// initCode = creationCode + packed(constructorArguments)
func ComputeCounterfactualAddress(sender, salt, initCode []byte) string {
	start, _ := hex.DecodeString("ff")
	address := fmt.Sprintf("0x%s", hex.EncodeToString(crypto.Keccak256(start, sender, salt, crypto.Keccak256(initCode))[12:]))

	return address
}
