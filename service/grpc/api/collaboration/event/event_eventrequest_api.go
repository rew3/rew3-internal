package event

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for EventRequest
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_EVENTREQUEST                   api.APIOperation = "collaboration_addEventRequest"
	UPDATE_EVENTREQUEST                 api.APIOperation = "collaboration_updateEventRequest"
	DELETE_EVENTREQUEST                 api.APIOperation = "collaboration_deleteEventRequest"
	CLONE_EVENTREQUEST                  api.APIOperation = "collaboration_cloneEventRequest"
	CHANGE_EVENTREQUEST_OWNER           api.APIOperation = "collaboration_changeEventRequestOwner"
	BULK_ADD_EVENTREQUEST          api.APIOperation = "collaboration_bulkAddEventRequest"
	BULK_UPDATE_EVENTREQUEST           api.APIOperation = "collaboration_bulkUpdateEventRequest"
	BULK_DELETE_EVENTREQUEST           api.APIOperation = "collaboration_bulkDeleteEventRequest"
	BULK_CHANGE_EVENTREQUEST_OWNER           api.APIOperation = "collaboration_bulkChangeEventRequestOwner"

	// READ APIs.
	LIST_EVENTREQUESTS     api.APIOperation = "collaboration_listEventRequests"
	COUNT_EVENTREQUESTS    api.APIOperation = "collaboration_countEventRequests"
	GET_EVENTREQUEST_BY_ID api.APIOperation = "ccollaboration_getEventRequestById"
    
)