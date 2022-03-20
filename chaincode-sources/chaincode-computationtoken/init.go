package main

import (
	"log"

	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-computationtoken/tokenchaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&tokenchaincode.ComputationTokenSmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
