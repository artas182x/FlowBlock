package tokenapi

import (
	"time"
)

// In debug mode for example file checksums are skipped, because example data contains invalid ones
const DEBUG_MODE = true

type Method struct {
	Name        string     `json:"name"`
	Args        []Argument `json:arguments`
	RetType     string     `json:"retType"`
	Description string     `json:"description"`
}

type Ret struct {
	RetValue string `json:retValue`
	RetType  string `json:retType`
}

type Argument struct {
	Value string `json:value`
	Type  string `json:type`
	Name  string `json:name`
}

type Token struct {
	ID                 string     `json:"ID"` //Must be string
	UserRequested      string     `json:userRequested`
	ChaincodeName      string     `json:chaincodeName`
	Method             string     `json:method`
	Arguments          []Argument `json:arguments`
	Ret                Ret        `json:"ret,omitempty" metadata:"ret,optional"`
	TimeRequested      time.Time  `json:timeRequested`
	ExpirationTime     time.Time  `json:expirationTime`
	Description        string     `json:description`
	DirectlyExecutable bool       `json:directlyExecutable`
}

type Interface struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Options struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type Connection struct {
	ID   string `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Node struct {
	Type               string      `json:"type"`
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	ChaincodeName      string      `json:"chaincodeName"`
	TokenId            string      `json:"tokenId"`
	MethodName         string      `json:"methodName"`
	Options            []Options   `json:"options"`
	Interfaces         []Interface `json:"interfaces"`
	DirectlyExecutable bool        `json:"directlyExecutable"`
}

type Flow struct {
	Nodes       []Node       `json:"nodes"`
	Connections []Connection `json:"connections"`
}
