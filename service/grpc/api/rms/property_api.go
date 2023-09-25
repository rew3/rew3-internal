package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_PROPERTY          api.APIOperation = "rms_addProperty"
	UPDATE_PROPERTY       api.APIOperation = "rms_updateProperty"
	DELETE_PROPERTY       api.APIOperation = "rms_deleteProperty"
	CLONE_PROPERTY        api.APIOperation = "rms_cloneProperty"
	CHANGE_PROPERTY_OWNER api.APIOperation = "rms_changePropertyOwner"

	// READ Operations.
	LIST_PROPERTIES    api.APIOperation = "rms_listProperties"
	COUNT_PROPERTIES   api.APIOperation = "rms_countProperties"
	GET_PROPERTY_BY_ID api.APIOperation = "rms_getPropertyById"
)
