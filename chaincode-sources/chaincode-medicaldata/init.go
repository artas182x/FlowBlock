package main

import (
	"log"

	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-medicaldata/medicaldatachaincode"
	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-medicaldata/patientchaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&patientchaincode.PatientSmartContract{}, &medicaldatachaincode.MedicalDataSmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
