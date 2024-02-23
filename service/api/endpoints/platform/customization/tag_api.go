package customization

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Tag
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TAG    endpoints.Endpoint = "customization_addTag"
	UPDATE_TAG endpoints.Endpoint = "customization_updateTag"
	DELETE_TAG endpoints.Endpoint = "customization_deleteTag"

	// READ APIs.
	LIST_TAGS     endpoints.Endpoint = "customization_listTags"
	COUNT_TAGS    endpoints.Endpoint = "customization_countTags"
	GET_TAG_BY_ID endpoints.Endpoint = "ccustomization_getTagById"
)
