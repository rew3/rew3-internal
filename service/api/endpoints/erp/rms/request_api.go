package rms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_REQUEST          endpoints.Endpoint = "rms_addRequest"
	UPDATE_REQUEST       endpoints.Endpoint = "rms_updateRequest"
	DELETE_REQUEST       endpoints.Endpoint = "rms_deleteRequest"
	CLONE_REQUEST        endpoints.Endpoint = "rms_cloneRequest"
	CHANGE_REQUEST_OWNER endpoints.Endpoint = "rms_changeRequestOwner"
	SET_REQUEST_FAVORITE endpoints.Endpoint = "crm_setRequestFavorite"

	// READ Operations.
	LIST_REQUESTS     endpoints.Endpoint = "rms_listRequests"
	COUNT_REQUESTS    endpoints.Endpoint = "rms_countRequests"
	GET_REQUEST_BY_ID endpoints.Endpoint = "rms_getRequestById"
)
