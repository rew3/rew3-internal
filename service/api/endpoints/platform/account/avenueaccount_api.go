package account

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for AvenueAccount
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AVENUEACCOUNT    endpoints.Endpoint = "account_addAvenueAccount"
	UPDATE_AVENUEACCOUNT endpoints.Endpoint = "account_updateAvenueAccount"
	DELETE_AVENUEACCOUNT endpoints.Endpoint = "account_deleteAvenueAccount"

	// READ APIs.
	LIST_AVENUEACCOUNT      endpoints.Endpoint = "account_listAvenueAccount"
	COUNT_AVENUEACCOUNT     endpoints.Endpoint = "account_countAvenueAccount"
	GET_AVENUEACCOUNT_BY_ID endpoints.Endpoint = "caccount_getAvenueAccountById"
)
