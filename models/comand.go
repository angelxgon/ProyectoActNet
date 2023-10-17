package models

type Commands struct {
	TypeTrans    string      `json:"typeTrans"`
	Sale         SaleRequest `json:"sale"`
	Refund       SaleRequest `json:"refund"`
	Cancellation SaleRequest `json:"cancellation"`
	EchoPinPad   SaleRequest `json:"echoPinPad"`
	StatusPinPad SaleRequest `json:"statusPinPad"`
	ConfigPinPad SaleRequest `json:"configPinPad"`
	InitPinPad   SaleRequest `json:"initPinPad"`
	ResetPinPad  SaleRequest `json:"resetPinPad"`
}
