package consts

type TransactionStageType string

const (
	TS_OPEN           TransactionStageType = "OPEN"
	TS_OPEN_PRE_CLOSE TransactionStageType = "PRE_CLOSE"
	TS_OPEN_PAYMENT   TransactionStageType = "PAYMENT"
	TS_OPEN_CLOSED    TransactionStageType = "CLOSED"
)
