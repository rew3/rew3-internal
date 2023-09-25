package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_TRANSACTION          api.APIOperation = "rms_addTransaction"
	UPDATE_TRANSACTION       api.APIOperation = "rms_updateTransaction"
	DELETE_TRANSACTION       api.APIOperation = "rms_deleteTransaction"
	CLONE_TRANSACTION        api.APIOperation = "rms_cloneTransaction"
	CHANGE_TRANSACTION_OWNER api.APIOperation = "rms_changeTransactionOwner"

	// READ Operations.
	LIST_TRANSACTIONS     api.APIOperation = "rms_listTransactions"
	COUNT_TRANSACTIONS    api.APIOperation = "rms_countTransactions"
	GET_TRANSACTION_BY_ID api.APIOperation = "rms_getTransactionById"
)
