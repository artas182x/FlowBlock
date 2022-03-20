package controllers

import (
	"encoding/json"

	"encoding/base64"

	"github.com/artas182x/FlowBlock/backend/models"
	"github.com/artas182x/FlowBlock/backend/services"
	"github.com/artas182x/FlowBlock/backend/vars"
	"github.com/artas182x/FlowBlock/chaincode-medicaldata/medicaldatastructs"
	"github.com/gin-gonic/gin"
)

type RequestMedicalDataInput struct {
	MedicalEntryName   string
	DateStartTimestamp string
	DateEndTimestamp   string
}

// @Summary GetMedicalData
// @Schemes
// @Produce json
// @Success 200
// @Tags MedicalData
// @Param input body RequestMedicalDataInput true "Request medical data"
// @Router /v1/medicaldata/request [post]
// @Security Bearer
func GetMedicalData(c *gin.Context) {

	var input RequestMedicalDataInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	user, _ := c.Get(vars.IdentityKey)

	out, err := services.EvaluateTransaction(user.(*models.User).Login, vars.MedicaldataChaincodeName, vars.MedicaldataSmartContractName, "GetMedicalEntries", input.MedicalEntryName, input.DateStartTimestamp, input.DateEndTimestamp, "")

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if len(out) == 0 {
		c.JSON(204, "")
		return
	}

	var medicaldata []medicaldatastructs.MedicalEntry
	err = json.Unmarshal(out, &medicaldata)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		for i, entry := range medicaldata {

			idEncoded := base64.StdEncoding.EncodeToString([]byte(entry.ID))

			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
			}

			medicaldata[i].ID = idEncoded

		}

		c.JSON(200, medicaldata)
	}
}
