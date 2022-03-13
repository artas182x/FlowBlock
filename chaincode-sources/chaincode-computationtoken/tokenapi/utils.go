package tokenapi

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Helper func that converts params string array to [][]byte array compatible with Hyperledger
func ParamsToHyperledgerArgs(params []string) [][]byte {
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}
	return queryArgs
}

// Check if token in valid (only date and time and owner is checked)
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
