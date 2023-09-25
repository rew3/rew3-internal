package consts

type TransactionAgentType string

const (
	TA_SELLING  TransactionAgentType = "SELLING"
	TA_LISTING  TransactionAgentType = "LISTING"
	TA_BOTH     TransactionAgentType = "BOTH"
	TA_DUAL     TransactionAgentType = "DUAL"
	TA_RENTAL   TransactionAgentType = "RENTAL"
	TA_REFERRAL TransactionAgentType = "REFERRAL"
)
