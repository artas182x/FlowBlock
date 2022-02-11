package models

type Method struct {
	Name        string
	Description string
	RetType     string
	Arguments   []Argument
}

type Argument struct {
	Name string
	Type string
}
