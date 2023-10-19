package notification

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Notification
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_NOTIFICATION               api.APIOperation = "notification_addNotification"
	UPDATE_NOTIFICATION            api.APIOperation = "notification_updateNotification"
	DELETE_NOTIFICATION            api.APIOperation = "notification_deleteNotification"
	CLONE_NOTIFICATION             api.APIOperation = "notification_cloneNotification"
	CHANGE_NOTIFICATION_OWNER      api.APIOperation = "notification_changeNotificationOwner"
	BULK_ADD_NOTIFICATION          api.APIOperation = "notification_bulkAddNotification"
	BULK_UPDATE_NOTIFICATION       api.APIOperation = "notification_bulkUpdateNotification"
	BULK_DELETE_NOTIFICATION       api.APIOperation = "notification_bulkDeleteNotification"
	BULK_CHANGE_NOTIFICATION_OWNER api.APIOperation = "notification_bulkChangeNotificationOwner"

	// READ APIs.
	LIST_NOTIFICATIONS     api.APIOperation = "notification_listNotifications"
	COUNT_NOTIFICATIONS    api.APIOperation = "notification_countNotifications"
	GET_NOTIFICATION_BY_ID api.APIOperation = "cnotification_getNotificationById"
)
