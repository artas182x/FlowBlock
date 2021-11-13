package main

import (
	"log"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-examplealghorytm/examplealghorytm"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&examplealghorytm.ExampleAlghorytmSmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
