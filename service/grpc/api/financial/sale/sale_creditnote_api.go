package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for CreditNote
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_CREDITNOTE                   api.APIOperation = "financial_addCreditNote"
	UPDATE_CREDITNOTE                 api.APIOperation = "financial_updateCreditNote"
	DELETE_CREDITNOTE                 api.APIOperation = "financial_deleteCreditNote"
	CLONE_CREDITNOTE                  api.APIOperation = "financial_cloneCreditNote"
	CHANGE_CREDITNOTE_OWNER           api.APIOperation = "financial_changeCreditNoteOwner"
	BULK_ADD_CREDITNOTE          api.APIOperation = "financial_bulkAddCreditNote"
	BULK_UPDATE_CREDITNOTE           api.APIOperation = "financial_bulkUpdateCreditNote"
	BULK_DELETE_CREDITNOTE           api.APIOperation = "financial_bulkDeleteCreditNote"
	BULK_CHANGE_CREDITNOTE_OWNER           api.APIOperation = "financial_bulkChangeCreditNoteOwner"

	// READ APIs.
	LIST_CREDITNOTES     api.APIOperation = "financial_listCreditNotes"
	COUNT_CREDITNOTES    api.APIOperation = "financial_countCreditNotes"
	GET_CREDITNOTE_BY_ID api.APIOperation = "cfinancial_getCreditNoteById"
    
)