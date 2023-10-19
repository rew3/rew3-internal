package account

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for User
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_USER               api.APIOperation = "account_addUser"
	UPDATE_USER            api.APIOperation = "account_updateUser"
	DELETE_USER            api.APIOperation = "account_deleteUser"
	CLONE_USER             api.APIOperation = "account_cloneUser"
	CHANGE_USER_OWNER      api.APIOperation = "account_changeUserOwner"
	BULK_ADD_USER          api.APIOperation = "account_bulkAddUser"
	BULK_UPDATE_USER       api.APIOperation = "account_bulkUpdateUser"
	BULK_DELETE_USER       api.APIOperation = "account_bulkDeleteUser"
	BULK_CHANGE_USER_OWNER api.APIOperation = "account_bulkChangeUserOwner"

	// READ APIs.
	LIST_USERS     api.APIOperation = "account_listUsers"
	COUNT_USERS    api.APIOperation = "account_countUsers"
	GET_USER_BY_ID api.APIOperation = "caccount_getUserById"
)
