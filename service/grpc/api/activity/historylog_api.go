package activity

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for HistoryLog
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_HISTORYLOG         api.APIOperation = "activity_addHistoryLog"
	UPDATE_HISTORYLOG      api.APIOperation = "activity_updateHistoryLog"
	DELETE_HISTORYLOG      api.APIOperation = "activity_deleteHistoryLog"
	BULK_ADD_HISTORYLOG    api.APIOperation = "activity_bulkAddHistoryLog"
	BULK_UPDATE_HISTORYLOG api.APIOperation = "activity_bulkUpdateHistoryLog"
	BULK_DELETE_HISTORYLOG api.APIOperation = "activity_bulkDeleteHistoryLog"

	// READ APIs.
	LIST_HISTORYLOG      api.APIOperation = "activity_listHistoryLog"
	COUNT_HISTORYLOG     api.APIOperation = "activity_countHistoryLog"
	GET_HISTORYLOG_BY_ID api.APIOperation = "cactivity_getHistoryLogById"
)
