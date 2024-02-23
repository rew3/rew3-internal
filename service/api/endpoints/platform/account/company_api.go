package account

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Company
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_COMPANY    endpoints.Endpoint = "account_addCompany"
	UPDATE_COMPANY endpoints.Endpoint = "account_updateCompany"
	DELETE_COMPANY endpoints.Endpoint = "account_deleteCompany"

	// READ APIs.
	LIST_COMPANIES    endpoints.Endpoint = "account_listCompanies"
	COUNT_COMPANIES   endpoints.Endpoint = "account_countCompanies"
	GET_COMPANY_BY_ID endpoints.Endpoint = "caccount_getCompanyById"
)
