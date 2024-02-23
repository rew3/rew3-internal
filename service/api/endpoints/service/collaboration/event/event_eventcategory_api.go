package event

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for EventCategory
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EVENTCATEGORY               endpoints.Endpoint = "collaboration_addEventCategory"
	UPDATE_EVENTCATEGORY            endpoints.Endpoint = "collaboration_updateEventCategory"
	DELETE_EVENTCATEGORY            endpoints.Endpoint = "collaboration_deleteEventCategory"
	CLONE_EVENTCATEGORY             endpoints.Endpoint = "collaboration_cloneEventCategory"
	CHANGE_EVENTCATEGORY_OWNER      endpoints.Endpoint = "collaboration_changeEventCategoryOwner"
	BULK_ADD_EVENTCATEGORY          endpoints.Endpoint = "collaboration_bulkAddEventCategory"
	BULK_UPDATE_EVENTCATEGORY       endpoints.Endpoint = "collaboration_bulkUpdateEventCategory"
	BULK_DELETE_EVENTCATEGORY       endpoints.Endpoint = "collaboration_bulkDeleteEventCategory"
	BULK_CHANGE_EVENTCATEGORY_OWNER endpoints.Endpoint = "collaboration_bulkChangeEventCategoryOwner"

	// READ APIs.
	LIST_EVENTCATEGORIES    endpoints.Endpoint = "collaboration_listEventCategories"
	COUNT_EVENTCATEGORIES   endpoints.Endpoint = "collaboration_countEventCategories"
	GET_EVENTCATEGORY_BY_ID endpoints.Endpoint = "ccollaboration_getEventCategoryById"
)
