package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_BLOG                    api.APIOperation = "cms_addBlog"
	UPDATE_BLOG                 api.APIOperation = "cms_updateBlog"
	DELETE_BLOG                 api.APIOperation = "cms_deleteBlog"
	CLONE_BLOG                  api.APIOperation = "cms_cloneBlog"
	CHANGE_BLOG_OWNER           api.APIOperation = "cms_changeBlogOwner"
	CHANGE_BLOG_FAVORITE_STATUS api.APIOperation = "cms_changeBlogFavoriteStatus"

	ADD_BLOG_CATEGORY    api.APIOperation = "cms_addBlogCategory"
	UPDATE_BLOG_CATEGORY api.APIOperation = "cms_updateBlogCategory"
	DELETE_BLOG_CATEGORY api.APIOperation = "cms_deleteBlogCategory"

	// READ Operations.
	LIST_BLOGS     api.APIOperation = "cms_listBlogs"
	COUNT_BLOGS    api.APIOperation = "cms_countBlogs"
	GET_BLOG_BY_ID api.APIOperation = "cms_getBlogById"

	LIST_BLOG_CATEGORIES    api.APIOperation = "cms_listBlogCategories"
	COUNT_BLOG_CATEGORIES   api.APIOperation = "cms_countBlogCategories"
	GET_BLOG_CATEGORY_BY_ID api.APIOperation = "cms_getBlogCategoryById"
)
