package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	log.Println("============ application-golang starts ============")

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	userList := []string{"Admin@org1.example.com", "doctor1@org1.example.com", "doctor2@org1.example.com", "doctor3@org1.example.com", "patient1@org1.example.com",
		"patient2@org1.example.com", "patient3@org1.example.com", "Admin@org2.example.com", "doctor11@org2.example.com", "patient11@org2.example.com"}

	for _, user := range userList {
		err = populateWallet(wallet, user)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents for user %s: %v", user, err)
		}
	}

}

func populateWallet(wallet *gateway.Wallet, userOrg string) error {
	log.Printf("============ Populating wallet for user %s ============", userOrg)
	org := strings.Split(userOrg, "@")[1]
	credPath := filepath.Join(
		"/app",
		"peerOrganizations",
		org,
		"users",
		userOrg,
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

	return wallet.Put(userOrg, identity)
}
