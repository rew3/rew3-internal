package customization

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for ListView
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_LISTVIEW            endpoints.Endpoint = "customization_addListView"
	UPDATE_LISTVIEW         endpoints.Endpoint = "customization_updateListView"
	DELETE_LISTVIEW         endpoints.Endpoint = "customization_deleteListView"
	UPDATE_LISTVIEW_FILTERS endpoints.Endpoint = "customization_updateListViewFilters"
	SET_LISTVIEW_ACTIVE     endpoints.Endpoint = "customization_setListViewActive"

	// READ APIs.
	LIST_LISTVIEWS     endpoints.Endpoint = "customization_listListViews"
	COUNT_LISTVIEWS    endpoints.Endpoint = "customization_countListViews"
	GET_LISTVIEW_BY_ID endpoints.Endpoint = "ccustomization_getListViewById"
)
