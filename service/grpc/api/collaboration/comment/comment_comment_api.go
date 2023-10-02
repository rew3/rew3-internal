package comment

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Comment
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_COMMENT                   api.APIOperation = "collaboration_addComment"
	UPDATE_COMMENT                 api.APIOperation = "collaboration_updateComment"
	DELETE_COMMENT                 api.APIOperation = "collaboration_deleteComment"
	CLONE_COMMENT                  api.APIOperation = "collaboration_cloneComment"
	CHANGE_COMMENT_OWNER           api.APIOperation = "collaboration_changeCommentOwner"
	BULK_ADD_COMMENT          api.APIOperation = "collaboration_bulkAddComment"
	BULK_UPDATE_COMMENT           api.APIOperation = "collaboration_bulkUpdateComment"
	BULK_DELETE_COMMENT           api.APIOperation = "collaboration_bulkDeleteComment"
	BULK_CHANGE_COMMENT_OWNER           api.APIOperation = "collaboration_bulkChangeCommentOwner"

	// READ APIs.
	LIST_COMMENTS     api.APIOperation = "collaboration_listComments"
	COUNT_COMMENTS    api.APIOperation = "collaboration_countComments"
	GET_COMMENT_BY_ID api.APIOperation = "ccollaboration_getCommentById"
    
)