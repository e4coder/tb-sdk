package v070

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/e4coder/tb-sdk/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func encodeAbiParameters(types []string, values [][]byte) ([]byte, error) {
	if len(types) != len(values) {
		return nil, fmt.Errorf("mismatch between types and values length")
	}

	var encodedData []byte
	for i := 0; i < len(types); i++ {
		switch types[i] {
		case "pad":
			encodedData = append(encodedData, abi.Encode(values[i])...)
		case "nopad":
			encodedData = append(encodedData, values[i]...)
		}
	}

	return encodedData, nil
}

func PackUserOperation(request *PackedUserOperation) ([]byte, error) {
	initCodeHash := crypto.Keccak256(request.InitCode)
	callDataHash := crypto.Keccak256(request.CallData)
	paymasterAndDataHash := crypto.Keccak256(request.PaymasterAndData)
	packedUserOpType := "PackedUserOperation(address sender,uint256 nonce,bytes initCode,bytes callData,bytes32 accountGasLimits,uint256 preVerificationGas,bytes32 gasFees,bytes paymasterAndData)"

	packedUserOpTypeHash := crypto.Keccak256([]byte(packedUserOpType))

	return encodeAbiParameters(
		[]string{
			"nopad",
			"pad",   // sender
			"pad",   // nonce
			"nopad", // keccak256(initCode)
			"nopad", // keccak256(callData)
			"nopad", // accountGasLimits
			"pad",   // preVerificationGas
			"nopad", // gasFees
			"nopad", // keccak256(paymasterAndData)
		},
		[][]byte{
			packedUserOpTypeHash,
			request.Sender.Bytes(),
			request.Nonce.Bytes(),
			initCodeHash,
			callDataHash,
			request.AccountGasLimits,
			request.PreVerificationGas.Bytes(),
			request.GasFees,
			paymasterAndDataHash,
		},
	)
}

// TODO: Not working properly fix this
//
// in the meanwhile you can use GetUseOpHash from the entrypoint contract
func GetUserOperationHash(request *PackedUserOperation, chainID *big.Int, entryPointAddress *common.Address) ([]byte, error) {
	uoBytes, err := PackUserOperation(request)
	if err != nil {
		return nil, err
	}

	uoBytesHash := crypto.Keccak256(uoBytes)

	domainSeperator := BuildDomainSeperatorV4(chainID, entryPointAddress)

	userOpHash, err := ToTypedDatahash(domainSeperator, uoBytesHash)

	if err != nil {
		return nil, errors.Join(errors.New("failure to build userOpHash"), err)
	}

	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(userOpHash), nil
}

func EIP191Message(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	if len(message) != 32 {
		return nil, errors.New("'message' must be 32 bytes long")
	}

	prefix := "\x19Ethereum Signed Message:\n32" + string(message)

	hash := crypto.Keccak256([]byte(prefix))

	fmt.Printf("\nSIG PREFIXED HASH %x\n\n\n", hash)

	signature, err := crypto.Sign(hash, privKey)
	if err != nil {
		return nil, err
	}

	// return CompactSignature(signature)

	// TODO: fix +27 please!
	// Encode the signature in Ethereum's compact signature format.
	out := append(signature[:64], signature[64]+27) //nolint:all

	return out, nil
}

//	string constant internal DOMAIN_NAME = "ERC4337";
//
// string constant internal DOMAIN_VERSION = "1";
//
//	bytes32 private constant TYPE_HASH =
//
// keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
//
//	function _buildDomainSeparator() private view returns (bytes32) {
//	        return keccak256(abi.encode(TYPE_HASH, _hashedName, _hashedVersion, block.chainid, address(this)));
//	    }
const EIP712_ENTRYPOINT_DOMAIN_NAME = "ERC4337"
const EIP712_ENTRYPOINT_DOMAIN_VERSION = "1"
const EIP712_ENTRYPOINT_DOMAIN_TYPE = "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"

func BuildDomainSeperatorV4(chainId *big.Int, entryPointAddress *common.Address) []byte {
	typeHash := crypto.Keccak256([]byte(EIP712_ENTRYPOINT_DOMAIN_TYPE))

	return crypto.Keccak256(abi.Encode(typeHash, []byte(EIP712_ENTRYPOINT_DOMAIN_NAME), []byte(EIP712_ENTRYPOINT_DOMAIN_VERSION), chainId.Bytes(), entryPointAddress.Bytes()))
}

//
//     function toTypedDataHash(bytes32 domainSeparator, bytes32 structHash) internal pure returns (bytes32 digest) {
//         assembly ("memory-safe") {
//             let ptr := mload(0x40)
//             mstore(ptr, hex"19_01")
//             mstore(add(ptr, 0x02), domainSeparator)
//             mstore(add(ptr, 0x22), structHash)
//             digest := keccak256(ptr, 0x42)
//         }
//     }
// }
//

const TYPED_DATA_PREFIX_Ox19_01 = "1901"

func ToTypedDatahash(domainSeparator []byte, structHash []byte) ([]byte, error) {
	Ox19_01, err := hex.DecodeString(TYPED_DATA_PREFIX_Ox19_01)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("invalid 1901 %s %x", TYPED_DATA_PREFIX_Ox19_01, Ox19_01), err)
	}
	typedDataHaSH := crypto.Keccak256(abi.EncodePacked(Ox19_01, domainSeparator, structHash))

	return typedDataHaSH, nil
}
