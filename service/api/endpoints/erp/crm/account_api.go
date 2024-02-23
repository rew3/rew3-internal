package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_ACCOUNT          endpoints.Endpoint = "crm_addAccount"
	UPDATE_ACCOUNT       endpoints.Endpoint = "crm_updateAccount"
	DELETE_ACCOUNT       endpoints.Endpoint = "crm_deleteAccount"
	CLONE_ACCOUNT        endpoints.Endpoint = "crm_cloneAccount"
	CHANGE_ACCOUNT_OWNER endpoints.Endpoint = "crm_changeAccountOwner"
	SET_ACCOUNT_FAVORITE endpoints.Endpoint = "crm_setAccountFavorite"

	// READ Operations.
	LIST_ACCOUNTS     endpoints.Endpoint = "crm_listAccounts"
	COUNT_ACCOUNTS    endpoints.Endpoint = "crm_countAccounts"
	GET_ACCOUNT_BY_ID endpoints.Endpoint = "crm_getAccountById"
)
