package examplealghorytm

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-medicaldata/medicaldatastructs"
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

func isNonceValid(ctx contractapi.TransactionContextInterface, nonceStr string) (bool, error) {
	creatorByte, err := ctx.GetStub().GetCreator()
	if err != nil {
		return false, err
	}

	creator := base64.StdEncoding.EncodeToString(creatorByte)

	fmt.Printf("ExampleAlghorytmSmartContract:isNonceValid: Comparing GetCreator(): %s vs nonce: %s\n", creator, nonceStr)
	return creator == nonceStr, nil
}

func (s *ExampleAlghorytmSmartContract) AvgBloodPreasure(ctx contractapi.TransactionContextInterface, nonce string, patientID string, startDateTimestamp string, endDateTimestamp string) (string, error) {

	isNonceValid, err := isNonceValid(ctx, nonce)
	if err != nil {
		return "", err
	}
	if !isNonceValid {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:AvgBloodPreasure: Nonce is invalid")
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "SystolicBloodPreasure", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	fmt.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:AvgBloodPreasure: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
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

func (s *ExampleAlghorytmSmartContract) MaxHeartRate(ctx contractapi.TransactionContextInterface, nonce string, patientID string, startDateTimestamp string, endDateTimestamp string) (string, error) {

	isNonceValid, err := isNonceValid(ctx, nonce)
	if err != nil {
		return "", err
	}
	if !isNonceValid {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:MaxHeartRate: Nonce is invalid")
	}

	params := []string{"MedicalDataSmartContract:GetPatientMedicalEntries", patientID, "HeartRate", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	fmt.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:MaxHeartRate: failed to query chaincode. Status: %d Payload: %s Message: %s", response.Status, response.Payload, response.Message)
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
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

func (s *ExampleAlghorytmSmartContract) LongRunningMethod(ctx contractapi.TransactionContextInterface, nonce string, token string, patientID string) (string, error) {

	isNonceValid, err := isNonceValid(ctx, nonce)
	if err != nil {
		return "", err
	}
	if !isNonceValid {
		return "", fmt.Errorf("ExampleAlghorytmSmartContract:MaxHeartRate: Nonce is invalid")
	}

	time.Sleep(60 * time.Second)

	// TODO put value somewhere

	return "0", nil
}

// Returns all available computation methods
func (s *ExampleAlghorytmSmartContract) ListAvailableMethods(ctx contractapi.TransactionContextInterface) ([]Method, error) {
	return METHODS, nil
}
