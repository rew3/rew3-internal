package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_FORM                    endpoints.Endpoint = "cms_addForm"
	UPDATE_FORM                 endpoints.Endpoint = "cms_updateForm"
	DELETE_FORM                 endpoints.Endpoint = "cms_deleteForm"
	CLONE_FORM                  endpoints.Endpoint = "cms_cloneForm"
	CHANGE_FORM_OWNER           endpoints.Endpoint = "cms_changeFormOwner"
	CHANGE_FORM_FAVORITE_STATUS endpoints.Endpoint = "cms_changeFormFavoriteStatus"

	ADD_FORM_CATEGORY    endpoints.Endpoint = "cms_addFormCategory"
	UPDATE_FORM_CATEGORY endpoints.Endpoint = "cms_updateFormCategory"
	DELETE_FORM_CATEGORY endpoints.Endpoint = "cms_deleteFormCategory"

	// READ Operations.
	LIST_FORMS     endpoints.Endpoint = "cms_listForms"
	COUNT_FORMS    endpoints.Endpoint = "cms_countForms"
	GET_FORM_BY_ID endpoints.Endpoint = "cms_getFormById"

	LIST_FORM_CATEGORIES    endpoints.Endpoint = "cms_listFormCategories"
	COUNT_FORM_CATEGORIES   endpoints.Endpoint = "cms_countFormCategories"
	GET_FORM_CATEGORY_BY_ID endpoints.Endpoint = "cms_getFormCategoryById"
)
