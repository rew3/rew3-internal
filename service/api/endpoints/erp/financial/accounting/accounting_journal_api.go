package accounting

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Journal
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_JOURNAL               endpoints.Endpoint = "financial_addJournal"
	UPDATE_JOURNAL            endpoints.Endpoint = "financial_updateJournal"
	DELETE_JOURNAL            endpoints.Endpoint = "financial_deleteJournal"
	CLONE_JOURNAL             endpoints.Endpoint = "financial_cloneJournal"
	CHANGE_JOURNAL_OWNER      endpoints.Endpoint = "financial_changeJournalOwner"
	BULK_ADD_JOURNAL          endpoints.Endpoint = "financial_bulkAddJournal"
	BULK_UPDATE_JOURNAL       endpoints.Endpoint = "financial_bulkUpdateJournal"
	BULK_DELETE_JOURNAL       endpoints.Endpoint = "financial_bulkDeleteJournal"
	BULK_CHANGE_JOURNAL_OWNER endpoints.Endpoint = "financial_bulkChangeJournalOwner"

	// READ APIs.
	LIST_JOURNALS     endpoints.Endpoint = "financial_listJournals"
	COUNT_JOURNALS    endpoints.Endpoint = "financial_countJournals"
	GET_JOURNAL_BY_ID endpoints.Endpoint = "cfinancial_getJournalById"
)
