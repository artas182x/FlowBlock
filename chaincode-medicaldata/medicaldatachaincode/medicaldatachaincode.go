package medicaldatachaincode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-computationtoken/tokenapi"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-medicaldata/medicaldatastructs"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-medicaldata/patientchaincode"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MedicalDataSmartContract struct {
	contractapi.Contract
}

const INDEX_NAME = "medicalData"
const KEY_NAME = "medicalData"

// Example data to fill ledger for testing purposes. Should be removed in production.
func (s *MedicalDataSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	id1, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", ctx.GetStub().GetTxID() + "1"})
	id2, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", ctx.GetStub().GetTxID() + "2"})
	id3, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", ctx.GetStub().GetTxID() + "3"})
	id4, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US", "HeartRate", ctx.GetStub().GetTxID() + "4"})
	id5, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient2,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", ctx.GetStub().GetTxID() + "5"})
	id6, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient2,OU=client,O=Hyperledger,ST=North Carolina,C=US", "SystolicBloodPreasure", ctx.GetStub().GetTxID() + "6"})
	id7, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "HeartRate", ctx.GetStub().GetTxID() + "7"})
	id8, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "8"})
	id9, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "9"})
	id10, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "10"})
	id11, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "11"})
	id12, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "12"})
	id13, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "13"})
	id14, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US", "ChestXRay", ctx.GetStub().GetTxID() + "14"})

	entries := []medicaldatastructs.MedicalEntry{
		{
			ID:                id1,
			PatientID:         "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "SystolicBloodPreasure",
			MedicalEntryValue: "130",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.January, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id2,
			PatientID:         "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "SystolicBloodPreasure",
			MedicalEntryValue: "150",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.January, 17, 12, 30, 30, 34, time.UTC),
		},
		{
			ID:                id3,
			PatientID:         "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "SystolicBloodPreasure",
			MedicalEntryValue: "143",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.February, 15, 16, 12, 15, 34, time.UTC),
		},
		{
			ID:                id4,
			PatientID:         "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "HeartRate",
			MedicalEntryValue: "80",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.February, 15, 16, 12, 15, 34, time.UTC),
		},
		{
			ID:                id5,
			PatientID:         "CN=patient2,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "SystolicBloodPreasure",
			MedicalEntryValue: "110",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.January, 11, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id6,
			PatientID:         "CN=patient2,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "SystolicBloodPreasure",
			MedicalEntryValue: "113",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.July, 25, 9, 12, 15, 34, time.UTC),
		},
		{
			ID:                id7,
			PatientID:         "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "HeartRate",
			MedicalEntryValue: "150",
			MedicalEntryType:  "int64",
			DateAdded:         time.Date(2021, time.July, 26, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id8,
			PatientID:         "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "pneumonia.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2021, time.December, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id9,
			PatientID:         "CN=patient1,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "normal.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2022, time.January, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id10,
			PatientID:         "CN=patient2,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "pneumonia2.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2021, time.November, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id11,
			PatientID:         "CN=patient3,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "normal2.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2021, time.August, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id12,
			PatientID:         "CN=patient4,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "normal3.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2021, time.July, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id13,
			PatientID:         "CN=patient4,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "pneumonia3.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2021, time.June, 10, 12, 12, 15, 34, time.UTC),
		},
		{
			ID:                id14,
			PatientID:         "CN=patient5,OU=client,O=Hyperledger,ST=North Carolina,C=US",
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: "normal4.jpeg",
			MedicalEntryType:  "s3img",
			DateAdded:         time.Date(2021, time.May, 10, 12, 12, 15, 34, time.UTC),
		},
	}

	for _, entry := range entries {
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(entry.ID, entryJSON)
		if err != nil {
			return fmt.Errorf("MedicalDataSmartContract:InitLedger: failed to put to world state. %v", err)
		}
	}

	return nil
}

