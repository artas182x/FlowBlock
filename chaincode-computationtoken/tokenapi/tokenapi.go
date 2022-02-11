package tokenapi

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

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
	ID             string    `json:"ID"` //Must be string
	UserRequested  string    `json:userRequested`
	ChaincodeName  string    `json:chaincodeName`
	Method         string    `json:method`
	Arguments      string    `json:arguments`
	Ret            Ret       `json:"ret,omitempty" metadata:"ret,optional"`
	TimeRequested  time.Time `json:timeRequested`
	ExpirationTime time.Time `json.expirationTime`
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
