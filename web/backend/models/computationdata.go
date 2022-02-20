package models

import "github.com/gocelery/gocelery"

type ComputationData struct {
	TokenId         string
	UserCertificate string
	TaskResult      *gocelery.AsyncResult
}
