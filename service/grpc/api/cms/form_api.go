package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_FORM                    api.APIOperation = "cms_addForm"
	UPDATE_FORM                 api.APIOperation = "cms_updateForm"
	DELETE_FORM                 api.APIOperation = "cms_deleteForm"
	CLONE_FORM                  api.APIOperation = "cms_cloneForm"
	CHANGE_FORM_OWNER           api.APIOperation = "cms_changeFormOwner"
	CHANGE_FORM_FAVORITE_STATUS api.APIOperation = "cms_changeFormFavoriteStatus"

	ADD_FORM_CATEGORY    api.APIOperation = "cms_addFormCategory"
	UPDATE_FORM_CATEGORY api.APIOperation = "cms_updateFormCategory"
	DELETE_FORM_CATEGORY api.APIOperation = "cms_deleteFormCategory"

	// READ Operations.
	LIST_FORMS     api.APIOperation = "cms_listForms"
	COUNT_FORMS    api.APIOperation = "cms_countForms"
	GET_FORM_BY_ID api.APIOperation = "cms_getFormById"

	LIST_FORM_CATEGORIES    api.APIOperation = "cms_listFormCategories"
	COUNT_FORM_CATEGORIES   api.APIOperation = "cms_countFormCategories"
	GET_FORM_CATEGORY_BY_ID api.APIOperation = "cms_getFormCategoryById"
)
