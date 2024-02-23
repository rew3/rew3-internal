package purchase

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Expense
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EXPENSE               endpoints.Endpoint = "financial_addExpense"
	UPDATE_EXPENSE            endpoints.Endpoint = "financial_updateExpense"
	DELETE_EXPENSE            endpoints.Endpoint = "financial_deleteExpense"
	CLONE_EXPENSE             endpoints.Endpoint = "financial_cloneExpense"
	CHANGE_EXPENSE_OWNER      endpoints.Endpoint = "financial_changeExpenseOwner"
	BULK_ADD_EXPENSE          endpoints.Endpoint = "financial_bulkAddExpense"
	BULK_UPDATE_EXPENSE       endpoints.Endpoint = "financial_bulkUpdateExpense"
	BULK_DELETE_EXPENSE       endpoints.Endpoint = "financial_bulkDeleteExpense"
	BULK_CHANGE_EXPENSE_OWNER endpoints.Endpoint = "financial_bulkChangeExpenseOwner"

	// READ APIs.
	LIST_EXPENSES     endpoints.Endpoint = "financial_listExpenses"
	COUNT_EXPENSES    endpoints.Endpoint = "financial_countExpenses"
	GET_EXPENSE_BY_ID endpoints.Endpoint = "cfinancial_getExpenseById"
)
