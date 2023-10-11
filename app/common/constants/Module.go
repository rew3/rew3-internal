package constants

type Module string

const (
	// ERP
	CRM       Module = "crm"
	RMS       Module = "rms"
	CMS       Module = "cms"
	DMS       Module = "dms"
	FINANCIAL Module = "financial"

	// REW3 and Platform.
	PROJECT       Module = "project"
	COLLABORATION Module = "collaboration"
	ACTIVITY      Module = "activity"
	LISTVIEW      Module = "listview"
	TAGS          Module = "tags"

	ACCOUNT Module = "accountService"
	UTILITY Module = "utility"
)
