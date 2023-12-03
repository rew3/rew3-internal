package constants

type Module string

const (
	// ERP
	CRM       Module = "crm"
	RMS       Module = "rms"
	CMS       Module = "cms"
	DMS       Module = "dms"
	FINANCIAL Module = "financial"
	PROJECT   Module = "project"

	// REW3 and Platform.
	COLLABORATION Module = "collaboration"
	ACTIVITY      Module = "activity"
	CUSTOMIZATION Module = "listview"

	ACCOUNT Module = "accountService"
	UTILITY Module = "utility"

	CHAT Module = "chat"
)
