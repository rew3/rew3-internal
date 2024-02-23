package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_BLOG                    endpoints.Endpoint = "cms_addBlog"
	UPDATE_BLOG                 endpoints.Endpoint = "cms_updateBlog"
	DELETE_BLOG                 endpoints.Endpoint = "cms_deleteBlog"
	CLONE_BLOG                  endpoints.Endpoint = "cms_cloneBlog"
	CHANGE_BLOG_OWNER           endpoints.Endpoint = "cms_changeBlogOwner"
	CHANGE_BLOG_FAVORITE_STATUS endpoints.Endpoint = "cms_changeBlogFavoriteStatus"

	ADD_BLOG_CATEGORY    endpoints.Endpoint = "cms_addBlogCategory"
	UPDATE_BLOG_CATEGORY endpoints.Endpoint = "cms_updateBlogCategory"
	DELETE_BLOG_CATEGORY endpoints.Endpoint = "cms_deleteBlogCategory"

	// READ Operations.
	LIST_BLOGS     endpoints.Endpoint = "cms_listBlogs"
	COUNT_BLOGS    endpoints.Endpoint = "cms_countBlogs"
	GET_BLOG_BY_ID endpoints.Endpoint = "cms_getBlogById"

	LIST_BLOG_CATEGORIES    endpoints.Endpoint = "cms_listBlogCategories"
	COUNT_BLOG_CATEGORIES   endpoints.Endpoint = "cms_countBlogCategories"
	GET_BLOG_CATEGORY_BY_ID endpoints.Endpoint = "cms_getBlogCategoryById"
)
