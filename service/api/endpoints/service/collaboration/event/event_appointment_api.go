package event

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Appointment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_APPOINTMENT               endpoints.Endpoint = "collaboration_addAppointment"
	UPDATE_APPOINTMENT            endpoints.Endpoint = "collaboration_updateAppointment"
	DELETE_APPOINTMENT            endpoints.Endpoint = "collaboration_deleteAppointment"
	CLONE_APPOINTMENT             endpoints.Endpoint = "collaboration_cloneAppointment"
	CHANGE_APPOINTMENT_OWNER      endpoints.Endpoint = "collaboration_changeAppointmentOwner"
	BULK_ADD_APPOINTMENT          endpoints.Endpoint = "collaboration_bulkAddAppointment"
	BULK_UPDATE_APPOINTMENT       endpoints.Endpoint = "collaboration_bulkUpdateAppointment"
	BULK_DELETE_APPOINTMENT       endpoints.Endpoint = "collaboration_bulkDeleteAppointment"
	BULK_CHANGE_APPOINTMENT_OWNER endpoints.Endpoint = "collaboration_bulkChangeAppointmentOwner"

	// READ APIs.
	LIST_APPOINTMENTS     endpoints.Endpoint = "collaboration_listAppointments"
	COUNT_APPOINTMENTS    endpoints.Endpoint = "collaboration_countAppointments"
	GET_APPOINTMENT_BY_ID endpoints.Endpoint = "ccollaboration_getAppointmentById"
)
