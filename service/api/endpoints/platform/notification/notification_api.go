package notification

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Notification
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_NOTIFICATION               endpoints.Endpoint = "notification_addNotification"
	UPDATE_NOTIFICATION            endpoints.Endpoint = "notification_updateNotification"
	DELETE_NOTIFICATION            endpoints.Endpoint = "notification_deleteNotification"
	CLONE_NOTIFICATION             endpoints.Endpoint = "notification_cloneNotification"
	CHANGE_NOTIFICATION_OWNER      endpoints.Endpoint = "notification_changeNotificationOwner"
	BULK_ADD_NOTIFICATION          endpoints.Endpoint = "notification_bulkAddNotification"
	BULK_UPDATE_NOTIFICATION       endpoints.Endpoint = "notification_bulkUpdateNotification"
	BULK_DELETE_NOTIFICATION       endpoints.Endpoint = "notification_bulkDeleteNotification"
	BULK_CHANGE_NOTIFICATION_OWNER endpoints.Endpoint = "notification_bulkChangeNotificationOwner"

	// READ APIs.
	LIST_NOTIFICATIONS     endpoints.Endpoint = "notification_listNotifications"
	COUNT_NOTIFICATIONS    endpoints.Endpoint = "notification_countNotifications"
	GET_NOTIFICATION_BY_ID endpoints.Endpoint = "cnotification_getNotificationById"
)
