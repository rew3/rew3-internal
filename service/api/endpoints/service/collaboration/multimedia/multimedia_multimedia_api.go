package multimedia

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Multimedia
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_MULTIMEDIA               endpoints.Endpoint = "collaboration_addMultimedia"
	UPDATE_MULTIMEDIA            endpoints.Endpoint = "collaboration_updateMultimedia"
	DELETE_MULTIMEDIA            endpoints.Endpoint = "collaboration_deleteMultimedia"
	CLONE_MULTIMEDIA             endpoints.Endpoint = "collaboration_cloneMultimedia"
	CHANGE_MULTIMEDIA_OWNER      endpoints.Endpoint = "collaboration_changeMultimediaOwner"
	BULK_ADD_MULTIMEDIA          endpoints.Endpoint = "collaboration_bulkAddMultimedia"
	BULK_UPDATE_MULTIMEDIA       endpoints.Endpoint = "collaboration_bulkUpdateMultimedia"
	BULK_DELETE_MULTIMEDIA       endpoints.Endpoint = "collaboration_bulkDeleteMultimedia"
	BULK_CHANGE_MULTIMEDIA_OWNER endpoints.Endpoint = "collaboration_bulkChangeMultimediaOwner"

	// READ APIs.
	LIST_MULTIMEDIA      endpoints.Endpoint = "collaboration_listMultimedia"
	COUNT_MULTIMEDIA     endpoints.Endpoint = "collaboration_countMultimedia"
	GET_MULTIMEDIA_BY_ID endpoints.Endpoint = "ccollaboration_getMultimediaById"
)
