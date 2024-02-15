package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ViewPreference
	// ----------------------------------------------------
	// WRITE APIs.
	UPDATE_VIEWPREFERENCE api.APIOperation = "customization_updateViewPreference"

	// READ APIs.
	GET_VIEWPREFERENCE_BY_ENTITY api.APIOperation = "ccustomization_getViewPreferenceByEntity"
)
