package account

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for User
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_USER               endpoints.Endpoint = "account_addUser"
	UPDATE_USER            endpoints.Endpoint = "account_updateUser"
	DELETE_USER            endpoints.Endpoint = "account_deleteUser"
	CLONE_USER             endpoints.Endpoint = "account_cloneUser"
	CHANGE_USER_OWNER      endpoints.Endpoint = "account_changeUserOwner"
	BULK_ADD_USER          endpoints.Endpoint = "account_bulkAddUser"
	BULK_UPDATE_USER       endpoints.Endpoint = "account_bulkUpdateUser"
	BULK_DELETE_USER       endpoints.Endpoint = "account_bulkDeleteUser"
	BULK_CHANGE_USER_OWNER endpoints.Endpoint = "account_bulkChangeUserOwner"

	// READ APIs.
	LIST_USERS     endpoints.Endpoint = "account_listUsers"
	COUNT_USERS    endpoints.Endpoint = "account_countUsers"
	GET_USER_BY_ID endpoints.Endpoint = "caccount_getUserById"
)
