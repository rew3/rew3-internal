package call

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Call
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_CALL                   api.APIOperation = "collaboration_addCall"
	UPDATE_CALL                 api.APIOperation = "collaboration_updateCall"
	DELETE_CALL                 api.APIOperation = "collaboration_deleteCall"
	CLONE_CALL                  api.APIOperation = "collaboration_cloneCall"
	CHANGE_CALL_OWNER           api.APIOperation = "collaboration_changeCallOwner"
	BULK_ADD_CALL          api.APIOperation = "collaboration_bulkAddCall"
	BULK_UPDATE_CALL           api.APIOperation = "collaboration_bulkUpdateCall"
	BULK_DELETE_CALL           api.APIOperation = "collaboration_bulkDeleteCall"
	BULK_CHANGE_CALL_OWNER           api.APIOperation = "collaboration_bulkChangeCallOwner"

	// READ APIs.
	LIST_CALLS     api.APIOperation = "collaboration_listCalls"
	COUNT_CALLS    api.APIOperation = "collaboration_countCalls"
	GET_CALL_BY_ID api.APIOperation = "ccollaboration_getCallById"
    
)