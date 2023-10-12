package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ViewPreference
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_VIEWPREFERENCE    api.APIOperation = "customization_addViewPreference"
	UPDATE_VIEWPREFERENCE api.APIOperation = "customization_updateViewPreference"
	DELETE_VIEWPREFERENCE api.APIOperation = "customization_deleteViewPreference"

	// READ APIs.
	LIST_VIEWPREFERENCES     api.APIOperation = "customization_listViewPreferences"
	COUNT_VIEWPREFERENCES    api.APIOperation = "customization_countViewPreferences"
	GET_VIEWPREFERENCE_BY_ID api.APIOperation = "ccustomization_getViewPreferenceById"
)
