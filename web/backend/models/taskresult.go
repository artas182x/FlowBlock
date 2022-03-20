package models

import "github.com/artas182x/FlowBlock/chaincode-sources/chaincode-computationtoken/tokenapi"

type TaskResult struct {
	Finished bool
	Result   tokenapi.Token
	Error    error
	TaskID   string
}
