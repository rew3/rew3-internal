package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_SHOWING          api.APIOperation = "rms_addShowing"
	UPDATE_SHOWING       api.APIOperation = "rms_updateShowing"
	DELETE_SHOWING       api.APIOperation = "rms_deleteShowing"
	CLONE_SHOWING        api.APIOperation = "rms_cloneShowing"
	CHANGE_SHOWING_OWNER api.APIOperation = "rms_changeShowingOwner"

	// READ Operations.
	LIST_SHOWINGS     api.APIOperation = "rms_listShowings"
	COUNT_SHOWINGS    api.APIOperation = "rms_countShowings"
	GET_SHOWING_BY_ID api.APIOperation = "rms_getShowingById"
)
