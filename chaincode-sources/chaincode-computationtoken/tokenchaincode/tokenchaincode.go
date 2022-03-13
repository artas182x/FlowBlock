package tokenchaincode

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const INDEX_NAME = "tokens"
const KEY_NAME = "tokens"

type ComputationTokenSmartContract struct {
	contractapi.Contract
}

// RequestToken issues a new token to the world state with given details.
func (s *ComputationTokenSmartContract) RequestToken(ctx contractapi.TransactionContextInterface, chaincodeName string, method string, arguments string, description string, directlyExecutable string) (*tokenapi.Token, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	userCN := x509.Subject.ToRDNSequence().String()

	fmt.Printf("[ComputationTokenSmartContract:RequestToken] requesting token: chaincode: %s method: %s args: %s desc: %s directlyExecutable: %s\n", chaincodeName, method, arguments, description, directlyExecutable)

	timestampRequested, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	timeRequested := time.Unix(int64(timestampRequested.GetSeconds()), int64(timestampRequested.GetNanos())).UTC()

	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, userCN, strconv.FormatInt(timestampRequested.GetSeconds(), 10), method, arguments})

	expirationTime := timeRequested.Add(time.Minute * time.Duration(tokenapi.TOKEN_VALID_MINUTES))

	err = cid.AssertAttributeValue(ctx.GetStub(), "RequestTokenRole", "1")
	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:RequestToken: No access to RequestToken")
	}

	params := []string{"ListAvailableMethods"}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode(chaincodeName, queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ComputationTokenSmartContract:RequestToken: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var methods []tokenapi.Method
	err = json.Unmarshal(response.Payload, &methods)
	found := false
	if err != nil {
		return nil, err
	}

	var selectedMethod tokenapi.Method

	for _, element := range methods {
		if element.Name == method {
			if len(strings.Split(element.Args, ";")) == len(strings.Split(arguments, ";")) {
				found = true
				selectedMethod = element
				break
			}
		}
	}
	if !found {
		return nil, fmt.Errorf("ComputationTokenSmartContract:RequestToken: Method %s with arguments %s not found in %s", method, arguments, chaincodeName)
	}

	var mergedArgsString string = ""
	argsArray := strings.Split(arguments, ";")

	for num, arg := range strings.Split(selectedMethod.Args, ";") {
		argsSplit := strings.Split(arg, ":")
		mergedArgsString += fmt.Sprintf("%s:%s:%s", argsArray[num], argsSplit[0], argsSplit[1])
		if num < len(argsArray)-1 {
			mergedArgsString += ";"
		}
	}

	token := tokenapi.Token{
		ID:                 base64.URLEncoding.EncodeToString([]byte(id)),
		UserRequested:      x509.Subject.ToRDNSequence().String(),
		ChaincodeName:      chaincodeName,
		Method:             method,
		Arguments:          mergedArgsString,
		TimeRequested:      timeRequested,
		ExpirationTime:     expirationTime,
		Description:        description,
		DirectlyExecutable: strings.ToLower(directlyExecutable) == "true",
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	fmt.Printf("[ComputationTokenSmartContract:RequestToken] token submitted successfully")

	err = ctx.GetStub().PutState(id, tokenJSON)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// GetAvailableMethods returns method available by given chaincode
func (s *ComputationTokenSmartContract) GetAvailableMethods(ctx contractapi.TransactionContextInterface, chaincodeName string) ([]tokenapi.Method, error) {
	params := []string{"ListAvailableMethods"}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode(chaincodeName, queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ComputationTokenSmartContract:GetAvailableMethods: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var methods []tokenapi.Method
	err := json.Unmarshal(response.Payload, &methods)
	if err != nil {
		return nil, err
	}

	return methods, nil
}

// ReadUserTokens returns all tokens that belong to current user
func (s *ComputationTokenSmartContract) ReadUserTokens(ctx contractapi.TransactionContextInterface) ([]*tokenapi.Token, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	id := []string{KEY_NAME, x509.Subject.ToRDNSequence().String()}

	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(INDEX_NAME, id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var tokens []*tokenapi.Token
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var token tokenapi.Token
		err = json.Unmarshal(queryResponse.Value, &token)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, &token)
	}

	return tokens, nil
}

// ReadToken returns the token entry stored in the world state with given id.
func (s *ComputationTokenSmartContract) ReadToken(ctx contractapi.TransactionContextInterface, id string) (*tokenapi.Token, error) {
	return tokenapi.ReadToken(ctx, id)
}

// Compute method launches alghoritm provided it has correct token
func (s *ComputationTokenSmartContract) Compute(ctx contractapi.TransactionContextInterface, id string) (*tokenapi.Token, error) {
	return tokenapi.Compute(ctx, id)
}

// GetAllEntriesAdmin returns all roken entries found in world state. Only admin can execute this
func (s *ComputationTokenSmartContract) GetAllEntriesAdmin(ctx contractapi.TransactionContextInterface) ([]*tokenapi.Token, error) {
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

	var tokens []*tokenapi.Token
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var token tokenapi.Token
		err = json.Unmarshal(queryResponse.Value, &token)
		if err != nil {
			return nil, err
		}

		token.ID = base64.URLEncoding.EncodeToString([]byte(token.ID))

		tokens = append(tokens, &token)
	}

	return tokens, nil
}
