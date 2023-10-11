package listview

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ListView
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_LISTVIEW    api.APIOperation = "listview_addListView"
	UPDATE_LISTVIEW api.APIOperation = "listview_updateListView"
	DELETE_LISTVIEW api.APIOperation = "listview_deleteListView"
	CLONE_LISTVIEW  api.APIOperation = "listview_cloneListView"

	// READ APIs.
	LIST_LISTVIEWS     api.APIOperation = "listview_listListViews"
	COUNT_LISTVIEWS    api.APIOperation = "listview_countListViews"
	GET_LISTVIEW_BY_ID api.APIOperation = "clistview_getListViewById"
)
