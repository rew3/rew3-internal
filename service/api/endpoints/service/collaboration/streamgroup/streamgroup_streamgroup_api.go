package streamgroup

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for StreamGroup
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_STREAMGROUP               endpoints.Endpoint = "collaboration_addStreamGroup"
	UPDATE_STREAMGROUP            endpoints.Endpoint = "collaboration_updateStreamGroup"
	DELETE_STREAMGROUP            endpoints.Endpoint = "collaboration_deleteStreamGroup"
	CLONE_STREAMGROUP             endpoints.Endpoint = "collaboration_cloneStreamGroup"
	CHANGE_STREAMGROUP_OWNER      endpoints.Endpoint = "collaboration_changeStreamGroupOwner"
	BULK_ADD_STREAMGROUP          endpoints.Endpoint = "collaboration_bulkAddStreamGroup"
	BULK_UPDATE_STREAMGROUP       endpoints.Endpoint = "collaboration_bulkUpdateStreamGroup"
	BULK_DELETE_STREAMGROUP       endpoints.Endpoint = "collaboration_bulkDeleteStreamGroup"
	BULK_CHANGE_STREAMGROUP_OWNER endpoints.Endpoint = "collaboration_bulkChangeStreamGroupOwner"

	// READ APIs.
	LIST_STREAMGROUPS     endpoints.Endpoint = "collaboration_listStreamGroups"
	COUNT_STREAMGROUPS    endpoints.Endpoint = "collaboration_countStreamGroups"
	GET_STREAMGROUP_BY_ID endpoints.Endpoint = "ccollaboration_getStreamGroupById"
)
