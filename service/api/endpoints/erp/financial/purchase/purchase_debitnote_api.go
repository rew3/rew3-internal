package purchase

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for DebitNote
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DEBITNOTE               endpoints.Endpoint = "financial_addDebitNote"
	UPDATE_DEBITNOTE            endpoints.Endpoint = "financial_updateDebitNote"
	DELETE_DEBITNOTE            endpoints.Endpoint = "financial_deleteDebitNote"
	CLONE_DEBITNOTE             endpoints.Endpoint = "financial_cloneDebitNote"
	CHANGE_DEBITNOTE_OWNER      endpoints.Endpoint = "financial_changeDebitNoteOwner"
	BULK_ADD_DEBITNOTE          endpoints.Endpoint = "financial_bulkAddDebitNote"
	BULK_UPDATE_DEBITNOTE       endpoints.Endpoint = "financial_bulkUpdateDebitNote"
	BULK_DELETE_DEBITNOTE       endpoints.Endpoint = "financial_bulkDeleteDebitNote"
	BULK_CHANGE_DEBITNOTE_OWNER endpoints.Endpoint = "financial_bulkChangeDebitNoteOwner"

	// READ APIs.
	LIST_DEBITNOTES     endpoints.Endpoint = "financial_listDebitNotes"
	COUNT_DEBITNOTES    endpoints.Endpoint = "financial_countDebitNotes"
	GET_DEBITNOTE_BY_ID endpoints.Endpoint = "cfinancial_getDebitNoteById"
)
