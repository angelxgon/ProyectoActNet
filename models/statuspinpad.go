package models

type StatusPinPadRequest struct {
	TransactionId string `json:"transactionId"`
}

type StatusPinPadResponse struct {
	TransactionId   string `json:"transactionId"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
