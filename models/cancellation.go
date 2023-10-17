package models

type CancellationRequest struct {
	TransactionId string `json:"transactionId"`
	Amount        string `json:"amount"`
	ApprovalCode  string `json:"approvalCode"`
	CardNumber    string `json:"cardNumber"`
}

type CancellationResponse struct {
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
