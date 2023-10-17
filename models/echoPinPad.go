package models

type EchoPinPadRequest struct {
	TransactionId string `json:"transactionId"`
}

type EchoPinPadResponse struct {
	TransactionId   string `json:"transactionId"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
