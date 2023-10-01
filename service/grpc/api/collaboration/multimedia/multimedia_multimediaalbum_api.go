package multimedia

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for MultimediaAlbum
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_MULTIMEDIAALBUM               api.APIOperation = "collaboration_addMultimediaAlbum"
	UPDATE_MULTIMEDIAALBUM            api.APIOperation = "collaboration_updateMultimediaAlbum"
	DELETE_MULTIMEDIAALBUM            api.APIOperation = "collaboration_deleteMultimediaAlbum"
	CLONE_MULTIMEDIAALBUM             api.APIOperation = "collaboration_cloneMultimediaAlbum"
	CHANGE_MULTIMEDIAALBUM_OWNER      api.APIOperation = "collaboration_changeMultimediaAlbumOwner"
	BULK_ADD_MULTIMEDIAALBUM          api.APIOperation = "collaboration_bulkAddMultimediaAlbum"
	BULK_UPDATE_MULTIMEDIAALBUM       api.APIOperation = "collaboration_bulkUpdateMultimediaAlbum"
	BULK_DELETE_MULTIMEDIAALBUM       api.APIOperation = "collaboration_bulkDeleteMultimediaAlbum"
	BULK_CHANGE_MULTIMEDIAALBUM_OWNER api.APIOperation = "collaboration_bulkChangeMultimediaAlbumOwner"

	// READ APIs.
	LIST_MULTIMEDIAALBUMS     api.APIOperation = "collaboration_listMultimediaAlbums"
	COUNT_MULTIMEDIAALBUMS    api.APIOperation = "collaboration_countMultimediaAlbums"
	GET_MULTIMEDIAALBUM_BY_ID api.APIOperation = "ccollaboration_getMultimediaAlbumById"
)
