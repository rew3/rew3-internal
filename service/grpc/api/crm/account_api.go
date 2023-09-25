package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_ACCOUNT                    api.APIOperation = "crm_addAccount"
	UPDATE_ACCOUNT                 api.APIOperation = "crm_updateAccount"
	DELETE_ACCOUNT                 api.APIOperation = "crm_deleteAccount"
	CLONE_ACCOUNT                  api.APIOperation = "crm_cloneAccount"
	CHANGE_ACCOUNT_OWNER           api.APIOperation = "crm_changeAccountOwner"
	CHANGE_ACCOUNT_FAVORITE_STATUS api.APIOperation = "crm_changeAccountFavoriteStatus"

	// READ Operations.
	LIST_ACCOUNTS     api.APIOperation = "crm_listAccounts"
	COUNT_ACCOUNTS    api.APIOperation = "crm_countAccounts"
	GET_ACCOUNT_BY_ID api.APIOperation = "crm_getAccountById"
)
