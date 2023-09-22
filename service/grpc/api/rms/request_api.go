package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_REQUEST                    api.APIOperation = "addRequest"
	UPDATE_REQUEST                 api.APIOperation = "updateRequest"
	DELETE_REQUEST                 api.APIOperation = "deleteRequest"
	CLONE_REQUEST                  api.APIOperation = "cloneRequest"
	CHANGE_REQUEST_OWNER           api.APIOperation = "changeRequestOwner"

	// READ Operations.
	LIST_REQUESTS     api.APIOperation = "listRequests"
	COUNT_REQUESTS    api.APIOperation = "countRequests"
	GET_REQUEST_BY_ID api.APIOperation = "getRequestById"
)
