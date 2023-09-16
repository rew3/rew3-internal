package constants

type Entity string

const (
	CRM_CONTACT              Entity = "crm_contact"
	CRM_ACCOUNT              Entity = "crm_account"
	CRM_LEAD                 Entity = "crm_lead"
	CRM_OPPORTUNITY          Entity = "crm_opportunity"
	CRM_OPPORTUNITY_PIPELINE Entity = "crm_opportunity_pipeline"
	CRM_CASE                 Entity = "crm_case"
	CRM_CAMPAIGN             Entity = "crm_campaign"

	//RMS Entities
	RMS_LISTING Entity = "rms_listing"
	RMS_REQUEST Entity = "rms_request"

	// Add more later.
)
