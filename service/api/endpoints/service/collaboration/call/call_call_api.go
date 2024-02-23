package call

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Call
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CALL               endpoints.Endpoint = "collaboration_addCall"
	UPDATE_CALL            endpoints.Endpoint = "collaboration_updateCall"
	DELETE_CALL            endpoints.Endpoint = "collaboration_deleteCall"
	CLONE_CALL             endpoints.Endpoint = "collaboration_cloneCall"
	CHANGE_CALL_OWNER      endpoints.Endpoint = "collaboration_changeCallOwner"
	BULK_ADD_CALL          endpoints.Endpoint = "collaboration_bulkAddCall"
	BULK_UPDATE_CALL       endpoints.Endpoint = "collaboration_bulkUpdateCall"
	BULK_DELETE_CALL       endpoints.Endpoint = "collaboration_bulkDeleteCall"
	BULK_CHANGE_CALL_OWNER endpoints.Endpoint = "collaboration_bulkChangeCallOwner"

	// READ APIs.
	LIST_CALLS     endpoints.Endpoint = "collaboration_listCalls"
	COUNT_CALLS    endpoints.Endpoint = "collaboration_countCalls"
	GET_CALL_BY_ID endpoints.Endpoint = "ccollaboration_getCallById"
)
