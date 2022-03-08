package main

import (
	"log"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-examplealgorithm/examplealgorithm"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&examplealgorithm.ExampleAlgorithmSmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
