package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_ACCOUNT                    api.APIOperation = "addAccount"
	UPDATE_ACCOUNT                 api.APIOperation = "updateAccount"
	DELETE_ACCOUNT                 api.APIOperation = "deleteAccount"
	CLONE_ACCOUNT                  api.APIOperation = "cloneAccount"
	CHANGE_ACCOUNT_OWNER           api.APIOperation = "changeOwner"
	CHANGE_ACCOUNT_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	// READ Operations.
	LIST_ACCOUNTS     api.APIOperation = "listAccounts"
	COUNT_ACCOUNTS    api.APIOperation = "countAccounts"
	GET_ACCOUNT_BY_ID api.APIOperation = "getAccountById"
)
