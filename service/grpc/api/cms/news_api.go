package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_NEWS                    api.APIOperation = "cms_addNews"
	UPDATE_NEWS                 api.APIOperation = "cms_updateNews"
	DELETE_NEWS                 api.APIOperation = "cms_deleteNews"
	CLONE_NEWS                  api.APIOperation = "cms_cloneNews"
	CHANGE_NEWS_OWNER           api.APIOperation = "cms_changeNewsOwner"
	CHANGE_NEWS_FAVORITE_STATUS api.APIOperation = "cms_changeNewsFavoriteStatus"

	ADD_NEWS_CATEGORY    api.APIOperation = "cms_addNewsCategory"
	UPDATE_NEWS_CATEGORY api.APIOperation = "cms_updateNewsCategory"
	DELETE_NEWS_CATEGORY api.APIOperation = "cms_deleteNewsCategory"

	// READ Operations.
	LIST_NEWS      api.APIOperation = "cms_listNews"
	COUNT_NEWS     api.APIOperation = "cms_countNews"
	GET_NEWS_BY_ID api.APIOperation = "cms_getNewsById"

	LIST_NEWS_CATEGORIES    api.APIOperation = "cms_listNewsCategories"
	COUNT_NEWS_CATEGORIES   api.APIOperation = "cms_countNewsCategories"
	GET_NEWS_CATEGORY_BY_ID api.APIOperation = "cms_getNewsCategoryById"
)
