package examplealgorithm

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-medicaldata/medicaldatastructs"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ExampleAlgorithmSmartContract struct {
	contractapi.Contract
}

const BASE_DIRECTORY = "/tmp/exalg/"

func addMedicalValue(ctx contractapi.TransactionContextInterface, patientID string, medicalEntryName string, medicalEntryType string, medicalEntryValue string, nonce string) error {
	params := []string{"MedicalDataSmartContract:AddMedicalEntry", patientID, medicalEntryName, medicalEntryType, medicalEntryValue, nonce}
	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return fmt.Errorf("ExampleAlgorithmSmartContract:addMedicalValue: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	return nil
}

// Calculates average blood preasure for given patient and data range
func (s *ExampleAlgorithmSmartContract) AvgBloodPreasure(ctx contractapi.TransactionContextInterface, nonce string, patientID string, startDateTimestamp string, endDateTimestamp string) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}
	if !isNonceValid {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "SystolicBloodPreasure", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

	log.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't read data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	val := 0.0

	for _, element := range medicalEntries {
		intVar, err := strconv.Atoi(element.MedicalEntryValue)
		if err != nil {
			ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain"), RetType: "string"}
			return &ret, nil
		}
		val = val + float64(intVar)
	}
	val = val / float64(len(medicalEntries))

	retVal := fmt.Sprint(val)
	retType := METHODS[0].RetType

	ret := tokenapi.Ret{RetValue: retVal, RetType: retType}

	err = addMedicalValue(ctx, patientID, "AvgBloodPreasure", retType, retVal, nonce)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// Calculates maximum heart rate for given patient and data range
func (s *ExampleAlgorithmSmartContract) MaxHeartRate(ctx contractapi.TransactionContextInterface, nonce string, patientID string, startDateTimestamp string, endDateTimestamp string) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}
	if !isNonceValid {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "HeartRate", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

	log.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't read data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	max := 0

	for _, element := range medicalEntries {
		intVar, err := strconv.Atoi(element.MedicalEntryValue)
		if err != nil {
			ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain"), RetType: "string"}
			return &ret, nil
		}
		if intVar > max {
			max = intVar
		}
	}

	retVal := fmt.Sprint(max)
	retType := METHODS[0].RetType

	ret := tokenapi.Ret{RetValue: retVal, RetType: retType}

	err = addMedicalValue(ctx, patientID, "MaxHeartRate", retType, retVal, nonce)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// Some long running algorithm as an example
func (s *ExampleAlgorithmSmartContract) LongRunningMethod(ctx contractapi.TransactionContextInterface, nonce string, patientID string) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}
	if !isNonceValid {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}

	time.Sleep(60 * time.Second)

	ret := tokenapi.Ret{RetValue: "0", RetType: METHODS[2].RetType}

	return &ret, nil
}