// AddEntry issues a new entry to the world state with given details.
func (s *MedicalDataSmartContract) AddMedicalEntry(ctx contractapi.TransactionContextInterface, patientID string, medicalEntryName string, medicalEntryType string, medicalEntryValue string, nonce string) error {
	patientContract := new(patientchaincode.PatientSmartContract)

	canWrite, err := patientContract.CanWrite(ctx, patientID)

	if err != nil {
		return err
	}

	canCompute, err := patientContract.CanCompute(ctx, patientID)

	if err != nil {
		return err
	}

	if !canWrite && !canCompute {
		fmt.Errorf("MedicalDataSmartContract:AddMedicalEntry: user does not have write/compute permissions")
	} else if !canWrite && canCompute {
		isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
		if err != nil {
			return err
		}
		if !isNonceValid {
			fmt.Errorf("MedicalDataSmartContract:AddMedicalEntry: nonce %s is not valid", nonce)
		}
	}

	id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, patientID, medicalEntryName, ctx.GetStub().GetTxID()})

	timestampRequested, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return err
	}
	timeRequested := time.Unix(int64(timestampRequested.GetSeconds()), int64(timestampRequested.GetNanos())).UTC()

	medicalEntry := medicaldatastructs.MedicalEntry{
		ID:                id,
		PatientID:         patientID,
		MedicalEntryType:  medicalEntryType,
		MedicalEntryName:  medicalEntryName,
		MedicalEntryValue: medicalEntryValue,
		DateAdded:         timeRequested,
	}
	medicalEntryJSON, err := json.Marshal(medicalEntry)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, medicalEntryJSON)
}

// ReadMedicalEntry returns the medical entry stored in the world state with given id.
func (s *MedicalDataSmartContract) ReadMedicalEntry(ctx contractapi.TransactionContextInterface, id string, nonce string) (*medicaldatastructs.MedicalEntry, error) {
	medicalEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("MedicalDataSmartContract:ReadMedicalEntry: failed to read from world state: %v", err)
	}
	if medicalEntryJSON == nil {
		return nil, fmt.Errorf("MedicalDataSmartContract:ReadMedicalEntry: the medical entry %s does not exist", id)
	}

	var medicalEntry medicaldatastructs.MedicalEntry
	err = json.Unmarshal(medicalEntryJSON, &medicalEntry)
	if err != nil {
		return nil, err
	}

	patientContract := new(patientchaincode.PatientSmartContract)

	canRead, err := patientContract.CanRead(ctx, medicalEntry.PatientID)

	if err != nil {
		return nil, err
	}

	canCompute, err := patientContract.CanRead(ctx, medicalEntry.PatientID)

	if err != nil {
		return nil, err
	}

	if !canRead && !canCompute {
		return nil, fmt.Errorf("MedicalDataSmartContract:ReadMedicalEntry: user does not have read/compute permissions to %s", id)
	} else if !canRead && canCompute {
		isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
		if err != nil {
			return nil, err
		}
		if !isNonceValid {
			return nil, fmt.Errorf("MedicalDataSmartContract:ReadMedicalEntry: nonce %s is not valid", nonce)
		}
	}

	return &medicalEntry, nil
}

// UpdateMedicalEntry updates an existing medical entry in the world state with provided parameters.
func (s *MedicalDataSmartContract) UpdateMedicalEntry(ctx contractapi.TransactionContextInterface, id string, patientID string, medicalEntryName string, medicalEntryType string, medicalEntryValue string) error {
	exists, err := s.MedicalEntryExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("MedicalDataSmartContract:UpdateMedicalEntry: the medicalEntry %s does not exist", id)
	}

	medicalEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("MedicalDataSmartContract:UpdateMedicalEntry: failed to read from world state: %v", err)
	}
	if medicalEntryJSON == nil {
		return fmt.Errorf("MedicalDataSmartContract:UpdateMedicalEntry: the medical entry %s does not exist", id)
	}

	var medicalEntry medicaldatastructs.MedicalEntry
	err = json.Unmarshal(medicalEntryJSON, &medicalEntry)
	if err != nil {
		return err
	}

	patientContract := new(patientchaincode.PatientSmartContract)

	canWrite, err := patientContract.CanWrite(ctx, medicalEntry.PatientID)

	if err != nil {
		return err
	}

	if !canWrite {
		return fmt.Errorf("MedicalDataSmartContract:UpdateMedicalEntry: user does not have write permissions to %s", id)
	}

	if id != medicalEntry.ID || medicalEntryName != medicalEntry.MedicalEntryName {
		return fmt.Errorf("MedicalDataSmartContract:UpdateMedicalEntry: can't change id and medical entry name")
	}

	// overwriting original medicalEntry with new
	medicalEntry = medicaldatastructs.MedicalEntry{
		ID:                id,
		PatientID:         patientID,
		MedicalEntryName:  medicalEntryName,
		MedicalEntryType:  medicalEntryType,
		MedicalEntryValue: medicalEntryValue,
		DateAdded:         medicalEntry.DateAdded,
	}
	medicalEntryJSON, err = json.Marshal(medicalEntry)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, medicalEntryJSON)
}

