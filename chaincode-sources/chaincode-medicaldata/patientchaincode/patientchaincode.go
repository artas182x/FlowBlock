package patientchaincode

import (
	"encoding/json"
	"fmt"

	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-medicaldata/medicaldatastructs"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PatientSmartContract struct {
	contractapi.Contract
}

const INDEX_NAME = "patientData"
const KEY_NAME = "patientData"

// AddEntry issues a new entry to the world state with given details. Assuming that patient is adding himself to the system
func (s *PatientSmartContract) AddPatientEntry(ctx contractapi.TransactionContextInterface, orgReadAllowed []string, usersReadAllowed []string, orgComputationAllowed []string, usersComputationAllowed []string, orgWriteAllowed []string, usersWriteAllowed []string) error {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, x509.Subject.ToRDNSequence().String()})
	patientEntry := medicaldatastructs.PatientEntry{
		ID:                      id,
		OrgReadAllowed:          orgReadAllowed,
		UsersReadAllowed:        usersReadAllowed,
		OrgComputationAllowed:   orgComputationAllowed,
		UsersComputationAllowed: usersComputationAllowed,
		OrgWriteAllowed:         orgWriteAllowed,
		UsersWriteAllowed:       usersWriteAllowed,
	}
	patientEntryJSON, err := json.Marshal(patientEntry)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, patientEntryJSON)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ReadPatientEntry returns the patient entry stored in the world state with given id.
func (s *PatientSmartContract) ReadPatientEntry(ctx contractapi.TransactionContextInterface, patientName string) (*medicaldatastructs.PatientEntry, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, patientName})
	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("PatientSmartContract:ReadPatientEntry: failed to read from world state: %v", err)
	}
	if patientEntryJSON == nil {
		return nil, fmt.Errorf("PatientSmartContract:ReadPatientEntry: the patient entry %s does not exist", id)
	}

	var patientEntry medicaldatastructs.PatientEntry
	err = json.Unmarshal(patientEntryJSON, &patientEntry)
	if err != nil {
		return nil, err
	}

	if x509.Subject.ToRDNSequence().String() != patientName && !stringInSlice(x509.Subject.ToRDNSequence().String(), patientEntry.UsersReadAllowed) && !stringInSlice(x509.Issuer.Organization[0], patientEntry.OrgReadAllowed) {
		return nil, fmt.Errorf("PatientSmartContract:ReadPatientEntry: user not allowed to read this value")
	}

	//We do not allow computation on personal data like name or surname. Only access to medical data

	return &patientEntry, nil
}

// UpdatePatientEntry updates an existing patient entry in the world state with provided parameters.
func (s *PatientSmartContract) UpdatePatientEntry(ctx contractapi.TransactionContextInterface, orgReadAllowed []string, usersReadAllowed []string, orgComputationAllowed []string, usersComputationAllowed []string, orgWriteAllowed []string, usersWriteAllowed []string) error {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	id := x509.Subject.CommonName
	exists, err := s.EntryExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("PatientSmartContract:UpdatePatientEntry: the patientEntry %s does not exist", id)
	}

	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("PatientSmartContract:UpdatePatientEntry: failed to read from world state: %v", err)
	}
	if patientEntryJSON == nil {
		return fmt.Errorf("PatientSmartContract:UpdatePatientEntry: the patient entry %s does not exist", id)
	}

	var patientEntry medicaldatastructs.PatientEntry
	err = json.Unmarshal(patientEntryJSON, &patientEntry)
	if err != nil {
		return err
	}

	// overwriting original patientEntry with new
	patientEntry = medicaldatastructs.PatientEntry{
		ID:                      id,
		OrgReadAllowed:          orgReadAllowed,
		UsersReadAllowed:        usersReadAllowed,
		OrgComputationAllowed:   orgComputationAllowed,
		UsersComputationAllowed: usersComputationAllowed,
		OrgWriteAllowed:         orgWriteAllowed,
		UsersWriteAllowed:       usersWriteAllowed,
	}
	patientEntryJSON, err = json.Marshal(patientEntry)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, patientEntryJSON)
}

// DeletePatientEntry deletes an given patientEntry from the world state.
func (s *PatientSmartContract) DeletePatientEntry(ctx contractapi.TransactionContextInterface) error {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	id := x509.Subject.CommonName
	exists, err := s.EntryExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("PatientSmartContract:DeletePatientEntry: the patientEntry %s does not exist", id)
	}

	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("PatientSmartContract:DeletePatientEntry: failed to read from world state: %v", err)
	}
	if patientEntryJSON == nil {
		return fmt.Errorf("PatientSmartContract:DeletePatientEntry: the patient entry %s does not exist", id)
	}

	var patientEntry medicaldatastructs.PatientEntry
	err = json.Unmarshal(patientEntryJSON, &patientEntry)
	if err != nil {
		return err
	}

	return ctx.GetStub().DelState(id)
}

