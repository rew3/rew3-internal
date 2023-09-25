package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_CASE                    api.APIOperation = "crm_addCase"
	UPDATE_CASE                 api.APIOperation = "crm_updateCase"
	DELETE_CASE                 api.APIOperation = "crm_deleteCase"
	CLONE_CASE                  api.APIOperation = "crm_cloneCase"
	CHANGE_CASE_OWNER           api.APIOperation = "crm_changeCaseOwner"
	CHANGE_CASE_FAVORITE_STATUS api.APIOperation = "crm_changeCaseFavoriteStatus"

	// READ Operations.
	LIST_CASES     api.APIOperation = "crm_listCases"
	COUNT_CASES    api.APIOperation = "crm_countCases"
	GET_CASE_BY_ID api.APIOperation = "crm_getCaseById"
)
