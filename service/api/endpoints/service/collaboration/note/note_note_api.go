package note

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Note
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_NOTE               endpoints.Endpoint = "collaboration_addNote"
	UPDATE_NOTE            endpoints.Endpoint = "collaboration_updateNote"
	DELETE_NOTE            endpoints.Endpoint = "collaboration_deleteNote"
	CLONE_NOTE             endpoints.Endpoint = "collaboration_cloneNote"
	CHANGE_NOTE_OWNER      endpoints.Endpoint = "collaboration_changeNoteOwner"
	BULK_ADD_NOTE          endpoints.Endpoint = "collaboration_bulkAddNote"
	BULK_UPDATE_NOTE       endpoints.Endpoint = "collaboration_bulkUpdateNote"
	BULK_DELETE_NOTE       endpoints.Endpoint = "collaboration_bulkDeleteNote"
	BULK_CHANGE_NOTE_OWNER endpoints.Endpoint = "collaboration_bulkChangeNoteOwner"

	// READ APIs.
	LIST_NOTES     endpoints.Endpoint = "collaboration_listNotes"
	COUNT_NOTES    endpoints.Endpoint = "collaboration_countNotes"
	GET_NOTE_BY_ID endpoints.Endpoint = "ccollaboration_getNoteById"
)
