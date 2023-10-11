package listview

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ViewPreference
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_VIEWPREFERENCE    api.APIOperation = "listview_addViewPreference"
	UPDATE_VIEWPREFERENCE api.APIOperation = "listview_updateViewPreference"
	DELETE_VIEWPREFERENCE api.APIOperation = "listview_deleteViewPreference"

	// READ APIs.
	LIST_VIEWPREFERENCES     api.APIOperation = "listview_listViewPreferences"
	COUNT_VIEWPREFERENCES    api.APIOperation = "listview_countViewPreferences"
	GET_VIEWPREFERENCE_BY_ID api.APIOperation = "clistview_getViewPreferenceById"
)
