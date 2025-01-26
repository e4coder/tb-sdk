package abi

import "github.com/ethereum/go-ethereum/common"

// Encodes 32 byte padded values
//
// TODO: input type ...interface{} output ([]byte, error)
func Encode(args ...[]byte) []byte {
	encodedArgs := []byte{}
	for _, v := range args {
		encodedArgs = append(encodedArgs, common.LeftPadBytes(v, 32)...)
	}

	return encodedArgs
}

// Encodes bytes with custom padding
//
// TODO: should only accept args with length less then
func EncodeCustomPad(pad int, args ...[]byte) []byte {
	encodedArgs := []byte{}
	for _, v := range args {
		encodedArgs = append(encodedArgs, common.LeftPadBytes(v, pad)...)
	}

	return encodedArgs
}

// Closely encode bytes
//
// TODO: input type ...interface{} output ([]byte, error)
func EncodePacked(args ...[]byte) []byte {
	packedArgs := []byte{}
	for _, v := range args {
		packedArgs = append(packedArgs, v...)
	}

	return packedArgs
}
