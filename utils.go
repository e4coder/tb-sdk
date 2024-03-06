package tbsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

func PrepareRPCCall(endpoint, method string, params interface{}) (*http.Request, error) {
	requestBody := &RpcRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		"POST",
		endpoint,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, errors.Join(
			ErrRequestCreationFailed,
			fmt.Errorf("endpoint->%s,body->[:%s:]", endpoint, string(body)),
			err,
		)
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func HandleRpcRequest(request *http.Request, client *http.Client) (*RpcResponse, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody := &RpcResponse{}
	err = json.NewDecoder(response.Body).Decode(responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// TODO: input type ...interface{} output ([]byte, error)
func AbiEncode(args ...[]byte) []byte {
	encodedArgs := []byte{}
	for _, v := range args {
		encodedArgs = append(encodedArgs, common.LeftPadBytes(v, 32)...)
	}

	return encodedArgs
}

// TODO: input type ...interface{} output ([]byte, error)
func AbiEncodePacked(args ...[]byte) []byte {
	packedArgs := []byte{}
	for _, v := range args {
		packedArgs = append(packedArgs, v...)
	}

	return packedArgs
}
