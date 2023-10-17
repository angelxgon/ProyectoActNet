package models

type SaleRequest struct {
	TransactionId string `json:"transactionId"`
	Amount        string `json:"amount"`
	Mounths       string `json:"mounths"`
}

type SaleResponse struct {
	Amount                       string `json:"amount"`
	CardNumberMasked             string `json:"cardNumberMasked"`
	EntryMode                    string `json:"entryMode"`
	ResponseCode                 string `json:"responseCode"`
	ResponseMessage              string `json:"responseMessage"`
	TerminalSerialNumber         string `json:"terminalSerialNumber"`
	ApprovalCode                 string `json:"approvalCode"`
	CardBrand                    string `json:"cardBrand"`
	CardIssuer                   string `json:"cardIssuer"`
	AccountingNature             string `json:"accountingNature"`
	Affilitation                 string `json:"affilitation"`
	AID                          string `json:"AID"`
	ARQC                         string `json:"ARQC"`
	CardHolderVerificationMethod string `json:"cardHolderVerificationMethod"`
}
