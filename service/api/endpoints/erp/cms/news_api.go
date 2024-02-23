package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_NEWS                    endpoints.Endpoint = "cms_addNews"
	UPDATE_NEWS                 endpoints.Endpoint = "cms_updateNews"
	DELETE_NEWS                 endpoints.Endpoint = "cms_deleteNews"
	CLONE_NEWS                  endpoints.Endpoint = "cms_cloneNews"
	CHANGE_NEWS_OWNER           endpoints.Endpoint = "cms_changeNewsOwner"
	CHANGE_NEWS_FAVORITE_STATUS endpoints.Endpoint = "cms_changeNewsFavoriteStatus"

	ADD_NEWS_CATEGORY    endpoints.Endpoint = "cms_addNewsCategory"
	UPDATE_NEWS_CATEGORY endpoints.Endpoint = "cms_updateNewsCategory"
	DELETE_NEWS_CATEGORY endpoints.Endpoint = "cms_deleteNewsCategory"

	// READ Operations.
	LIST_NEWS      endpoints.Endpoint = "cms_listNews"
	COUNT_NEWS     endpoints.Endpoint = "cms_countNews"
	GET_NEWS_BY_ID endpoints.Endpoint = "cms_getNewsById"

	LIST_NEWS_CATEGORIES    endpoints.Endpoint = "cms_listNewsCategories"
	COUNT_NEWS_CATEGORIES   endpoints.Endpoint = "cms_countNewsCategories"
	GET_NEWS_CATEGORY_BY_ID endpoints.Endpoint = "cms_getNewsCategoryById"
)
