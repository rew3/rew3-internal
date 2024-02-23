package calendar

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for CalendarCategory
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CALENDARCATEGORY               endpoints.Endpoint = "collaboration_addCalendarCategory"
	UPDATE_CALENDARCATEGORY            endpoints.Endpoint = "collaboration_updateCalendarCategory"
	DELETE_CALENDARCATEGORY            endpoints.Endpoint = "collaboration_deleteCalendarCategory"
	CLONE_CALENDARCATEGORY             endpoints.Endpoint = "collaboration_cloneCalendarCategory"
	CHANGE_CALENDARCATEGORY_OWNER      endpoints.Endpoint = "collaboration_changeCalendarCategoryOwner"
	BULK_ADD_CALENDARCATEGORY          endpoints.Endpoint = "collaboration_bulkAddCalendarCategory"
	BULK_UPDATE_CALENDARCATEGORY       endpoints.Endpoint = "collaboration_bulkUpdateCalendarCategory"
	BULK_DELETE_CALENDARCATEGORY       endpoints.Endpoint = "collaboration_bulkDeleteCalendarCategory"
	BULK_CHANGE_CALENDARCATEGORY_OWNER endpoints.Endpoint = "collaboration_bulkChangeCalendarCategoryOwner"

	// READ APIs.
	LIST_CALENDARCATEGORIES    endpoints.Endpoint = "collaboration_listCalendarCategories"
	COUNT_CALENDARCATEGORIES   endpoints.Endpoint = "collaboration_countCalendarCategories"
	GET_CALENDARCATEGORY_BY_ID endpoints.Endpoint = "ccollaboration_getCalendarCategoryById"
)
