package event

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for EventCategory
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_EVENTCATEGORY                   api.APIOperation = "collaboration_addEventCategory"
	UPDATE_EVENTCATEGORY                 api.APIOperation = "collaboration_updateEventCategory"
	DELETE_EVENTCATEGORY                 api.APIOperation = "collaboration_deleteEventCategory"
	CLONE_EVENTCATEGORY                  api.APIOperation = "collaboration_cloneEventCategory"
	CHANGE_EVENTCATEGORY_OWNER           api.APIOperation = "collaboration_changeEventCategoryOwner"
	BULK_ADD_EVENTCATEGORY          api.APIOperation = "collaboration_bulkAddEventCategory"
	BULK_UPDATE_EVENTCATEGORY           api.APIOperation = "collaboration_bulkUpdateEventCategory"
	BULK_DELETE_EVENTCATEGORY           api.APIOperation = "collaboration_bulkDeleteEventCategory"
	BULK_CHANGE_EVENTCATEGORY_OWNER           api.APIOperation = "collaboration_bulkChangeEventCategoryOwner"

	// READ APIs.
	LIST_EVENTCATEGORIES     api.APIOperation = "collaboration_listEventCategories"
	COUNT_EVENTCATEGORIES    api.APIOperation = "collaboration_countEventCategories"
	GET_EVENTCATEGORY_BY_ID api.APIOperation = "ccollaboration_getEventCategoryById"
    
)