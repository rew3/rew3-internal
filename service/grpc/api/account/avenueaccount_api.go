package account

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for AvenueAccount
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AVENUEACCOUNT    api.APIOperation = "account_addAvenueAccount"
	UPDATE_AVENUEACCOUNT api.APIOperation = "account_updateAvenueAccount"
	DELETE_AVENUEACCOUNT api.APIOperation = "account_deleteAvenueAccount"

	// READ APIs.
	LIST_AVENUEACCOUNT      api.APIOperation = "account_listAvenueAccount"
	COUNT_AVENUEACCOUNT     api.APIOperation = "account_countAvenueAccount"
	GET_AVENUEACCOUNT_BY_ID api.APIOperation = "caccount_getAvenueAccountById"
)
