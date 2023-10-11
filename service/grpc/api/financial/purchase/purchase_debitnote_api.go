package purchase

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for DebitNote
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DEBITNOTE               api.APIOperation = "financial_addDebitNote"
	UPDATE_DEBITNOTE            api.APIOperation = "financial_updateDebitNote"
	DELETE_DEBITNOTE            api.APIOperation = "financial_deleteDebitNote"
	CLONE_DEBITNOTE             api.APIOperation = "financial_cloneDebitNote"
	CHANGE_DEBITNOTE_OWNER      api.APIOperation = "financial_changeDebitNoteOwner"
	BULK_ADD_DEBITNOTE          api.APIOperation = "financial_bulkAddDebitNote"
	BULK_UPDATE_DEBITNOTE       api.APIOperation = "financial_bulkUpdateDebitNote"
	BULK_DELETE_DEBITNOTE       api.APIOperation = "financial_bulkDeleteDebitNote"
	BULK_CHANGE_DEBITNOTE_OWNER api.APIOperation = "financial_bulkChangeDebitNoteOwner"

	// READ APIs.
	LIST_DEBITNOTES     api.APIOperation = "financial_listDebitNotes"
	COUNT_DEBITNOTES    api.APIOperation = "financial_countDebitNotes"
	GET_DEBITNOTE_BY_ID api.APIOperation = "cfinancial_getDebitNoteById"
)
