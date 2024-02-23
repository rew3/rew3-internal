package group

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Group
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_GROUP               endpoints.Endpoint = "collaboration_addGroup"
	UPDATE_GROUP            endpoints.Endpoint = "collaboration_updateGroup"
	DELETE_GROUP            endpoints.Endpoint = "collaboration_deleteGroup"
	CLONE_GROUP             endpoints.Endpoint = "collaboration_cloneGroup"
	CHANGE_GROUP_OWNER      endpoints.Endpoint = "collaboration_changeGroupOwner"
	BULK_ADD_GROUP          endpoints.Endpoint = "collaboration_bulkAddGroup"
	BULK_UPDATE_GROUP       endpoints.Endpoint = "collaboration_bulkUpdateGroup"
	BULK_DELETE_GROUP       endpoints.Endpoint = "collaboration_bulkDeleteGroup"
	BULK_CHANGE_GROUP_OWNER endpoints.Endpoint = "collaboration_bulkChangeGroupOwner"

	// READ APIs.
	LIST_GROUPS     endpoints.Endpoint = "collaboration_listGroups"
	COUNT_GROUPS    endpoints.Endpoint = "collaboration_countGroups"
	GET_GROUP_BY_ID endpoints.Endpoint = "ccollaboration_getGroupById"
)
