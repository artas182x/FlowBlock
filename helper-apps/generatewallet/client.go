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

	userList := []string{"Admin@org1.example.com", "Admin@org3.example.com", "Admin@org4.example.com", "doctor1@org1.example.com", "doctor2@org1.example.com", "university1@org1.example.com", "patient1@org1.example.com",
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
	orgNum := getStringInBetween(org, "org", ".example.com")
	credPath := filepath.Join(
		".",
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
		log.Println("Certificate not found. Skipping")
		return nil
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

	orgMSP := "Org" + orgNum + "MSP"

	identity := gateway.NewX509Identity(orgMSP, string(cert), string(key))

	return wallet.Put(userOrg, identity)
}

func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	e += s + e - 1
	return str[s:e]
}
