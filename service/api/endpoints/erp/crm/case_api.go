package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_CASE          endpoints.Endpoint = "crm_addCase"
	UPDATE_CASE       endpoints.Endpoint = "crm_updateCase"
	DELETE_CASE       endpoints.Endpoint = "crm_deleteCase"
	CLONE_CASE        endpoints.Endpoint = "crm_cloneCase"
	CHANGE_CASE_OWNER endpoints.Endpoint = "crm_changeCaseOwner"
	SET_CASE_FAVORITE endpoints.Endpoint = "crm_setCaseFavorite"

	// READ Operations.
	LIST_CASES     endpoints.Endpoint = "crm_listCases"
	COUNT_CASES    endpoints.Endpoint = "crm_countCases"
	GET_CASE_BY_ID endpoints.Endpoint = "crm_getCaseById"
)
