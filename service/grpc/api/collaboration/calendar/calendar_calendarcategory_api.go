package calendar

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for CalendarCategory
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_CALENDARCATEGORY                   api.APIOperation = "collaboration_addCalendarCategory"
	UPDATE_CALENDARCATEGORY                 api.APIOperation = "collaboration_updateCalendarCategory"
	DELETE_CALENDARCATEGORY                 api.APIOperation = "collaboration_deleteCalendarCategory"
	CLONE_CALENDARCATEGORY                  api.APIOperation = "collaboration_cloneCalendarCategory"
	CHANGE_CALENDARCATEGORY_OWNER           api.APIOperation = "collaboration_changeCalendarCategoryOwner"
	BULK_ADD_CALENDARCATEGORY          api.APIOperation = "collaboration_bulkAddCalendarCategory"
	BULK_UPDATE_CALENDARCATEGORY           api.APIOperation = "collaboration_bulkUpdateCalendarCategory"
	BULK_DELETE_CALENDARCATEGORY           api.APIOperation = "collaboration_bulkDeleteCalendarCategory"
	BULK_CHANGE_CALENDARCATEGORY_OWNER           api.APIOperation = "collaboration_bulkChangeCalendarCategoryOwner"

	// READ APIs.
	LIST_CALENDARCATEGORIES     api.APIOperation = "collaboration_listCalendarCategories"
	COUNT_CALENDARCATEGORIES    api.APIOperation = "collaboration_countCalendarCategories"
	GET_CALENDARCATEGORY_BY_ID api.APIOperation = "ccollaboration_getCalendarCategoryById"
    
)