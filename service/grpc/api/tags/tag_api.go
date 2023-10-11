package tags

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Tag
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TAG         api.APIOperation = "tags_addTag"
	UPDATE_TAG      api.APIOperation = "tags_updateTag"
	DELETE_TAG      api.APIOperation = "tags_deleteTag"
	CLONE_TAG       api.APIOperation = "tags_cloneTag"
	BULK_ADD_TAG    api.APIOperation = "tags_bulkAddTag"
	BULK_UPDATE_TAG api.APIOperation = "tags_bulkUpdateTag"
	BULK_DELETE_TAG api.APIOperation = "tags_bulkDeleteTag"

	// READ APIs.
	LIST_TAGS     api.APIOperation = "tags_listTags"
	COUNT_TAGS    api.APIOperation = "tags_countTags"
	GET_TAG_BY_ID api.APIOperation = "ctags_getTagById"
)
