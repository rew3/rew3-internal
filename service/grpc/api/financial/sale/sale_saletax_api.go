package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for SaleTax
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_SALETAX               api.APIOperation = "financial_addSaleTax"
	UPDATE_SALETAX            api.APIOperation = "financial_updateSaleTax"
	DELETE_SALETAX            api.APIOperation = "financial_deleteSaleTax"
	CLONE_SALETAX             api.APIOperation = "financial_cloneSaleTax"
	CHANGE_SALETAX_OWNER      api.APIOperation = "financial_changeSaleTaxOwner"
	BULK_ADD_SALETAX          api.APIOperation = "financial_bulkAddSaleTax"
	BULK_UPDATE_SALETAX       api.APIOperation = "financial_bulkUpdateSaleTax"
	BULK_DELETE_SALETAX       api.APIOperation = "financial_bulkDeleteSaleTax"
	BULK_CHANGE_SALETAX_OWNER api.APIOperation = "financial_bulkChangeSaleTaxOwner"

	// READ APIs.
	LIST_SALETAXES    api.APIOperation = "financial_listSaleTaxes"
	COUNT_SALETAXES   api.APIOperation = "financial_countSaleTaxes"
	GET_SALETAX_BY_ID api.APIOperation = "cfinancial_getSaleTaxById"
)
