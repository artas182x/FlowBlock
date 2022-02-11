package controllers

import (
	"encoding/json"
	"strings"

	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/models"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/services"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/vars"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-computationtoken/tokenapi"
	"github.com/gin-gonic/gin"
)

type AvailableMethodsStruct struct {
	ChaincodeName string `form:"chaincodeName" json:"chaincodeName" binding:"required"`
}

type RequestTokenInput struct {
	ChaincodeName string   `form:"chaincodeName" json:"chaincodeName" binding:"required"`
	Method        string   `form:"method" json:"method" binding:"required"`
	Arguments     []string `form:"arguments" json:"arguments" binding:"required"`
}

// @Summary GetAvailableMethods
// @Schemes
// @Produce json
// @Success 200
// @Tags Computation
// @Param chaincode_name path string true "Chaincode name"
// @Router /v1/computation/availablemethods/{chaincode_name} [get]
// @Security Bearer
func GetAvailableMethods(c *gin.Context) {
	chaincodeName := c.Param("chaincode_name")

	user, _ := c.Get(vars.IdentityKey)

	out, err := services.EvaluateTransaction(user.(*models.User).Login, vars.ComputationTokenChaincodeName, vars.ComputationTokenSmartContractName, "GetAvailableMethods", chaincodeName)

	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}

	var methodsResponse []tokenapi.Method
	err = json.Unmarshal(out, &methodsResponse)
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}

	var methods []models.Method

	for _, method := range methodsResponse {

		var arguments []models.Argument

		for _, argString := range strings.Split(method.Args, ";") {
			argStrSplit := strings.Split(argString, ":")
			arguments = append(arguments, models.Argument{Name: argStrSplit[0], Type: argStrSplit[1]})
		}

		var method = models.Method{Name: method.Name, Description: method.Description, RetType: method.RetType, Arguments: arguments}
		methods = append(methods, method)
	}

	c.JSON(200, methods)

}

// @Summary RequestToken
// @Schemes
// @Produce json
// @Success 200
// @Tags Computation
// @Param input body RequestTokenInput true "Request tokeninput data"
// @Router /v1/computation/requesttoken [post]
// @Security Bearer
func RequestToken(c *gin.Context) {
	var input RequestTokenInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user, _ := c.Get(vars.IdentityKey)

	var argsString string

	for index, arg := range input.Arguments {
		argsString += arg
		if index < len(input.Arguments)-1 {
			argsString += ";"
		}
	}

	out, err := services.SubmitTransaction(user.(*models.User).Login, vars.ComputationTokenChaincodeName, vars.ComputationTokenSmartContractName, "RequestToken", input.ChaincodeName, input.Method, argsString)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var token tokenapi.Token
	err = json.Unmarshal(out, &token)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, token)
	}

}

// @Summary ReadUserTokens
// @Schemes
// @Produce json
// @Success 200
// @Tags Computation
// @Router /v1/computation/usertokens [get]
// @Security Bearer
func ReadUserTokens(c *gin.Context) {
	user, _ := c.Get(vars.IdentityKey)

	out, err := services.EvaluateTransaction(user.(*models.User).Login, vars.ComputationTokenChaincodeName, vars.ComputationTokenSmartContractName, "ReadUserTokens")

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var token []tokenapi.Token

	if len(out) == 0 {
		c.JSON(204, "")
		return
	}

	err = json.Unmarshal(out, &token)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, token)
	}

}

// @Summary ReadToken
// @Schemes
// @Produce json
// @Success 200
// @Tags Computation
// @Param token_id path string true "Token id"
// @Router /v1/computation/token/{token_id} [get]
// @Security Bearer
func ReadToken(c *gin.Context) {
	tokenId := c.Param("token_id")

	user, _ := c.Get(vars.IdentityKey)

	out, err := services.EvaluateTransaction(user.(*models.User).Login, vars.ComputationTokenChaincodeName, vars.ComputationTokenSmartContractName, "ReadToken", tokenId)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var token tokenapi.Token
	err = json.Unmarshal(out, &token)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, token)
	}
}

// @Summary StartComputation
// @Schemes
// @Produce json
// @Success 200
// @Tags Computation
// @Param token_id path string true "Token id"
// @Router /v1/computation/token/{token_id}/start [post]
// @Security Bearer
func StartComputation(c *gin.Context) {
	tokenId := c.Param("token_id")

	user, _ := c.Get(vars.IdentityKey)

	task, err := services.QueueComputation(user.(*models.User).Login, tokenId)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, task)
	}

}

// @Summary GetQueue
// @Schemes
// @Produce json
// @Success 200
// @Tags Computation
// @Router /v1/computation/queue [get]
// @Security Bearer
func GetQueue(c *gin.Context) {

	user, _ := c.Get(vars.IdentityKey)

	out := services.GetUsersRunningComputations(user.(*models.User).Login)

	if len(out) == 0 {
		c.JSON(204, gin.H{})
	} else {
		c.JSON(200, out)
	}

}
