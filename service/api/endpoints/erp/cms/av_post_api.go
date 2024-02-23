package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_AVENUE_POST                    endpoints.Endpoint = "cms_addAvenuePost"
	UPDATE_AVENUE_POST                 endpoints.Endpoint = "cms_updateAvenuePost"
	DELETE_AVENUE_POST                 endpoints.Endpoint = "cms_deleteAvenuePost"
	CLONE_AVENUE_POST                  endpoints.Endpoint = "cms_cloneAvenuePost"
	CHANGE_AVENUE_POST_OWNER           endpoints.Endpoint = "cms_changeAvenuePostOwner"
	CHANGE_AVENUE_POST_FAVORITE_STATUS endpoints.Endpoint = "cms_changeAvenuePostFavoriteStatus"

	// READ Operations.
	LIST_AVENUE_POSTS     endpoints.Endpoint = "cms_listAvenuePosts"
	COUNT_AVENUE_POSTS    endpoints.Endpoint = "cms_countAvenuePosts"
	GET_AVENUE_POST_BY_ID endpoints.Endpoint = "cms_getAvenuePostById"
)
