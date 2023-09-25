package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_POST                    api.APIOperation = "cms_addPost"
	UPDATE_POST                 api.APIOperation = "cms_updatePost"
	DELETE_POST                 api.APIOperation = "cms_deletePost"
	CLONE_POST                  api.APIOperation = "cms_clonePost"
	CHANGE_POST_OWNER           api.APIOperation = "cms_changePostOwner"
	CHANGE_POST_FAVORITE_STATUS api.APIOperation = "cms_changePostFavoriteStatus"

	ADD_POST_CATEGORY    api.APIOperation = "cms_addPostCategory"
	UPDATE_POST_CATEGORY api.APIOperation = "cms_updatePostCategory"
	DELETE_POST_CATEGORY api.APIOperation = "cms_deletePostCategory"

	// READ Operations.
	LIST_POSTS     api.APIOperation = "cms_listPosts"
	COUNT_POSTS    api.APIOperation = "cms_countPosts"
	GET_POST_BY_ID api.APIOperation = "cms_getPostById"

	LIST_POST_CATEGORIES    api.APIOperation = "cms_listPostCategories"
	COUNT_POST_CATEGORIES   api.APIOperation = "cms_countPostCategories"
	GET_POST_CATEGORY_BY_ID api.APIOperation = "cms_getPostCategoryById"
)
