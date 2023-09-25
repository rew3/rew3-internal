package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_REQUEST          api.APIOperation = "rms_addRequest"
	UPDATE_REQUEST       api.APIOperation = "rms_updateRequest"
	DELETE_REQUEST       api.APIOperation = "rms_deleteRequest"
	CLONE_REQUEST        api.APIOperation = "rms_cloneRequest"
	CHANGE_REQUEST_OWNER api.APIOperation = "rms_changeRequestOwner"

	// READ Operations.
	LIST_REQUESTS     api.APIOperation = "rms_listRequests"
	COUNT_REQUESTS    api.APIOperation = "rms_countRequests"
	GET_REQUEST_BY_ID api.APIOperation = "rms_getRequestById"
)
