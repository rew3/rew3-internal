package accounting

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Journal
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_JOURNAL                   api.APIOperation = "financial_addJournal"
	UPDATE_JOURNAL                 api.APIOperation = "financial_updateJournal"
	DELETE_JOURNAL                 api.APIOperation = "financial_deleteJournal"
	CLONE_JOURNAL                  api.APIOperation = "financial_cloneJournal"
	CHANGE_JOURNAL_OWNER           api.APIOperation = "financial_changeJournalOwner"
	BULK_ADD_JOURNAL          api.APIOperation = "financial_bulkAddJournal"
	BULK_UPDATE_JOURNAL           api.APIOperation = "financial_bulkUpdateJournal"
	BULK_DELETE_JOURNAL           api.APIOperation = "financial_bulkDeleteJournal"
	BULK_CHANGE_JOURNAL_OWNER           api.APIOperation = "financial_bulkChangeJournalOwner"

	// READ APIs.
	LIST_JOURNALS     api.APIOperation = "financial_listJournals"
	COUNT_JOURNALS    api.APIOperation = "financial_countJournals"
	GET_JOURNAL_BY_ID api.APIOperation = "cfinancial_getJournalById"
    
)