package tokenapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// In debug mode for example file checksums are skipped, because example data contains invalid ones
const DEBUG_MODE = true

const TOKEN_VALID_MINUTES = 10

type Method struct {
	Name        string `json:"name"`
	Args        string `json:"args"`
	RetType     string `json:"retType"`
	Description string `json:"description"`
}

type Ret struct {
	RetValue string `json:retValue`
	RetType  string `json:retType`
}

type Token struct {
	ID                 string    `json:"ID"` //Must be string
	UserRequested      string    `json:userRequested`
	ChaincodeName      string    `json:chaincodeName`
	Method             string    `json:method`
	Arguments          string    `json:arguments`
	Ret                Ret       `json:"ret,omitempty" metadata:"ret,optional"`
	TimeRequested      time.Time `json:timeRequested`
	ExpirationTime     time.Time `json:expirationTime`
	Description        string    `json:description`
	DirectlyExecutable bool      `json:directlyExecutable`
}

type Interface struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Options struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type Connection struct {
	ID   string `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Node struct {
	Type               string      `json:"type"`
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	ChaincodeName      string      `json:"chaincodeName"`
	TokenId            string      `json:"tokenId"`
	MethodName         string      `json:"methodName"`
	Options            []Options   `json:"options"`
	Interfaces         []Interface `json:"interfaces"`
	DirectlyExecutable bool        `json:"directlyExecutable"`
}

type Flow struct {
	Nodes       []Node       `json:"nodes"`
	Connections []Connection `json:"connections"`
}

// Check whether nonce provided in parameter is equal to actual nonce. Can be used to check
// whether method has been executed by other method/chaincode. We should use this function
// to forbid users from direct access to algorithm function. Each algorithm should use this
// function before computation starts
// Nonce is in fact ctx.GetStub().GetCreator() converted to base64
func IsNonceValid(ctx contractapi.TransactionContextInterface, nonceStr string) (bool, error) {
	creatorByte, err := ctx.GetStub().GetCreator()
	if err != nil {
		return false, err
	}

	creator := base64.StdEncoding.EncodeToString(creatorByte)

	fmt.Printf("tokenapi:isNonceValid: Comparing GetCreator(): %s vs nonce: %s\n", creator, nonceStr)
	return creator == nonceStr, nil
}

// Compute method launches alghoritm provided it has correct token
func Compute(ctx contractapi.TransactionContextInterface, id string) (*Token, error) {
	decodedId, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:ReadToken: failed to decode ID: %v", err)
	}
	tokenJSON, err := ctx.GetStub().GetState(string(decodedId))
	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:Compute: failed to read from world state: %v", err)
	}
	if tokenJSON == nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:Compute: token %s does not exist", id)
	}

	var token Token
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		return nil, err
	}

	tokenValid, err := isTokenValid(ctx, token)

	if err != nil {
		return nil, err
	}

	if !tokenValid {
		return nil, fmt.Errorf("ComputationTokenSmartContract:Compute: token is invalid")
	}

	nonce_byte, _ := ctx.GetStub().GetCreator()
	nonce := base64.StdEncoding.EncodeToString(nonce_byte)

	params := []string{token.Method, nonce}

	args := strings.Split(token.Arguments, ";")
	for _, arg := range args {
		params = append(params, strings.Split(arg, ":")[0])
	}

	log.Printf("Starting computing: %s %+q\n", token.Method, params)

	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("examplealgorithm", queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ComputationTokenSmartContract:Compute: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var ret Ret
	err = json.Unmarshal(response.Payload, &ret)
	if err != nil {
		return nil, err
	}

	token.Ret = ret

	tokenJSON, err = json.Marshal(token)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(string(decodedId), tokenJSON)
	if err != nil {
		return nil, err
	}

	log.Printf("Finished computing: %s %+q\n", token.Method, params)

	return &token, nil
}

// ReadToken returns the token entry stored in the world state with given id.
func ReadToken(ctx contractapi.TransactionContextInterface, id string) (*Token, error) {
	decodedId, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:ReadToken: failed to decode ID: %v", err)
	}
	tokenJSON, err := ctx.GetStub().GetState(string(decodedId))
	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:ReadToken: failed to read from world state: %v", err)
	}
	if tokenJSON == nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:ReadToken: token %s does not exist", id)
	}

	var token Token
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		return nil, err
	}

	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	if token.UserRequested != x509.Subject.ToRDNSequence().String() {
		return nil, fmt.Errorf("ComputationTokenSmartContract:ReadToken: set not a owner of this token")
	}

	return &token, nil
}
