package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_POST                    endpoints.Endpoint = "cms_addPost"
	UPDATE_POST                 endpoints.Endpoint = "cms_updatePost"
	DELETE_POST                 endpoints.Endpoint = "cms_deletePost"
	CLONE_POST                  endpoints.Endpoint = "cms_clonePost"
	CHANGE_POST_OWNER           endpoints.Endpoint = "cms_changePostOwner"
	CHANGE_POST_FAVORITE_STATUS endpoints.Endpoint = "cms_changePostFavoriteStatus"

	ADD_POST_CATEGORY    endpoints.Endpoint = "cms_addPostCategory"
	UPDATE_POST_CATEGORY endpoints.Endpoint = "cms_updatePostCategory"
	DELETE_POST_CATEGORY endpoints.Endpoint = "cms_deletePostCategory"

	// READ Operations.
	LIST_POSTS     endpoints.Endpoint = "cms_listPosts"
	COUNT_POSTS    endpoints.Endpoint = "cms_countPosts"
	GET_POST_BY_ID endpoints.Endpoint = "cms_getPostById"

	LIST_POST_CATEGORIES    endpoints.Endpoint = "cms_listPostCategories"
	COUNT_POST_CATEGORIES   endpoints.Endpoint = "cms_countPostCategories"
	GET_POST_CATEGORY_BY_ID endpoints.Endpoint = "cms_getPostCategoryById"
)
