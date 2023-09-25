package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_AVENUE_POST                    api.APIOperation = "cms_addAvenuePost"
	UPDATE_AVENUE_POST                 api.APIOperation = "cms_updateAvenuePost"
	DELETE_AVENUE_POST                 api.APIOperation = "cms_deleteAvenuePost"
	CLONE_AVENUE_POST                  api.APIOperation = "cms_cloneAvenuePost"
	CHANGE_AVENUE_POST_OWNER           api.APIOperation = "cms_changeAvenuePostOwner"
	CHANGE_AVENUE_POST_FAVORITE_STATUS api.APIOperation = "cms_changeAvenuePostFavoriteStatus"

	// READ Operations.
	LIST_AVENUE_POSTS     api.APIOperation = "cms_listAvenuePosts"
	COUNT_AVENUE_POSTS    api.APIOperation = "cms_countAvenuePosts"
	GET_AVENUE_POST_BY_ID api.APIOperation = "cms_getAvenuePostById"
)