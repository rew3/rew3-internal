package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_CASE                    api.APIOperation = "addCase"
	UPDATE_CASE                 api.APIOperation = "updateCase"
	DELETE_CASE                 api.APIOperation = "deleteCase"
	CLONE_CASE                  api.APIOperation = "cloneCase"
	CHANGE_CASE_OWNER           api.APIOperation = "changeOwner"
	CHANGE_CASE_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	// READ Operations.
	LIST_CASES     api.APIOperation = "listCases"
	COUNT_CASES    api.APIOperation = "countCases"
	GET_CASE_BY_ID api.APIOperation = "getCaseById"
)
