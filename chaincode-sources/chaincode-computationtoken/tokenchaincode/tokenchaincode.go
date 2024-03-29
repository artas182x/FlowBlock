package tokenchaincode

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const INDEX_NAME = "tokens"
const KEY_NAME = "tokens"

const TOKEN_VALID_MINUTES = 10

type ComputationTokenSmartContract struct {
	contractapi.Contract
}

// RequestToken issues a new token to the world state with given details.
func (s *ComputationTokenSmartContract) RequestToken(ctx contractapi.TransactionContextInterface, chaincodeName string, method string, arguments string, description string, directlyExecutable string) (*tokenapi.Token, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	userCN := x509.Subject.ToRDNSequence().String()

	log.Printf("[ComputationTokenSmartContract:RequestToken] requesting token: chaincode: %s method: %s args: %s desc: %s directlyExecutable: %s\n", chaincodeName, method, arguments, description, directlyExecutable)

	timestampRequested, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	timeRequested := time.Unix(int64(timestampRequested.GetSeconds()), int64(timestampRequested.GetNanos())).UTC()

	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, userCN, strconv.FormatInt(timestampRequested.GetSeconds(), 10), method, arguments})

	expirationTime := timeRequested.Add(time.Minute * time.Duration(TOKEN_VALID_MINUTES))

	err = cid.AssertAttributeValue(ctx.GetStub(), "RequestTokenRole", "1")
	if err != nil {
		return nil, fmt.Errorf("ComputationTokenSmartContract:RequestToken: No access to RequestToken")
	}

	params := []string{"ListAvailableMethods"}
	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

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

	var inputArgs []tokenapi.Argument
	err = json.Unmarshal([]byte(arguments), &inputArgs)
	if err != nil {
		return nil, err
	}

	var selectedMethod tokenapi.Method

	for _, element := range methods {
		if element.Name == method {
			if len(element.Args) == len(inputArgs) {
				found = true
				selectedMethod = element
				break
			}
		}
	}
	if !found {
		return nil, fmt.Errorf("ComputationTokenSmartContract:RequestToken: Method %s with arguments %s not found in %s", method, arguments, chaincodeName)
	}

	for i, arg := range selectedMethod.Args {
		inputArgs[i].Name = arg.Name
		inputArgs[i].Type = arg.Type
	}

	token := tokenapi.Token{
		ID:                 base64.URLEncoding.EncodeToString([]byte(id)),
		UserRequested:      x509.Subject.ToRDNSequence().String(),
		ChaincodeName:      chaincodeName,
		Method:             method,
		Arguments:          inputArgs,
		TimeRequested:      timeRequested,
		ExpirationTime:     expirationTime,
		Description:        description,
		DirectlyExecutable: strings.ToLower(directlyExecutable) == "true",
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	log.Printf("[ComputationTokenSmartContract:RequestToken] token submitted successfully")

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

	var token tokenapi.Token
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
func (s *ComputationTokenSmartContract) Compute(ctx contractapi.TransactionContextInterface, id string) (*tokenapi.Token, error) {
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

	var token tokenapi.Token
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

	log.Printf("Starting computing: %s %+q\n", token.Method, params)

	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

	for _, arg := range token.Arguments {

		if arg.Type == "tokenInputs" {
			subTokensIds := strings.Split(arg.Value, "|")

			var subTokens []tokenapi.Token

			for _, subTokenId := range subTokensIds {

				log.Printf("Subtoken compute: %s", subTokenId)

				tokenResponse, err := s.Compute(ctx, subTokenId)
				if err != nil {
					return nil, err
				}

				log.Printf("Subtoken compute success")

				subTokens = append(subTokens, *tokenResponse)
			}

			subTokensJSON, err := json.Marshal(subTokens)

			if err != nil {
				return nil, err
			}

			queryArgs = append(queryArgs, subTokensJSON)

		} else {
			queryArgs = append(queryArgs, []byte(arg.Value))
		}

	}

	response := ctx.GetStub().InvokeChaincode("examplealgorithm", queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ComputationTokenSmartContract:Compute: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var ret tokenapi.Ret
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

// Check if token in valid (only date and time and owner is checked)
func isTokenValid(ctx contractapi.TransactionContextInterface, token tokenapi.Token) (bool, error) {
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
