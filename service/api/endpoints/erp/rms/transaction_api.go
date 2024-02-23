package rms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_TRANSACTION          endpoints.Endpoint = "rms_addTransaction"
	UPDATE_TRANSACTION       endpoints.Endpoint = "rms_updateTransaction"
	DELETE_TRANSACTION       endpoints.Endpoint = "rms_deleteTransaction"
	CLONE_TRANSACTION        endpoints.Endpoint = "rms_cloneTransaction"
	CHANGE_TRANSACTION_OWNER endpoints.Endpoint = "rms_changeTransactionOwner"
	SET_TRANSACTION_FAVORITE endpoints.Endpoint = "crm_setTransactionFavorite"

	// READ Operations.
	LIST_TRANSACTIONS     endpoints.Endpoint = "rms_listTransactions"
	COUNT_TRANSACTIONS    endpoints.Endpoint = "rms_countTransactions"
	GET_TRANSACTION_BY_ID endpoints.Endpoint = "rms_getTransactionById"
)