// EntryExists returns true when patientEntry with given ID exists in world state
func (s *PatientSmartContract) EntryExists(ctx contractapi.TransactionContextInterface, patientID string) (bool, error) {
	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, patientID})
	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("PatientSmartContract:EntryExists: failed to read from world state: %v", err)
	}

	return patientEntryJSON != nil, nil
}

// CanRead returns whether user or org has permission to read values
func (s *PatientSmartContract) CanRead(ctx contractapi.TransactionContextInterface, patientID string) (bool, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())

	// Patient always can read own values
	if patientID == x509.Subject.ToRDNSequence().String() {
		return true, nil
	} else {
		err := cid.AssertAttributeValue(ctx.GetStub(), "ReadOthersData", "1")
		if err != nil {
			return false, nil
		}
	}

	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, patientID})
	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("PatientSmartContract:CanRead: failed to read from world state: %v", err)
	}
	if patientEntryJSON == nil {
		return false, fmt.Errorf("PatientSmartContract:CanRead: the patient entry %s does not exist", patientID)
	}

	var patientEntry medicaldatastructs.PatientEntry
	err = json.Unmarshal(patientEntryJSON, &patientEntry)
	if err != nil {
		return false, err
	}

	if !stringInSlice(x509.Subject.ToRDNSequence().String(), patientEntry.UsersReadAllowed) && !stringInSlice(x509.Issuer.Organization[0], patientEntry.OrgReadAllowed) {
		return false, nil
	}

	return true, nil
}

// CanRead returns whether user or org has permission to write values
func (s *PatientSmartContract) CanWrite(ctx contractapi.TransactionContextInterface, patientID string) (bool, error) {

	err := cid.AssertAttributeValue(ctx.GetStub(), "ReadOthersData", "1")
	if err != nil {
		return false, nil
	}

	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, patientID})
	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("PatientSmartContract:CanWrite: failed to read from world state: %v", err)
	}
	if patientEntryJSON == nil {
		return false, fmt.Errorf("PatientSmartContract:CanWrite: the patient entry %s does not exist", patientID)
	}

	var patientEntry medicaldatastructs.PatientEntry
	err = json.Unmarshal(patientEntryJSON, &patientEntry)
	if err != nil {
		return false, err
	}

	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	if !stringInSlice(x509.Subject.ToRDNSequence().String(), patientEntry.UsersWriteAllowed) && !stringInSlice(x509.Issuer.Organization[0], patientEntry.OrgWriteAllowed) {
		return false, nil
	}

	return true, nil
}

// CanCompute returns whether user or org has permission to read values
func (s *PatientSmartContract) CanCompute(ctx contractapi.TransactionContextInterface, patientID string) (bool, error) {
	err := cid.AssertAttributeValue(ctx.GetStub(), "RequestTokenRole", "1")
	if err != nil {
		return false, nil
	}
	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, patientID})
	patientEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("PatientSmartContract:CanCompute: failed to read from world state: %v", err)
	}
	if patientEntryJSON == nil {
		return false, fmt.Errorf("PatientSmartContract:CanCompute: the patient entry %s does not exist", patientID)
	}

	var patientEntry medicaldatastructs.PatientEntry
	err = json.Unmarshal(patientEntryJSON, &patientEntry)
	if err != nil {
		return false, err
	}

	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	if !stringInSlice(x509.Subject.ToRDNSequence().String(), patientEntry.OrgComputationAllowed) && !stringInSlice(x509.Issuer.Organization[0], patientEntry.OrgComputationAllowed) {
		return false, nil
	}

	return true, nil
}

// GetAllEntriesAdmin returns all medical entries found in world state. Only admin can execute this
func (s *PatientSmartContract) GetAllEntriesAdmin(ctx contractapi.TransactionContextInterface) ([]*medicaldatastructs.PatientEntry, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	err := cid.AssertAttributeValue(ctx.GetStub(), "hf.Type", "admin")

	if err != nil {
		return nil, fmt.Errorf("PatientSmartContract:GetAllEntriesAdmin: only admin can do this. Current user: %s", x509.Subject.ToRDNSequence().String())
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var medicalEntries []*medicaldatastructs.PatientEntry
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var patientEntry medicaldatastructs.PatientEntry
		err = json.Unmarshal(queryResponse.Value, &patientEntry)
		if err != nil {
			return nil, err
		}

		medicalEntries = append(medicalEntries, &patientEntry)
	}

	return medicalEntries, nil
}
