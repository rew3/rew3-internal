package comment

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Comment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_COMMENT               endpoints.Endpoint = "collaboration_addComment"
	UPDATE_COMMENT            endpoints.Endpoint = "collaboration_updateComment"
	DELETE_COMMENT            endpoints.Endpoint = "collaboration_deleteComment"
	CLONE_COMMENT             endpoints.Endpoint = "collaboration_cloneComment"
	CHANGE_COMMENT_OWNER      endpoints.Endpoint = "collaboration_changeCommentOwner"
	BULK_ADD_COMMENT          endpoints.Endpoint = "collaboration_bulkAddComment"
	BULK_UPDATE_COMMENT       endpoints.Endpoint = "collaboration_bulkUpdateComment"
	BULK_DELETE_COMMENT       endpoints.Endpoint = "collaboration_bulkDeleteComment"
	BULK_CHANGE_COMMENT_OWNER endpoints.Endpoint = "collaboration_bulkChangeCommentOwner"

	// READ APIs.
	LIST_COMMENTS     endpoints.Endpoint = "collaboration_listComments"
	COUNT_COMMENTS    endpoints.Endpoint = "collaboration_countComments"
	GET_COMMENT_BY_ID endpoints.Endpoint = "ccollaboration_getCommentById"
)
