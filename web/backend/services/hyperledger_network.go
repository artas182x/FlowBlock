package services

import (
	"log"
	"os"
	"path/filepath"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/models"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func GetNetwork(loginVals models.Login) (*gateway.Network, error) {
	ccpPath := os.Getenv("NETWORK_YAML")
	wallet := gateway.NewInMemoryWallet()
	identity := gateway.NewX509Identity(loginVals.MspID, loginVals.Certificate, loginVals.PrivateKey)
	wallet.Put("User", identity)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "User"),
	)
	if err != nil {
		log.Printf("Failed to connect to gateway: %v\n", err)
		return nil, err
	}
	defer gw.Close()

	network, err := gw.GetNetwork("medicalsystem")
	if err != nil {
		log.Printf("Failed to get network: %v\n", err)
		return nil, err
	}

	return network, nil

}

func EvaluateTransaction(login models.Login, chaincode string, contactName string, methodName string, args ...string) ([]byte, error) {

	network, err := GetNetwork(login)

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	contract := network.GetContractWithName(chaincode, contactName)
	log.Printf("--> Evaluate Transaction: %s %s %s\n", chaincode, contactName, methodName)
	result, err := contract.EvaluateTransaction(methodName, args...)
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		return nil, err
	}
	return result, nil
}

func SubmitTransaction(login models.Login, chaincode string, contactName string, methodName string, args ...string) ([]byte, error) {

	network, err := GetNetwork(login)

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	contract := network.GetContractWithName(chaincode, contactName)
	log.Printf("--> Evaluate Transaction: %s %s %s\n", chaincode, contactName, methodName)
	result, err := contract.SubmitTransaction(methodName, args...)
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		return nil, err
	}
	return result, nil
}
