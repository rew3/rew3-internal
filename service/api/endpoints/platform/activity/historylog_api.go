package activity

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for HistoryLog
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_HISTORYLOG         endpoints.Endpoint = "activity_addHistoryLog"
	UPDATE_HISTORYLOG      endpoints.Endpoint = "activity_updateHistoryLog"
	DELETE_HISTORYLOG      endpoints.Endpoint = "activity_deleteHistoryLog"
	BULK_ADD_HISTORYLOG    endpoints.Endpoint = "activity_bulkAddHistoryLog"
	BULK_UPDATE_HISTORYLOG endpoints.Endpoint = "activity_bulkUpdateHistoryLog"
	BULK_DELETE_HISTORYLOG endpoints.Endpoint = "activity_bulkDeleteHistoryLog"

	// READ APIs.
	LIST_HISTORYLOG      endpoints.Endpoint = "activity_listHistoryLog"
	COUNT_HISTORYLOG     endpoints.Endpoint = "activity_countHistoryLog"
	GET_HISTORYLOG_BY_ID endpoints.Endpoint = "cactivity_getHistoryLogById"
)
