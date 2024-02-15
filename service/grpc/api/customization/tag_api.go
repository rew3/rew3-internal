package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Tag
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TAG    api.APIOperation = "customization_addTag"
	UPDATE_TAG api.APIOperation = "customization_updateTag"
	DELETE_TAG api.APIOperation = "customization_deleteTag"

	// READ APIs.
	LIST_TAGS     api.APIOperation = "customization_listTags"
	COUNT_TAGS    api.APIOperation = "customization_countTags"
	GET_TAG_BY_ID api.APIOperation = "ccustomization_getTagById"
)
