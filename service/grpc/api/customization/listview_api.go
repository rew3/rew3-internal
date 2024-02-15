package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ListView
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_LISTVIEW    api.APIOperation = "customization_addListView"
	UPDATE_LISTVIEW api.APIOperation = "customization_updateListView"
	DELETE_LISTVIEW api.APIOperation = "customization_deleteListView"

	// READ APIs.
	LIST_LISTVIEWS     api.APIOperation = "customization_listListViews"
	COUNT_LISTVIEWS    api.APIOperation = "customization_countListViews"
	GET_LISTVIEW_BY_ID api.APIOperation = "ccustomization_getListViewById"
)
