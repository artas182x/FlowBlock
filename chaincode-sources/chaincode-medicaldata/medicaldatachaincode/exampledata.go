package medicaldatachaincode

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-medicaldata/medicaldatastructs"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-medicaldata/patientchaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const PNEUMONIA_IMG_NUM = 20
const NO_PNEUMONIA_IMG_NUM = 50

const HEART_RATE_GEN = 20
const BLOOD_PREASSURE_GEN = 30

var global_id_counter = 1

// Example data to fill ledger for testing purposes. Should be removed in production.
func (s *MedicalDataSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	entries := []medicaldatastructs.MedicalEntry{}

	entries = append(entries, genRandomXray(ctx)...)
	entries = append(entries, genHeartRate(ctx)...)
	entries = append(entries, genBloodPreasure(ctx)...)

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

func randDate(currId int) time.Time {
	min := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	sec := min + int64(currId)*1000
	return time.Unix(sec, 0)
}

func genHeartRate(ctx contractapi.TransactionContextInterface) []medicaldatastructs.MedicalEntry {

	entries := []medicaldatastructs.MedicalEntry{}

	for i := 1; i < HEART_RATE_GEN; i++ {

		patientId := i%patientchaincode.PATIENTS_NUM + 1
		id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId), "ChestXRay", fmt.Sprintf("%s%d", ctx.GetStub().GetTxID(), global_id_counter)})
		global_id_counter += 1

		preassure := 80 + i%70

		entry := medicaldatastructs.MedicalEntry{
			ID:                id,
			PatientID:         fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId),
			MedicalEntryName:  "HeartRate",
			MedicalEntryValue: fmt.Sprintf("%d", preassure),
			MedicalEntryType:  "int64",
			DateAdded:         randDate(i),
		}

		entries = append(entries, entry)
	}

	return entries
}

func genBloodPreasure(ctx contractapi.TransactionContextInterface) []medicaldatastructs.MedicalEntry {

	entries := []medicaldatastructs.MedicalEntry{}

	for i := 1; i < BLOOD_PREASSURE_GEN; i++ {

		patientId := i%patientchaincode.PATIENTS_NUM + 1
		id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId), "ChestXRay", fmt.Sprintf("%s%d", ctx.GetStub().GetTxID(), global_id_counter)})
		global_id_counter += 1

		preassure := 110 + i%40

		entry := medicaldatastructs.MedicalEntry{
			ID:                id,
			PatientID:         fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId),
			MedicalEntryName:  "SystolicBloodPreasure",
			MedicalEntryValue: fmt.Sprintf("%d", preassure),
			MedicalEntryType:  "int64",
			DateAdded:         randDate(i),
		}

		entries = append(entries, entry)
	}

	return entries
}

func GetSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func genRandomXray(ctx contractapi.TransactionContextInterface) []medicaldatastructs.MedicalEntry {

	entries := []medicaldatastructs.MedicalEntry{}

	for i := 1; i < NO_PNEUMONIA_IMG_NUM; i++ {

		patientId := i%patientchaincode.PATIENTS_NUM + 1
		id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId), "ChestXRay", fmt.Sprintf("%s%d", ctx.GetStub().GetTxID(), global_id_counter)})
		global_id_counter += 1

		fileName := fmt.Sprintf("normal%d.jpeg", i)

		entry := medicaldatastructs.MedicalEntry{
			ID:                id,
			PatientID:         fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId),
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: fmt.Sprintf("%s?%s", fileName, GetSHA256Hash(fileName)),
			MedicalEntryType:  "s3img",
			DateAdded:         randDate(i),
		}

		entries = append(entries, entry)
	}

	for i := 1; i < PNEUMONIA_IMG_NUM; i++ {

		patientId := i%patientchaincode.PATIENTS_NUM + 1
		id, _ := ctx.GetStub().CreateCompositeKey(INDEX_NAME, []string{KEY_NAME, fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId), "ChestXRay", fmt.Sprintf("%s%d", ctx.GetStub().GetTxID(), global_id_counter)})
		global_id_counter += 1

		fileName := fmt.Sprintf("pneumonia%d.jpeg", i)

		entry := medicaldatastructs.MedicalEntry{
			ID:                id,
			PatientID:         fmt.Sprintf("CN=patient%d,OU=client,O=Hyperledger,ST=North Carolina,C=US", patientId),
			MedicalEntryName:  "ChestXRay",
			MedicalEntryValue: fmt.Sprintf("%s?%s", fileName, GetSHA256Hash(fileName)),
			MedicalEntryType:  "s3img",
			DateAdded:         randDate(i),
		}

		entries = append(entries, entry)
	}

	return entries
}
