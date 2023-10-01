package group

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Group
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_GROUP               api.APIOperation = "collaboration_addGroup"
	UPDATE_GROUP            api.APIOperation = "collaboration_updateGroup"
	DELETE_GROUP            api.APIOperation = "collaboration_deleteGroup"
	CLONE_GROUP             api.APIOperation = "collaboration_cloneGroup"
	CHANGE_GROUP_OWNER      api.APIOperation = "collaboration_changeGroupOwner"
	BULK_ADD_GROUP          api.APIOperation = "collaboration_bulkAddGroup"
	BULK_UPDATE_GROUP       api.APIOperation = "collaboration_bulkUpdateGroup"
	BULK_DELETE_GROUP       api.APIOperation = "collaboration_bulkDeleteGroup"
	BULK_CHANGE_GROUP_OWNER api.APIOperation = "collaboration_bulkChangeGroupOwner"

	// READ APIs.
	LIST_GROUPS     api.APIOperation = "collaboration_listGroups"
	COUNT_GROUPS    api.APIOperation = "collaboration_countGroups"
	GET_GROUP_BY_ID api.APIOperation = "ccollaboration_getGroupById"
)
