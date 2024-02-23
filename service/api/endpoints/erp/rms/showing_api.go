package rms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_SHOWING          endpoints.Endpoint = "rms_addShowing"
	UPDATE_SHOWING       endpoints.Endpoint = "rms_updateShowing"
	DELETE_SHOWING       endpoints.Endpoint = "rms_deleteShowing"
	CLONE_SHOWING        endpoints.Endpoint = "rms_cloneShowing"
	CHANGE_SHOWING_OWNER endpoints.Endpoint = "rms_changeShowingOwner"
	SET_SHOWING_FAVORITE endpoints.Endpoint = "crm_setShowingFavorite"

	// READ Operations.
	LIST_SHOWINGS     endpoints.Endpoint = "rms_listShowings"
	COUNT_SHOWINGS    endpoints.Endpoint = "rms_countShowings"
	GET_SHOWING_BY_ID endpoints.Endpoint = "rms_getShowingById"
)
