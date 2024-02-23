package rms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_PROPERTY          endpoints.Endpoint = "rms_addProperty"
	UPDATE_PROPERTY       endpoints.Endpoint = "rms_updateProperty"
	DELETE_PROPERTY       endpoints.Endpoint = "rms_deleteProperty"
	CLONE_PROPERTY        endpoints.Endpoint = "rms_cloneProperty"
	CHANGE_PROPERTY_OWNER endpoints.Endpoint = "rms_changePropertyOwner"
	SET_PROPERTY_FAVORITE endpoints.Endpoint = "crm_setPropertyFavorite"

	// READ Operations.
	LIST_PROPERTIES    endpoints.Endpoint = "rms_listProperties"
	COUNT_PROPERTIES   endpoints.Endpoint = "rms_countProperties"
	GET_PROPERTY_BY_ID endpoints.Endpoint = "rms_getPropertyById"
)
