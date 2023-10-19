package account

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Company
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_COMPANY    api.APIOperation = "account_addCompany"
	UPDATE_COMPANY api.APIOperation = "account_updateCompany"
	DELETE_COMPANY api.APIOperation = "account_deleteCompany"

	// READ APIs.
	LIST_COMPANIES    api.APIOperation = "account_listCompanies"
	COUNT_COMPANIES   api.APIOperation = "account_countCompanies"
	GET_COMPANY_BY_ID api.APIOperation = "caccount_getCompanyById"
)
