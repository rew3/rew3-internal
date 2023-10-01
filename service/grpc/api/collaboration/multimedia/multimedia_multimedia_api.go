package multimedia

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Multimedia
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_MULTIMEDIA               api.APIOperation = "collaboration_addMultimedia"
	UPDATE_MULTIMEDIA            api.APIOperation = "collaboration_updateMultimedia"
	DELETE_MULTIMEDIA            api.APIOperation = "collaboration_deleteMultimedia"
	CLONE_MULTIMEDIA             api.APIOperation = "collaboration_cloneMultimedia"
	CHANGE_MULTIMEDIA_OWNER      api.APIOperation = "collaboration_changeMultimediaOwner"
	BULK_ADD_MULTIMEDIA          api.APIOperation = "collaboration_bulkAddMultimedia"
	BULK_UPDATE_MULTIMEDIA       api.APIOperation = "collaboration_bulkUpdateMultimedia"
	BULK_DELETE_MULTIMEDIA       api.APIOperation = "collaboration_bulkDeleteMultimedia"
	BULK_CHANGE_MULTIMEDIA_OWNER api.APIOperation = "collaboration_bulkChangeMultimediaOwner"

	// READ APIs.
	LIST_MULTIMEDIA      api.APIOperation = "collaboration_listMultimedia"
	COUNT_MULTIMEDIA     api.APIOperation = "collaboration_countMultimedia"
	GET_MULTIMEDIA_BY_ID api.APIOperation = "ccollaboration_getMultimediaById"
)
