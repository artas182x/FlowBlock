package patientchaincode

import (
	"encoding/json"
	"fmt"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-medicaldata/medicaldatastructs"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const PATIENTS_NUM = 20

// Example data to fill ledger for testing purposes. Should be removed in production.
func (s *PatientSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	id1, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id2, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient2,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id3, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id4, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient4,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id5, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient5,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id6, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient6,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id7, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient7,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id8, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient8,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id9, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient9,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id10, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient10,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id11, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient11,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id12, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient12,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id13, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient13,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id14, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient14,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id15, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient15,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id16, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient16,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id17, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient17,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id18, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient18,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id19, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient19,OU=client,O=Hyperledger,ST=North Carolina,C=US"})
	id20, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient20,OU=client,O=Hyperledger,ST=North Carolina,C=US"})

	entries := []medicaldatastructs.PatientEntry{
		{
			ID:                      id1,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id2,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id3,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id4,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id5,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id6,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id7,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id8,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id9,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id10,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id11,
			OrgReadAllowed:          []string{"org1.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id12,
			OrgReadAllowed:          []string{"org1.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id13,
			OrgReadAllowed:          []string{"org1.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id14,
			OrgReadAllowed:          []string{"org1.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id15,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id16,
			OrgReadAllowed:          []string{"org1.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com", "org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id17,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org2admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id18,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id19,
			OrgReadAllowed:          []string{"org1.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org1.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
		{
			ID:                      id20,
			OrgReadAllowed:          []string{"org2.example.com"},
			UsersReadAllowed:        []string{"CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US"},
			OrgComputationAllowed:   []string{"org2.example.com"},
			UsersComputationAllowed: []string{},
			OrgWriteAllowed:         []string{},
			UsersWriteAllowed:       []string{},
		},
	}

	for _, entry := range entries {
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(entry.ID, entryJSON)
		if err != nil {
			return fmt.Errorf("PatientSmartContract:InitLedger: failed to put to world state. %v", err)
		}
	}

	return nil
}
