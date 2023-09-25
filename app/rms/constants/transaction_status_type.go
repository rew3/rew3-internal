package consts

type TransactionStatusType string

const (
	TS_ACTIVE                 TransactionStatusType = "ACTIVE"
	TS_PENDING_UNDER_CONTRACT TransactionStatusType = "PENDING_UNDER_CONTRACT"
	TS_CLOSED                 TransactionStatusType = "CLOSED"
	TS_CANCELLED              TransactionStatusType = "CANCELLED"
)
