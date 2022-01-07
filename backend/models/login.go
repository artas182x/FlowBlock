package models

type Login struct {
	Certificate string `form:"certificate" json:"certificate" binding:"required"`
	PrivateKey  string `form:"privateKey" json:"privateKey" binding:"required"`
	MspID       string `form:"mspid" json:"mspid" binding:"required"`
}
