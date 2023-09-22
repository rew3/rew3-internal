package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_SHOWING                    api.APIOperation = "addShowing"
	UPDATE_SHOWING                 api.APIOperation = "updateShowing"
	DELETE_SHOWING                 api.APIOperation = "deleteShowing"
	CLONE_SHOWING                  api.APIOperation = "cloneShowing"
	CHANGE_SHOWING_OWNER           api.APIOperation = "changeShowingOwner"

	// READ Operations.
	LIST_SHOWINGS     api.APIOperation = "listShowings"
	COUNT_SHOWINGS    api.APIOperation = "countShowings"
	GET_SHOWING_BY_ID api.APIOperation = "getShowingById"
)
