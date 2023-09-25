package consts

type TransactionType string

const (
	T_SELLING  TransactionType = "SELLING"
	T_LISTING  TransactionType = "LISTING"
	T_BOTH     TransactionType = "BOTH"
	T_DUAL     TransactionType = "DUAL"
	T_RENTAL   TransactionType = "RENTAL"
	T_REFERRAL TransactionType = "REFERRAL"
)
