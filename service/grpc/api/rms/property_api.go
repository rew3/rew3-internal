package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_PROPERTY                    api.APIOperation = "addProperty"
	UPDATE_PROPERTY                 api.APIOperation = "updateProperty"
	DELETE_PROPERTY                 api.APIOperation = "deleteProperty"
	CLONE_PROPERTY                  api.APIOperation = "cloneProperty"
	CHANGE_PROPERTY_OWNER           api.APIOperation = "changePropertyOwner"

	// READ Operations.
	LIST_PROPERTYS     api.APIOperation = "listPropertys"
	COUNT_PROPERTYS    api.APIOperation = "countPropertys"
	GET_PROPERTY_BY_ID api.APIOperation = "getPropertyById"
)
