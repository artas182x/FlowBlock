package tokenchaincode

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const TOKEN_VALID_MINUTES = 5

type ComputationTokenSmartContract struct {
	contractapi.Contract
}

type Token struct {
	ID             string    `json:"ID"` //Must be string
	UserRequested  string    `json:userRequested`
	ChaincodeName  string    `json:chaincodeName`
	Method         string    `json:method`
	Arguments      string    `json:arguments`
	TimeRequested  time.Time `json:timeRequested`
	ExpirationTime time.Time `json.expirationTime`
}

type Method struct {
	Name        string `json:"name"`
	Args        string `json:"args"`
	RetType     string `json:"retType"`
	Description string `json:"description"`
}

// AddEntry issues a new entry to the world state with given details.
func (s *ComputationTokenSmartContract) RequestToken(ctx contractapi.TransactionContextInterface, chaincodeName string, method string, arguments string) (*Token, error) {
	id := ctx.GetStub().GetTxID()
	timestampRequested, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	timeRequested := time.Unix(int64(timestampRequested.GetSeconds()), int64(timestampRequested.GetNanos())).UTC()

	expirationTime := timeRequested.Add(time.Minute * time.Duration(TOKEN_VALID_MINUTES))

	x509, _ := cid.GetX509Certificate(ctx.GetStub())

	// TODO Permission logic here

	params := []string{"ListAvailableMethods"}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode(chaincodeName, queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ComputationTokenSmartContract:RequestToken: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var methods []Method
	err = json.Unmarshal(response.Payload, &methods)
	found := false
	if err != nil {
		return nil, err
	}
	for _, element := range methods {
		if element.Name == method {
			if len(strings.Split(element.Args, ";")) == len(strings.Split(arguments, ";")) {
				found = true
				break
			}
		}
	}
	if !found {
		return nil, fmt.Errorf("Method %s with arguments %s not found in %s", method, arguments, chaincodeName)
	}

	token := Token{
		ID:             id,
		UserRequested:  x509.Subject.ToRDNSequence().String(),
		ChaincodeName:  chaincodeName,
		Method:         method,
		Arguments:      arguments,
		TimeRequested:  timeRequested,
		ExpirationTime: expirationTime,
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(id, tokenJSON)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// ReadToken returns the token entry stored in the world state with given id.
func (s *ComputationTokenSmartContract) ReadToken(ctx contractapi.TransactionContextInterface, id string) (*Token, error) {
	tokenJSON, err := ctx.GetStub().GetState(id)
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

// Compute method launches alghoritm provided it has correct token
func (s *ComputationTokenSmartContract) Compute(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	tokenJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", fmt.Errorf("ComputationTokenSmartContract:Compute: failed to read from world state: %v", err)
	}
	if tokenJSON == nil {
		return "", fmt.Errorf("ComputationTokenSmartContract:Compute: token %s does not exist", id)
	}

	var token Token
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		return "", err
	}

	tokenValid, err := isTokenValid(ctx, token)

	if err != nil {
		return "", err
	}

	if !tokenValid {
		return "", fmt.Errorf("ComputationTokenSmartContract:Compute: token is invalid")
	}

	nonce_byte, _ := ctx.GetStub().GetCreator()
	nonce := base64.StdEncoding.EncodeToString(nonce_byte)

	params := []string{token.Method, nonce}

	args := strings.Split(token.Arguments, ";")
	params = append(params, args...)

	fmt.Printf("Starting computing: %+q\n", params)

	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("examplealghorytm", queryArgs, "")

	if response.Status != shim.OK {
		return "", fmt.Errorf("ComputationTokenSmartContract:Compute: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	return string(response.Payload), nil

}

// GetAllEntriesAdmin returns all roken entries found in world state. Only admin can execute this
func (s *ComputationTokenSmartContract) GetAllEntriesAdmin(ctx contractapi.TransactionContextInterface) ([]*Token, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	err := cid.AssertAttributeValue(ctx.GetStub(), "hf.Type", "admin")

	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:GetAllEntriesAdmin: only admin can do this. Current user: %s", x509.Subject.ToRDNSequence().String())
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var tokens []*Token
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var token Token
		err = json.Unmarshal(queryResponse.Value, &token)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, &token)
	}

	return tokens, nil
}

func isTokenValid(ctx contractapi.TransactionContextInterface, token Token) (bool, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	if token.UserRequested != x509.Subject.ToRDNSequence().String() {
		return false, fmt.Errorf("ComputationTokenSmartContract:isTokenValid: user not owner of this token")
	}

	currTimestamp, err := ctx.GetStub().GetTxTimestamp()

	if err != nil {
		return false, err
	}

	currTime := time.Unix(int64(currTimestamp.GetSeconds()), int64(currTimestamp.GetNanos())).UTC()
	timeDiff := currTime.Sub(token.TimeRequested)

	if timeDiff.Minutes() > TOKEN_VALID_MINUTES {
		return false, fmt.Errorf("ComputationTokenSmartContract:isTokenValid: token expired")
	}

	return true, nil
}

func (s *ComputationTokenSmartContract) CheckTokenValidity(ctx contractapi.TransactionContextInterface, tokenID string, method string, arguments string, chaincodeName string, basicCheck string) (bool, error) {

	fmt.Printf("ComputationTokenSmartContract:CheckTokenValidity: Veryfing %s token value for: %s %s (%s) Basic check: %s\n", tokenID, method, arguments, chaincodeName, basicCheck)

	tokenJSON, err := ctx.GetStub().GetState(tokenID)
	if err != nil {
		return false, fmt.Errorf("ComputationTokenSmartContract:checkTokenValidity: failed to read from world state: %v", err)
	}
	if tokenJSON == nil {
		return false, fmt.Errorf("ComputationTokenSmartContract:checkTokenValidity: token %s does not exist", tokenID)
	}

	var token Token
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		return false, err
	}

	tokenValid, err := isTokenValid(ctx, token)

	if err != nil {
		return false, err
	}

	fmt.Printf("ComputationTokenSmartContract:CheckTokenValidity: isTokenValid: %t\n", tokenValid)

	if !tokenValid {
		return false, nil
	}

	basicCheckBool, _ := strconv.ParseBool(basicCheck)

	if !basicCheckBool {
		if method != token.Method || arguments != token.Arguments || chaincodeName != token.ChaincodeName {
			fmt.Printf("ComputationTokenSmartContract:CheckTokenValidity: Result: %t\n", false)
			return false, nil
		}
	}

	fmt.Printf("ComputationTokenSmartContract:CheckTokenValidity: Result: %t\n", true)

	return true, nil
}
