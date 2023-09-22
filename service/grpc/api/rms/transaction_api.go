package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_TRANSACTION                    api.APIOperation = "addTransaction"
	UPDATE_TRANSACTION                 api.APIOperation = "updateTransaction"
	DELETE_TRANSACTION                 api.APIOperation = "deleteTransaction"
	CLONE_TRANSACTION                  api.APIOperation = "cloneTransaction"
	CHANGE_TRANSACTION_OWNER           api.APIOperation = "changeTransactionOwner"

	// READ Operations.
	LIST_TRANSACTIONS     api.APIOperation = "listTransactions"
	COUNT_TRANSACTIONS    api.APIOperation = "countTransactions"
	GET_TRANSACTION_BY_ID api.APIOperation = "getTransactionById"
)
