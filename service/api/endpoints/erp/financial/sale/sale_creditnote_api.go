package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for CreditNote
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CREDITNOTE               endpoints.Endpoint = "financial_addCreditNote"
	UPDATE_CREDITNOTE            endpoints.Endpoint = "financial_updateCreditNote"
	DELETE_CREDITNOTE            endpoints.Endpoint = "financial_deleteCreditNote"
	CLONE_CREDITNOTE             endpoints.Endpoint = "financial_cloneCreditNote"
	CHANGE_CREDITNOTE_OWNER      endpoints.Endpoint = "financial_changeCreditNoteOwner"
	BULK_ADD_CREDITNOTE          endpoints.Endpoint = "financial_bulkAddCreditNote"
	BULK_UPDATE_CREDITNOTE       endpoints.Endpoint = "financial_bulkUpdateCreditNote"
	BULK_DELETE_CREDITNOTE       endpoints.Endpoint = "financial_bulkDeleteCreditNote"
	BULK_CHANGE_CREDITNOTE_OWNER endpoints.Endpoint = "financial_bulkChangeCreditNoteOwner"

	// READ APIs.
	LIST_CREDITNOTES     endpoints.Endpoint = "financial_listCreditNotes"
	COUNT_CREDITNOTES    endpoints.Endpoint = "financial_countCreditNotes"
	GET_CREDITNOTE_BY_ID endpoints.Endpoint = "cfinancial_getCreditNoteById"
)
