package note

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Note
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_NOTE               api.APIOperation = "collaboration_addNote"
	UPDATE_NOTE            api.APIOperation = "collaboration_updateNote"
	DELETE_NOTE            api.APIOperation = "collaboration_deleteNote"
	CLONE_NOTE             api.APIOperation = "collaboration_cloneNote"
	CHANGE_NOTE_OWNER      api.APIOperation = "collaboration_changeNoteOwner"
	BULK_ADD_NOTE          api.APIOperation = "collaboration_bulkAddNote"
	BULK_UPDATE_NOTE       api.APIOperation = "collaboration_bulkUpdateNote"
	BULK_DELETE_NOTE       api.APIOperation = "collaboration_bulkDeleteNote"
	BULK_CHANGE_NOTE_OWNER api.APIOperation = "collaboration_bulkChangeNoteOwner"

	// READ APIs.
	LIST_NOTES     api.APIOperation = "collaboration_listNotes"
	COUNT_NOTES    api.APIOperation = "collaboration_countNotes"
	GET_NOTE_BY_ID api.APIOperation = "ccollaboration_getNoteById"
)
