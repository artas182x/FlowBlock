package medicaldatastructs

import "time"

type MedicalEntry struct {
	ID                string    `json:"ID"` //Must be string
	PatientID         string    `json:"PatientID"`
	MedicalEntryName  string    `json:"MedicalEntryName"`
	MedicalEntryType  string    `json:"MedicalEntryType"`
	MedicalEntryValue string    `json:"MedicalEntryValue"`
	DateAdded         time.Time `json:"DateAdded"`
}

type PatientEntry struct {
	ID                      string   `json:"ID"` //Must be string
	OrgReadAllowed          []string `json:"OrgReadAllowed,omitempty" metadata:"OrgReadAllowed,optional"`
	UsersReadAllowed        []string `json:"UsersReadAllowed,omitempty" metadata:"UsersReadAllowed,optional"`
	OrgComputationAllowed   []string `json:"OrgComputationAllowed,omitempty" metadata:"OrgComputationAllowed,optional"`
	UsersComputationAllowed []string `json:"UsersComputationAllowed,omitempty" metadata:"UsersComputationAllowed,optional"`
	OrgWriteAllowed         []string `json:"OrgWriteAllowed,omitempty" metadata:"OrgWriteAllowed,optional"`
	UsersWriteAllowed       []string `json:"UsersWriteAllowed,omitempty" metadata:"UsersWriteAllowed,optional"`
}
