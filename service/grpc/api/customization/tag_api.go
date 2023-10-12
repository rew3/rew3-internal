package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Tag
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TAG               api.APIOperation = "customization_addTag"
	UPDATE_TAG            api.APIOperation = "customization_updateTag"
	DELETE_TAG            api.APIOperation = "customization_deleteTag"
	CLONE_TAG             api.APIOperation = "customization_cloneTag"
	CHANGE_TAG_OWNER      api.APIOperation = "customization_changeTagOwner"
	BULK_ADD_TAG          api.APIOperation = "customization_bulkAddTag"
	BULK_UPDATE_TAG       api.APIOperation = "customization_bulkUpdateTag"
	BULK_DELETE_TAG       api.APIOperation = "customization_bulkDeleteTag"
	BULK_CHANGE_TAG_OWNER api.APIOperation = "customization_bulkChangeTagOwner"

	// READ APIs.
	LIST_TAGS     api.APIOperation = "customization_listTags"
	COUNT_TAGS    api.APIOperation = "customization_countTags"
	GET_TAG_BY_ID api.APIOperation = "ccustomization_getTagById"
)
