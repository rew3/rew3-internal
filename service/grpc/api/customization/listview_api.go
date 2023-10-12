package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ListView
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_LISTVIEW               api.APIOperation = "customization_addListView"
	UPDATE_LISTVIEW            api.APIOperation = "customization_updateListView"
	DELETE_LISTVIEW            api.APIOperation = "customization_deleteListView"
	CLONE_LISTVIEW             api.APIOperation = "customization_cloneListView"
	CHANGE_LISTVIEW_OWNER      api.APIOperation = "customization_changeListViewOwner"
	BULK_ADD_LISTVIEW          api.APIOperation = "customization_bulkAddListView"
	BULK_UPDATE_LISTVIEW       api.APIOperation = "customization_bulkUpdateListView"
	BULK_DELETE_LISTVIEW       api.APIOperation = "customization_bulkDeleteListView"
	BULK_CHANGE_LISTVIEW_OWNER api.APIOperation = "customization_bulkChangeListViewOwner"

	// READ APIs.
	LIST_LISTVIEWS     api.APIOperation = "customization_listListViews"
	COUNT_LISTVIEWS    api.APIOperation = "customization_countListViews"
	GET_LISTVIEW_BY_ID api.APIOperation = "ccustomization_getListViewById"
)
