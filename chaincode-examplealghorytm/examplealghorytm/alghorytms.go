package examplealghorytm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ExampleAlghorytmSmartContract struct {
	contractapi.Contract
}

type Method struct {
	Name        string `json:"name"`
	Args        string `json:"args"`
	RetType     string `json:"retType"`
	Description string `json:"description"`
}

type MedicalEntry struct {
	ID                string    `json:"ID"` //Must be string
	PatientID         string    `json:"PatientID"`
	MedicalEntryName  string    `json:"MedicalEntryName"`
	MedicalEntryValue string    `json:"MedicalEntryValue"`
	DateAdded         time.Time `json:"DateAdded"`
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

var METHODS = []Method{
	{
		Name:        "ExampleAlghorytmSmartContract:AvgBloodPreasure",
		Args:        "patientID:string;startDateTimestamp:string;endDateTimestamp:string",
		RetType:     "float32",
		Description: "Calculates average value of blood preasure",
	},
	{
		Name:        "ExampleAlghorytmSmartContract:MaxHeartRate",
		Args:        "patientID:string;startDateTimestamp:string;endDateTimestamp:string",
		RetType:     "int64",
		Description: "Calculates maximum value of heart rate",
	},
	{
		Name:        "ExampleAlghorytmSmartContract:LongRunningMethod",
		Args:        "patientID:string",
		RetType:     "int64",
		Description: "Sleeps and returns current timestamp",
	},
}

func isTokenValid(ctx contractapi.TransactionContextInterface, token string, method string, args string) (bool, error) {

	params := []string{token, "ComputationTokenSmartContract:CheckTokenValidity", args, "examplealghorytm", "false"}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("computationtoken", queryArgs, "")

	if response.Status != shim.OK {
		return false, fmt.Errorf("ExampleAlghorytmSmartContract:AvgBloodPreasure: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var isValid bool
	err := json.Unmarshal(response.Payload, &isValid)
	if err != nil {
		return false, err
	}

	return isValid, nil
}

func (s *ExampleAlghorytmSmartContract) AvgBloodPreasure(ctx contractapi.TransactionContextInterface, token string, patientID string, startDateTimestamp string, endDateTimestamp string) (string, error) {

	argsStr := fmt.Sprintf("%s,%s,%s", patientID, startDateTimestamp, endDateTimestamp)
	isValid, err := isTokenValid(ctx, token, "AvgBloodPreasure", argsStr)
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", fmt.Errorf("Token %s is not valid", token)
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "SystolicBloodPreasure", startDateTimestamp, endDateTimestamp, token}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	fmt.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:AvgBloodPreasure: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var medicalEntries []MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		return "", err
	}

	ret := 0.0

	for _, element := range medicalEntries {
		intVar, err := strconv.Atoi(element.MedicalEntryValue)
		if err != nil {
			return "", err
		}
		ret = ret + float64(intVar)
	}
	ret = ret / float64(len(medicalEntries))

	// TODO put value somewhere

	return fmt.Sprint(ret), nil
}

func (s *ExampleAlghorytmSmartContract) MaxHeartRate(ctx contractapi.TransactionContextInterface, token string, patientID string, startDateTimestamp string, endDateTimestamp string) (string, error) {
	argsStr := fmt.Sprintf("%s,%s,%s", patientID, startDateTimestamp, endDateTimestamp)
	isValid, err := isTokenValid(ctx, token, "MaxHeartRate", argsStr)
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", fmt.Errorf("Token %s is not valid", token)
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "HeartRate", startDateTimestamp, endDateTimestamp, token}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	fmt.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:MaxHeartRate: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var medicalEntries []MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		return "", err
	}

	max := 0

	for _, element := range medicalEntries {
		intVar, err := strconv.Atoi(element.MedicalEntryValue)
		if err != nil {
			return "", err
		}
		if intVar > max {
			max = intVar
		}
	}

	// TODO put value somewhere

	return fmt.Sprint(max), nil
}

func (s *ExampleAlghorytmSmartContract) LongRunningMethod(ctx contractapi.TransactionContextInterface, token string, patientID string) (string, error) {
	isValid, err := isTokenValid(ctx, token, "LongRunningMethod", patientID)
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", fmt.Errorf("Token %s is not valid", token)
	}

	time.Sleep(60 * time.Second)

	// TODO put value somewhere

	return "0", nil
}

// Returns all available computation methods
func (s *ExampleAlghorytmSmartContract) ListAvailableMethods(ctx contractapi.TransactionContextInterface) ([]Method, error) {
	return METHODS, nil
}