// DeleteMedicalEntry deletes an given medicalEntry from the world state.
func (s *MedicalDataSmartContract) DeleteMedicalEntry(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.MedicalEntryExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("MedicalDataSmartContract:DeleteMedicalEntry: the medicalEntry %s does not exist", id)
	}

	medicalEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("MedicalDataSmartContract:DeleteMedicalEntry: failed to read from world state: %v", err)
	}
	if medicalEntryJSON == nil {
		return fmt.Errorf("MedicalDataSmartContract:DeleteMedicalEntry: the medical entry %s does not exist", id)
	}

	var medicalEntry medicaldatastructs.MedicalEntry
	err = json.Unmarshal(medicalEntryJSON, &medicalEntry)
	if err != nil {
		return err
	}

	patientContract := new(patientchaincode.PatientSmartContract)

	canWrite, err := patientContract.CanWrite(ctx, medicalEntry.PatientID)

	if err != nil {
		return err
	}

	if !canWrite {
		return fmt.Errorf("MedicalDataSmartContract:DeleteMedicalEntry: user does not have write permissions to %s", id)
	}
	return ctx.GetStub().DelState(id)
}

// MedicalEntryExists returns true when medicalEntry with given ID exists in world state
func (s *MedicalDataSmartContract) MedicalEntryExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	medicalEntryJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("MedicalDataSmartContract:MedicalEntryExists: failed to read from world state: %v", err)
	}

	return medicalEntryJSON != nil, nil
}

