package models

type ConfigPinPadRequest struct {
	TransactionId string `json:"transactionId"`
}

type ConfigPinPadResponse struct {
	TransactionId   string `json:"transactionId"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
