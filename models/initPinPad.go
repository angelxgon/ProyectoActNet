package models

type InitPinPadRequest struct {
	TransactionId string `json:"transactionId"`
}

type InitPinPadResponse struct {
	TransactionId   string `json:"transactionId"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