// GetPatientMedicalEntries returns all medical entries found in world state for given patient
func (s *MedicalDataSmartContract) GetPatientMedicalEntries(ctx contractapi.TransactionContextInterface, patientID string, medicalEntryName string, dateStartTimestamp string, dateEndTimestamp string, nonce string) ([]*medicaldatastructs.MedicalEntry, error) {

	compositeKey := []string{KEY_NAME, patientID}
	if medicalEntryName != "" {
		compositeKey = append(compositeKey, medicalEntryName)
	}
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(INDEX_NAME, compositeKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	i, err := strconv.ParseInt(dateStartTimestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	dateStart := time.Unix(i, 0)

	i, err = strconv.ParseInt(dateEndTimestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	dateEnd := time.Unix(i, 0)

	fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: Parsed timestamps. DateStart: %s  DateEnd: %s\n", dateStart.String(), dateEnd.String())

	var medicalEntries []*medicaldatastructs.MedicalEntry
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: Deserialize JSON: %s\n", queryResponse.Value)

		var medicalEntry medicaldatastructs.MedicalEntry
		err = json.Unmarshal(queryResponse.Value, &medicalEntry)
		if err != nil {
			return nil, err
		}

		fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: Patient ID: %s\n", medicalEntry.PatientID)

		if medicalEntry.PatientID != patientID {
			continue
		}

		patientContract := new(patientchaincode.PatientSmartContract)

		canRead, err := patientContract.CanRead(ctx, medicalEntry.PatientID)

		fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: Can read: %t\n", canRead)

		if err != nil {
			return nil, err
		}

		canCompute, err := patientContract.CanCompute(ctx, medicalEntry.PatientID)

		fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: Can compute: %t\n", canCompute)

		if err != nil {
			return nil, err
		}

		if !canRead && !canCompute {
			continue
		} else if !canRead && canCompute {
			isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
			if err != nil {
				return nil, err
			}
			if !isNonceValid {
				continue
			}
		}

		fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: MedicalEntryName: %s\n", medicalEntry.MedicalEntryName)

		// check if medical entry name is correct or skip if we are looking for all entries
		if medicalEntry.MedicalEntryName != medicalEntryName && medicalEntryName != "" {
			continue
		}

		// Check if data range is correct
		if dateStart.After(medicalEntry.DateAdded) || dateEnd.Before(medicalEntry.DateAdded) {
			continue
		}

		fmt.Printf("MedicalDataSmartContract:GetPatientMedicalEntries: Appending %s to list\n", medicalEntry.ID)

		medicalEntries = append(medicalEntries, &medicalEntry)
	}

	return medicalEntries, nil
}

// GetAllEntriesAdmin returns all medical entries found in world state. Only admin can execute this
func (s *MedicalDataSmartContract) GetAllEntriesAdmin(ctx contractapi.TransactionContextInterface) ([]*medicaldatastructs.MedicalEntry, error) {
	x509, _ := cid.GetX509Certificate(ctx.GetStub())
	err := cid.AssertAttributeValue(ctx.GetStub(), "hf.Type", "admin")

	if err != nil {
		return nil, fmt.Errorf("MedicalDataSmartContract:GetAllEntriesAdmin: only admin can do this. Current user: %s", x509.Subject.ToRDNSequence().String())
	}

	compositeKey := []string{KEY_NAME}
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(INDEX_NAME, compositeKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var medicalEntries []*medicaldatastructs.MedicalEntry
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var medicalEntry medicaldatastructs.MedicalEntry
		err = json.Unmarshal(queryResponse.Value, &medicalEntry)
		if err != nil {
			return nil, err
		}

		medicalEntries = append(medicalEntries, &medicalEntry)
	}

	return medicalEntries, nil
}

// GetPatientMedicalEntries returns all medical entries found in world state
func (s *MedicalDataSmartContract) GetMedicalEntries(ctx contractapi.TransactionContextInterface, medicalEntryName string, dateStartTimestamp string, dateEndTimestamp string, nonce string) ([]*medicaldatastructs.MedicalEntry, error) {

	compositeKey := []string{KEY_NAME}

	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(INDEX_NAME, compositeKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	i, err := strconv.ParseInt(dateStartTimestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	dateStart := time.Unix(i, 0)

	i, err = strconv.ParseInt(dateEndTimestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	dateEnd := time.Unix(i, 0)

	fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: Parsed timestamps. DateStart: %s  DateEnd: %s\n", dateStart.String(), dateEnd.String())

	var medicalEntries []*medicaldatastructs.MedicalEntry
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: Deserialize JSON: %s\n", queryResponse.Value)

		var medicalEntry medicaldatastructs.MedicalEntry
		err = json.Unmarshal(queryResponse.Value, &medicalEntry)
		if err != nil {
			return nil, err
		}

		fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: Patient ID: %s\n", medicalEntry.PatientID)

		patientContract := new(patientchaincode.PatientSmartContract)

		canRead, err := patientContract.CanRead(ctx, medicalEntry.PatientID)

		fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: Can read: %t\n", canRead)

		if err != nil {
			return nil, err
		}

		canCompute, err := patientContract.CanCompute(ctx, medicalEntry.PatientID)

		fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: Can compute: %t\n", canCompute)

		if err != nil {
			return nil, err
		}

		if !canRead && !canCompute {
			continue
		} else if !canRead && canCompute {
			isNonceValid, err := tokenapi.IsNonceValid(ctx, nonce)
			if err != nil {
				return nil, err
			}
			if !isNonceValid {
				continue
			}
		}

		fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: MedicalEntryName: %s\n", medicalEntry.MedicalEntryName)

		// check if medical entry name is correct or skip if we are looking for all entries
		if medicalEntry.MedicalEntryName != medicalEntryName && medicalEntryName != "" {
			continue
		}

		// Check if data range is correct
		if dateStart.After(medicalEntry.DateAdded) || dateEnd.Before(medicalEntry.DateAdded) {
			continue
		}

		fmt.Printf("MedicalDataSmartContract:GetMedicalEntries: Appending %s to list\n", medicalEntry.ID)

		medicalEntries = append(medicalEntries, &medicalEntry)
	}

	return medicalEntries, nil
}
