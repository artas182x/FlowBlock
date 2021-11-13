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

type Token struct {
	ID             string    `json:"ID"` //Must be string
	UserRequested  string    `json:userRequested`
	ChaincodeName  string    `json:chaincodeName`
	Method         string    `json:method`
	Arguments      string    `json:arguments`
	TimeRequested  time.Time `json:timeRequested`
	ExpirationTime time.Time `json.expirationTime`
}

func IsNonceValid(ctx contractapi.TransactionContextInterface, nonceStr string) (bool, error) {
	creatorByte, err := ctx.GetStub().GetCreator()
	if err != nil {
		return false, err
	}

	creator := base64.StdEncoding.EncodeToString(creatorByte)

	fmt.Printf("tokenapi:isNonceValid: Comparing GetCreator(): %s vs nonce: %s\n", creator, nonceStr)
	return creator == nonceStr, nil
}
