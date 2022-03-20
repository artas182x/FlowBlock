package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Token struct {
	ID             string    `json:"ID"` //Must be string
	UserRequested  string    `json:userRequested`
	ChaincodeName  string    `json:chaincodeName`
	Method         string    `json:method`
	Arguments      string    `json:arguments`
	TimeRequested  time.Time `json:timeRequested`
	ExpirationTime time.Time `json.expirationTime`
}

func main() {
	log.Println("============ application-golang starts ============")

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	/*network, err := gw.GetNetwork("medicalsystem")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}*/

	//contract := network.GetContractWithName("medicaldata", "PatientSmartContract")

	/*log.Println("--> Submit Transaction: InitLedger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))
	*/

	b, err := ioutil.ReadFile("example.json")

	tokenValsStr := string(b)

	var tokenVals *[]tokenapi.Token

	log.Printf("%s", tokenValsStr)

	err = json.Unmarshal([]byte(tokenValsStr), tokenVals)

	/*contract := network.GetContractWithName("computationtoken", "ComputationTokenSmartContract")
	log.Println("--> Evaluate Transaction: ReadToken")
	result, err := contract.EvaluateTransaction("ReadToken", "AHRva2VucwB0b2tlbnMAQ049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUwAxNjQ3MTc2NzYxAEV4YW1wbGVBbGdvcml0aG1TbWFydENvbnRyYWN0OlhSYXlQbmV1bW9uaWFDYXNlcwAxNTc3ODgzODQwOzE1ODA0NzU5MDAA")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))
	*/

	/*
		contract = network.GetContractWithName("medicaldata", "MedicalDataSmartContract")
		log.Println("--> Evaluate Transaction: GetPatientMedicalEntries")
		result, err = contract.EvaluateTransaction("GetPatientMedicalEntries", "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", "1573657134", "1636815534", "786cd6b55a09f12441051dbbb9f44393b55187eb4cb726391fa4b1b5aa307baa")
		if err != nil {
			log.Fatalf("Failed to evaluate transaction: %v", err)
		}
		log.Println(string(result))

		contract = network.GetContractWithName("computationtoken", "ComputationTokenSmartContract")

		log.Println("--> Evaluate Transaction: RequestToken for AvgBloodPreasure")
		result, err = contract.SubmitTransaction("RequestToken", "examplealgorithm", "ExampleAlgorithmSmartContract:AvgBloodPreasure", "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US;1573657134;1636815534")
		if err != nil {
			log.Fatalf("Failed to evaluate transaction: %v", err)
		}
		log.Println(string(result))

		var token Token
		err = json.Unmarshal(result, &token)
		if err != nil {
			log.Fatalf("Failed decode json %v", err)
		}

		log.Println("--> Evaluate Transaction: Compute AvgBloodPreasure")
		result, err = contract.EvaluateTransaction("Compute", token.ID)
		if err != nil {
			log.Fatalf("Failed to evaluate transaction: %v", err)
		}
		log.Println(string(result))
	*/

	/*
		log.Println("--> Evaluate Transaction: RequestToken for LongRunningMethod")
		result, err = contract.SubmitTransaction("RequestToken", "examplealgorithm", "ExampleAlgorithmSmartContract:LongRunningMethod", "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US")
		if err != nil {
			log.Fatalf("Failed to evaluate transaction: %v", err)
		}
		log.Println(string(result))

		var token2 Token
		err = json.Unmarshal(result, &token2)
		if err != nil {
			log.Fatalf("Failed decode json %v", err)
		}

		log.Println("--> Evaluate Transaction: Compute LongRunningMethod")
		result, err = contract.EvaluateTransaction("Compute", token2.ID)
		if err != nil {
			log.Fatalf("Failed to evaluate transaction: %v", err)
		}
		log.Println(string(result))
	*/

	log.Println("============ application-golang ends ============")

}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"Admin@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
