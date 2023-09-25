package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_FORUM                    api.APIOperation = "cms_addForum"
	UPDATE_FORUM                 api.APIOperation = "cms_updateForum"
	DELETE_FORUM                 api.APIOperation = "cms_deleteForum"
	CLONE_FORUM                  api.APIOperation = "cms_cloneForum"
	CHANGE_FORUM_OWNER           api.APIOperation = "cms_changeForumOwner"
	CHANGE_FORUM_FAVORITE_STATUS api.APIOperation = "cms_changeForumFavoriteStatus"

	ADD_FORUM_CATEGORY    api.APIOperation = "cms_addForumCategory"
	UPDATE_FORUM_CATEGORY api.APIOperation = "cms_updateForumCategory"
	DELETE_FORUM_CATEGORY api.APIOperation = "cms_deleteForumCategory"

	// READ Operations.
	LIST_FORUMS     api.APIOperation = "cms_listForums"
	COUNT_FORUMS    api.APIOperation = "cms_countForums"
	GET_FORUM_BY_ID api.APIOperation = "cms_getForumById"

	LIST_FORUM_CATEGORIES    api.APIOperation = "cms_listForumCategories"
	COUNT_FORUM_CATEGORIES   api.APIOperation = "cms_countForumCategories"
	GET_FORUM_CATEGORY_BY_ID api.APIOperation = "cms_getForumCategoryById"
)
