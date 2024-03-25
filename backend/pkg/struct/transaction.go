package structure

type CreateQRPromptPayRequest struct {
	UserId            string  `json:"user_id" binding:"required"`
	TransactionAmount float64 `json:"transaction_amount" binding:"required"`
}

type CreateQRPromptPayResponse struct {
	PaymentIntentId     string  `json:"payment_intent_id"`
	PaymentClientSecret string  `json:"payment_client_secret"`
	PaymentMethodType   string  `json:"payment_method_type"`
	TransactionAmount   float64 `json:"transaction_amount"`
}

type CreateTransactionRequest struct {
	UserId            string  `json:"user_id" binding:"required"`
	TransactionAmount float64 `json:"transaction_amount" binding:"required"`
	Status            string  `json:"status" binding:"required"`
}

type CreateTransactionResponse struct {
	TransactionId string `json:"transaction_id"`
}

type UpdateTransactionRequest struct {
	TransactionId string `json:"transaction_id" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

type GetTransactionByPaymentIdRequest struct {
	PaymentIntentId string `json:"payment_intent_id" binding:"required"`
}

type GetTransactionByPaymentIdResponse struct {
	TransactionId     string  `json:"transaction_id"`
	UserId            string  `json:"user_id"`
	PaymentIntentId   string  `json:"payment_intent_id"`
	Status            string  `json:"status"`
	TransactionAmount float64 `json:"transaction_amount"`
}

// ? Use for ...
/*
	UpdateTransactionResponse, CreateQRPromptPayResponse
*/
type TransactionResponse struct {
	Response string
}
