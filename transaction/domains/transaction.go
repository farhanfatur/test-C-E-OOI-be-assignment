package domains

type StatusTransaction string

const (
	ProcessStatusTransaction   StatusTransaction = "process"
	OnPaymentStatusTransaction StatusTransaction = "on payment"
	PaidStatusTransaction      StatusTransaction = "paid"
	DoneStatusTransaction      StatusTransaction = "done"
)

type TransactionRequest struct {
	Amount      int               `json:"amount"`
	ToAddress   string            `json:"toAddress"`
	FromAddress string            `json:"fromAddress"`
	Currency    string            `json:"currency"`
	Status      StatusTransaction `json:"status"`
}

type WithdrawRequest struct {
	Charge    int `json:"charge"`
	DepositId int `json:"deposit_id"`
}
