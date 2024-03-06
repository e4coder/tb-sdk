package tbsdk

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"fmt"
	"math/big"
)

/*
*	TODO: Adopt sigutils in utils with AbiEncode and AbiEncodePacked
 */

func padHex(hex string) string {
	// add 24 zeros to the left
	return strings.Repeat("0", 24) + strings.TrimPrefix(hex, "0x")
}

func encodeAddress(value string) string {
	return padHex(strings.ToLower(string(value)))
}

func encodeAbiParameters(types []string, values []interface{}) ([]byte, error) {
	if len(types) != len(values) {
		return nil, fmt.Errorf("mismatch between types and values length")
	}

	var encodedData []byte

	for i := 0; i < len(types); i++ {
		switch types[i] {
		case "address":
			encodedStr := encodeAddress(values[i].(string))
			addressBytes := hexToBytes(encodedStr)
			encodedData = append(encodedData, addressBytes...)
		case "uint256":
			intBytes := values[i].(*big.Int).Bytes()
			// Ensure uint256 is 32 bytes long
			padSize := 32 - len(intBytes)
			padBytes := make([]byte, padSize)
			encodedData = append(encodedData, padBytes...)
			encodedData = append(encodedData, intBytes...)
		case "bytes32":
			bytes32Value := hexToBytes32(values[i].(string))
			encodedData = append(encodedData, bytes32Value[:]...)
		default:
			return nil, fmt.Errorf("unsupported type: %s", types[i])
		}
	}

	return encodedData, nil
}

func keccak256(data string) []byte {
	// remove 0x prefix if exists
	data = strings.TrimPrefix(data, "0x")

	// convert hex string to bytes
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		// fmt.Printf("failed to decode hex string\n",err)
		fmt.Printf("failed to decode hex string %s\n", err)
		return []byte{}
	}

	// calculate hash
	hash := crypto.Keccak256Hash(dataBytes).Bytes()
	return hash
}

func hexToBytes32(hexStr string) [32]byte {
	var bytes32Value [32]byte
	copy(bytes32Value[:], hexToBytes(hexStr))
	return bytes32Value
}

func hexToBytes(hexStr string) []byte {
	data, _ := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	return data
}

func hexToBigInt(hex string) *big.Int {
	// remove 0x prefix if exists
	hex = strings.TrimPrefix(hex, "0x")

	value := new(big.Int)
	value.SetString(hex, 16)
	return value
}

func GetUserOperationHash(request *PackedUserOp, chainID int64, entryPointAddress common.Address) ([]byte, error) {
	uoBytes, err := PackUserOperation(request)
	if err != nil {
		return nil, err
	}

	// uoBytes to hex string
	uoHex := hex.EncodeToString(uoBytes)

	uoHexKeccak := "0x" + hex.EncodeToString(keccak256(uoHex))

	eap, err := encodeAbiParameters(
		[]string{
			"bytes32",
			"address",
			"uint256",
		},
		[]interface{}{
			uoHexKeccak,
			entryPointAddress.String(),
			big.NewInt(chainID),
		},
	)

	if err != nil {
		return nil, err
	}

	out := keccak256(hex.EncodeToString(eap))
	return out, nil
}

func PackUserOperation(request *PackedUserOp) ([]byte, error) {
	// byte arrays
	hashedInitCode := "0x" + hex.EncodeToString(keccak256(request.InitCode))
	hashedCallData := "0x" + hex.EncodeToString(keccak256(request.CallData))
	hashedPaymasterAndData := "0x" + hex.EncodeToString(keccak256(request.PaymasterAndData))

	return encodeAbiParameters(
		[]string{
			"address",
			"uint256",
			"bytes32",
			"bytes32",
			"uint256",
			"uint256",
			"uint256",
			"uint256",
			"uint256",
			"bytes32",
		},
		[]interface{}{
			request.Sender,
			hexToBigInt(request.Nonce),
			hashedInitCode,
			hashedCallData,
			hexToBigInt(request.CallGasLimit),
			hexToBigInt(request.VerificationGasLimit),
			hexToBigInt(request.PreVerificationGas),
			hexToBigInt(request.MaxFeePerGas),
			hexToBigInt(request.MaxPriorityFeePerGas),
			hashedPaymasterAndData,
		},
	)
}

func SignDataWithEthereumPrivateKey(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	return signDataHashWithEthereumPrivateKey(data, privKey)
}

func signDataHashWithEthereumPrivateKey(dataToSign []byte, privateKeyECDSA *ecdsa.PrivateKey) ([]byte, error) {
	if len(dataToSign) != 32 {
		return nil, errors.New("dataToSign must be 32 bytes long")
	}

	// Prepend the "Ethereum Signed Message" prefix.
	prefix := "\x19Ethereum Signed Message:\n32" + string(dataToSign)

	// Hash the message using Keccak-256.
	hash := crypto.Keccak256([]byte(prefix))

	signature, err := crypto.Sign(hash, privateKeyECDSA)
	if err != nil {
		return nil, err
	}
	//return CompactSignature(signature)

	// TODO: fix +27 please!
	// Encode the signature in Ethereum's compact signature format.
	out := append(signature[:64], signature[64]+27)
	return out, nil
}
