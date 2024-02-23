package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for SaleTax
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_SALETAX               endpoints.Endpoint = "financial_addSaleTax"
	UPDATE_SALETAX            endpoints.Endpoint = "financial_updateSaleTax"
	DELETE_SALETAX            endpoints.Endpoint = "financial_deleteSaleTax"
	CLONE_SALETAX             endpoints.Endpoint = "financial_cloneSaleTax"
	CHANGE_SALETAX_OWNER      endpoints.Endpoint = "financial_changeSaleTaxOwner"
	BULK_ADD_SALETAX          endpoints.Endpoint = "financial_bulkAddSaleTax"
	BULK_UPDATE_SALETAX       endpoints.Endpoint = "financial_bulkUpdateSaleTax"
	BULK_DELETE_SALETAX       endpoints.Endpoint = "financial_bulkDeleteSaleTax"
	BULK_CHANGE_SALETAX_OWNER endpoints.Endpoint = "financial_bulkChangeSaleTaxOwner"

	// READ APIs.
	LIST_SALETAXES    endpoints.Endpoint = "financial_listSaleTaxes"
	COUNT_SALETAXES   endpoints.Endpoint = "financial_countSaleTaxes"
	GET_SALETAX_BY_ID endpoints.Endpoint = "cfinancial_getSaleTaxById"
)
