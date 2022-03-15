package examplealgorithm

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/wcharczuk/go-chart/v2"
)

type ChartType int64

const (
	Bar ChartType = iota
	Donut
)

func (s *ExampleAlgorithmSmartContract) CreateBarChart(ctx contractapi.TransactionContextInterface, nonce string, tokenValsStr string, title string) (*tokenapi.Ret, error) {
	return s.CreateChart(ctx, nonce, tokenValsStr, title, Bar)
}

func (s *ExampleAlgorithmSmartContract) CreateDonutChart(ctx contractapi.TransactionContextInterface, nonce string, tokenValsStr string, title string) (*tokenapi.Ret, error) {
	return s.CreateChart(ctx, nonce, tokenValsStr, title, Donut)
}

// Creates bar chart on data from other computations
func (s *ExampleAlgorithmSmartContract) CreateChart(ctx contractapi.TransactionContextInterface, nonce string, tokenValsStr string, title string, chartType ChartType) (*tokenapi.Ret, error) {

	isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}
	if !isNonceValid {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Security error"), RetType: "string"}
		return &ret, nil
	}

	var tokenVals []tokenapi.Token

	err = json.Unmarshal([]byte(tokenValsStr), &tokenVals)

	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Can't parse tokens JSON"), RetType: "string"}
		return &ret, nil
	}

	var bars []chart.Value

	for _, token := range tokenVals {

		desc := token.Description
		var val float64

		if token.Ret.RetType == "int64" {
			intVar, err := strconv.Atoi(token.Ret.RetValue)
			if err != nil {
				ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain (cast from str->int64)"), RetType: "string"}
				return &ret, nil
			}
			val = float64(intVar)
		} else if token.Ret.RetType == "float32" {
			floatVar, err := strconv.ParseFloat(token.Ret.RetValue, 32)
			if err != nil {
				ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain (cast from str->float32)"), RetType: "string"}
				return &ret, nil
			}
			val = float64(floatVar)
		} else if token.Ret.RetType == "float64" {
			floatVar, err := strconv.ParseFloat(token.Ret.RetValue, 64)
			if err != nil {
				ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain (cast from str->float64)"), RetType: "string"}
				return &ret, nil
			}
			val = float64(floatVar)
		} else {
			ret := tokenapi.Ret{RetValue: fmt.Sprintf("Error: unsupported token value (id: %s)\n", token.ID), RetType: "string"}
			return &ret, nil
		}

		bars = append(bars, chart.Value{Value: val, Label: desc})
	}

	timestampRequested, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't get tx timestamp"), RetType: "string"}
		return &ret, nil
	}

	fileName := fmt.Sprintf("%s_%d.png", strings.ReplaceAll(title, " ", "_"), timestampRequested.GetSeconds())

	if chartType == Bar {
		graph := chart.BarChart{
			Title: title,
			Background: chart.Style{
				Padding: chart.Box{
					Top: 40,
				},
			},
			Height:   512,
			BarWidth: 60,
			Bars:     bars,
		}
		f, _ := os.Create(BASE_DIRECTORY + fileName)
		graph.Render(chart.PNG, f)
		f.Close()
	} else {
		graph := chart.DonutChart{
			Title:  title,
			Height: 512,
			Width:  512,
			Values: bars,
		}

		for i := range bars {
			bars[i].Label = fmt.Sprintf("%s - %f", bars[i].Label, bars[i].Value)
		}

		f, _ := os.Create(BASE_DIRECTORY + fileName)
		graph.Render(chart.PNG, f)
		f.Close()
	}

	fileId, err := tokenapi.UploadToS3(ctx, BASE_DIRECTORY+fileName)

	if err != nil {
		log.Printf("Upload to S3 failed: %v", err)
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: uploading to S3 failed"), RetType: "string"}
		return &ret, nil
	}

	ret := tokenapi.Ret{RetValue: fileId, RetType: "s3img"}

	return &ret, nil
}
