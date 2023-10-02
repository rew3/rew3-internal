package event

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Appointment
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_APPOINTMENT                   api.APIOperation = "collaboration_addAppointment"
	UPDATE_APPOINTMENT                 api.APIOperation = "collaboration_updateAppointment"
	DELETE_APPOINTMENT                 api.APIOperation = "collaboration_deleteAppointment"
	CLONE_APPOINTMENT                  api.APIOperation = "collaboration_cloneAppointment"
	CHANGE_APPOINTMENT_OWNER           api.APIOperation = "collaboration_changeAppointmentOwner"
	BULK_ADD_APPOINTMENT          api.APIOperation = "collaboration_bulkAddAppointment"
	BULK_UPDATE_APPOINTMENT           api.APIOperation = "collaboration_bulkUpdateAppointment"
	BULK_DELETE_APPOINTMENT           api.APIOperation = "collaboration_bulkDeleteAppointment"
	BULK_CHANGE_APPOINTMENT_OWNER           api.APIOperation = "collaboration_bulkChangeAppointmentOwner"

	// READ APIs.
	LIST_APPOINTMENTS     api.APIOperation = "collaboration_listAppointments"
	COUNT_APPOINTMENTS    api.APIOperation = "collaboration_countAppointments"
	GET_APPOINTMENT_BY_ID api.APIOperation = "ccollaboration_getAppointmentById"
    
)