package examplealgorithm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-computationtoken/tokenapi"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-medicaldata/medicaldatastructs"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ExampleAlghorytmSmartContract struct {
	contractapi.Contract
}

var METHODS = []tokenapi.Method{
	{
		Name:        "ExampleAlghorytmSmartContract:AvgBloodPreasure",
		Args:        "patientID:string;startDateTimestamp:ts;endDateTimestamp:ts",
		RetType:     "float32",
		Description: "Calculates average value of blood preasure",
	},
	{
		Name:        "ExampleAlghorytmSmartContract:MaxHeartRate",
		Args:        "patientID:string;startDateTimestamp:ts;endDateTimestamp:ts",
		RetType:     "int64",
		Description: "Calculates maximum value of heart rate",
	},
	{
		Name:        "ExampleAlghorytmSmartContract:LongRunningMethod",
		Args:        "patientID:string",
		RetType:     "int64",
		Description: "Sleeps and returns 0 is there was no error",
	},
}

func addMedicalValue(ctx contractapi.TransactionContextInterface, patientID string, medicalEntryName string, medicalEntryType string, medicalEntryValue string, nonce string) error {
	params := []string{"MedicalDataSmartContract:AddMedicalEntry", patientID, medicalEntryName, medicalEntryType, medicalEntryValue, nonce}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return fmt.Errorf("ExampleAlghorytmSmartContract:addMedicalValue: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	return nil
}

// Calculates average blood preasure for given patient and data range
func (s *ExampleAlghorytmSmartContract) AvgBloodPreasure(ctx contractapi.TransactionContextInterface, nonce string, patientID string, startDateTimestamp string, endDateTimestamp string) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		return nil, err
	}
	if !isNonceValid {
		return nil, fmt.Errorf("ExampleAlghorytmSmartContract:AvgBloodPreasure: Nonce is invalid")
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "SystolicBloodPreasure", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	fmt.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ExampleAlghorytmSmartContract:AvgBloodPreasure: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		return nil, err
	}

	val := 0.0

	for _, element := range medicalEntries {
		intVar, err := strconv.Atoi(element.MedicalEntryValue)
		if err != nil {
			return nil, err
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
func (s *ExampleAlghorytmSmartContract) MaxHeartRate(ctx contractapi.TransactionContextInterface, nonce string, patientID string, startDateTimestamp string, endDateTimestamp string) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		return nil, err
	}
	if !isNonceValid {
		return nil, fmt.Errorf("ExampleAlghorytmSmartContract:MaxHeartRate: Nonce is invalid")
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "HeartRate", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	fmt.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return nil, fmt.Errorf("ExampleAlghorytmSmartContract:MaxHeartRate: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		return nil, err
	}

	max := 0

	for _, element := range medicalEntries {
		intVar, err := strconv.Atoi(element.MedicalEntryValue)
		if err != nil {
			return nil, err
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
func (s *ExampleAlghorytmSmartContract) LongRunningMethod(ctx contractapi.TransactionContextInterface, nonce string, patientID string) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		return nil, err
	}
	if !isNonceValid {
		return nil, fmt.Errorf("ExampleAlghorytmSmartContract:LongRunningMethod: Nonce is invalid")
	}

	time.Sleep(60 * time.Second)

	ret := tokenapi.Ret{RetValue: "0", RetType: METHODS[2].RetType}

	return &ret, nil
}

// Returns all available computation methods
func (s *ExampleAlghorytmSmartContract) ListAvailableMethods(ctx contractapi.TransactionContextInterface) ([]tokenapi.Method, error) {
	return METHODS, nil
}
