package examplealgorithm

import (
	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var METHODS = []tokenapi.Method{
	{
		Name:        "ExampleAlgorithmSmartContract:AvgBloodPreasure",
		Args:        []tokenapi.Argument{{Name: "patientID", Type: "string"}, {Name: "startDateTimestamp", Type: "ts"}, {Name: "endDateTimestamp", Type: "ts"}},
		RetType:     "float32",
		Description: "Calculates average value of blood preasure",
	},
	{
		Name:        "ExampleAlgorithmSmartContract:MaxHeartRate",
		Args:        []tokenapi.Argument{{Name: "patientID", Type: "string"}, {Name: "startDateTimestamp", Type: "ts"}, {Name: "endDateTimestamp", Type: "ts"}},
		RetType:     "int64",
		Description: "Calculates maximum value of heart rate",
	},
	{
		Name:        "ExampleAlgorithmSmartContract:LongRunningMethod",
		Args:        []tokenapi.Argument{{Name: "patientID", Type: "string"}},
		RetType:     "int64",
		Description: "Sleeps and returns 0 is there was no error",
	},
	{
		Name:        "ExampleAlgorithmSmartContract:PneumoniaImageClassification",
		Args:        []tokenapi.Argument{{Name: "medicalEntryId", Type: "string"}},
		RetType:     "string",
		Description: "Runs pneumonia image classification on specified medical entry id",
	},
	{
		Name:        "ExampleAlgorithmSmartContract:XRayPneumoniaCases",
		Args:        []tokenapi.Argument{{Name: "startDateTimestamp", Type: "ts"}, {Name: "endDateTimestamp", Type: "ts"}},
		RetType:     "int64",
		Description: "Calculates number of pneumonia cases over time based on XRay images",
	},
	{
		Name:        "ExampleAlgorithmSmartContract:CreateBarChart",
		Args:        []tokenapi.Argument{{Name: "tokenIds", Type: "tokenInputs"}, {Name: "title", Type: "string"}},
		RetType:     "s3img",
		Description: "Creates simple bar chart basign on values from other computation algorithms",
	},
	{
		Name:        "ExampleAlgorithmSmartContract:CreateDonutChart",
		Args:        []tokenapi.Argument{{Name: "tokenIds", Type: "tokenInputs"}, {Name: "title", Type: "string"}},
		RetType:     "s3img",
		Description: "Creates simple donut chart basign on values from other computation algorithms",
	},
}

// Returns all available computation methods
func (s *ExampleAlgorithmSmartContract) ListAvailableMethods(ctx contractapi.TransactionContextInterface) ([]tokenapi.Method, error) {
	return METHODS, nil
}
