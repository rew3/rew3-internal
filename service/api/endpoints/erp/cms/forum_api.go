package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_FORUM                    endpoints.Endpoint = "cms_addForum"
	UPDATE_FORUM                 endpoints.Endpoint = "cms_updateForum"
	DELETE_FORUM                 endpoints.Endpoint = "cms_deleteForum"
	CLONE_FORUM                  endpoints.Endpoint = "cms_cloneForum"
	CHANGE_FORUM_OWNER           endpoints.Endpoint = "cms_changeForumOwner"
	CHANGE_FORUM_FAVORITE_STATUS endpoints.Endpoint = "cms_changeForumFavoriteStatus"

	ADD_FORUM_CATEGORY    endpoints.Endpoint = "cms_addForumCategory"
	UPDATE_FORUM_CATEGORY endpoints.Endpoint = "cms_updateForumCategory"
	DELETE_FORUM_CATEGORY endpoints.Endpoint = "cms_deleteForumCategory"

	// READ Operations.
	LIST_FORUMS     endpoints.Endpoint = "cms_listForums"
	COUNT_FORUMS    endpoints.Endpoint = "cms_countForums"
	GET_FORUM_BY_ID endpoints.Endpoint = "cms_getForumById"

	LIST_FORUM_CATEGORIES    endpoints.Endpoint = "cms_listForumCategories"
	COUNT_FORUM_CATEGORIES   endpoints.Endpoint = "cms_countForumCategories"
	GET_FORUM_CATEGORY_BY_ID endpoints.Endpoint = "cms_getForumCategoryById"
)
