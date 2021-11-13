package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

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

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

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

	network, err := gw.GetNetwork("medicalsystem")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	//contract := network.GetContractWithName("medicaldata", "PatientSmartContract")

	/*log.Println("--> Submit Transaction: InitLedger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))
	*/

	contract := network.GetContractWithName("medicaldata", "MedicalDataSmartContract")
	log.Println("--> Evaluate Transaction: GetAllEntriesAdmin")
	result, err := contract.EvaluateTransaction("GetAllEntriesAdmin")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	contract = network.GetContractWithName("medicaldata", "MedicalDataSmartContract")
	log.Println("--> Evaluate Transaction: GetPatientMedicalEntries")
	result, err = contract.EvaluateTransaction("GetPatientMedicalEntries", "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", "1573657134", "1636815534", "786cd6b55a09f12441051dbbb9f44393b55187eb4cb726391fa4b1b5aa307baa")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	contract = network.GetContractWithName("computationtoken", "ComputationTokenSmartContract")

	log.Println("--> Evaluate Transaction: RequestToken")
	result, err = contract.SubmitTransaction("RequestToken", "examplealghorytm", "ExampleAlghorytmSmartContract:AvgBloodPreasure", "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US;1573657134;1636815534")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	var token Token
	err = json.Unmarshal(result, &token)
	if err != nil {
		log.Fatalf("Failed decode json %v", err)
	}

	log.Println("--> Evaluate Transaction: Compute")
	result, err = contract.EvaluateTransaction("Compute", token.ID)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("============ application-golang ends ============")

}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
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
