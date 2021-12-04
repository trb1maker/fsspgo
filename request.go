package fsspgo

type innerRequest struct {
	Type   requestResponseType `json:"type"`
	Params GroupParam          `json:"params"`
}

type groupRequest struct {
	Token   string         `json:"token"`
	Request []innerRequest `json:"request"`
}

type requestResponseType uint8

const (
	personRequestResponseType requestResponseType = iota + 1
	legalRequestRequestResponseType
	numberRequestRequestResponseType
)
