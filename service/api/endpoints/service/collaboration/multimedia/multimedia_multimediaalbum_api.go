package multimedia

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for MultimediaAlbum
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_MULTIMEDIAALBUM               endpoints.Endpoint = "collaboration_addMultimediaAlbum"
	UPDATE_MULTIMEDIAALBUM            endpoints.Endpoint = "collaboration_updateMultimediaAlbum"
	DELETE_MULTIMEDIAALBUM            endpoints.Endpoint = "collaboration_deleteMultimediaAlbum"
	CLONE_MULTIMEDIAALBUM             endpoints.Endpoint = "collaboration_cloneMultimediaAlbum"
	CHANGE_MULTIMEDIAALBUM_OWNER      endpoints.Endpoint = "collaboration_changeMultimediaAlbumOwner"
	BULK_ADD_MULTIMEDIAALBUM          endpoints.Endpoint = "collaboration_bulkAddMultimediaAlbum"
	BULK_UPDATE_MULTIMEDIAALBUM       endpoints.Endpoint = "collaboration_bulkUpdateMultimediaAlbum"
	BULK_DELETE_MULTIMEDIAALBUM       endpoints.Endpoint = "collaboration_bulkDeleteMultimediaAlbum"
	BULK_CHANGE_MULTIMEDIAALBUM_OWNER endpoints.Endpoint = "collaboration_bulkChangeMultimediaAlbumOwner"

	// READ APIs.
	LIST_MULTIMEDIAALBUMS     endpoints.Endpoint = "collaboration_listMultimediaAlbums"
	COUNT_MULTIMEDIAALBUMS    endpoints.Endpoint = "collaboration_countMultimediaAlbums"
	GET_MULTIMEDIAALBUM_BY_ID endpoints.Endpoint = "ccollaboration_getMultimediaAlbumById"
)
