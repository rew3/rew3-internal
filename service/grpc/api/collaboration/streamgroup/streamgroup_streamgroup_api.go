package streamgroup

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for StreamGroup
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_STREAMGROUP               api.APIOperation = "collaboration_addStreamGroup"
	UPDATE_STREAMGROUP            api.APIOperation = "collaboration_updateStreamGroup"
	DELETE_STREAMGROUP            api.APIOperation = "collaboration_deleteStreamGroup"
	CLONE_STREAMGROUP             api.APIOperation = "collaboration_cloneStreamGroup"
	CHANGE_STREAMGROUP_OWNER      api.APIOperation = "collaboration_changeStreamGroupOwner"
	BULK_ADD_STREAMGROUP          api.APIOperation = "collaboration_bulkAddStreamGroup"
	BULK_UPDATE_STREAMGROUP       api.APIOperation = "collaboration_bulkUpdateStreamGroup"
	BULK_DELETE_STREAMGROUP       api.APIOperation = "collaboration_bulkDeleteStreamGroup"
	BULK_CHANGE_STREAMGROUP_OWNER api.APIOperation = "collaboration_bulkChangeStreamGroupOwner"

	// READ APIs.
	LIST_STREAMGROUPS     api.APIOperation = "collaboration_listStreamGroups"
	COUNT_STREAMGROUPS    api.APIOperation = "collaboration_countStreamGroups"
	GET_STREAMGROUP_BY_ID api.APIOperation = "ccollaboration_getStreamGroupById"
)
