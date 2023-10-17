package models

type ResetPinPadRequest struct {
	TransactionId string `json:"transactionId"`
}

type ResetPinPadResponse struct {
	TransactionId   string `json:"transactionId"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
