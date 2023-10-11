package purchase

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Expense
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EXPENSE               api.APIOperation = "financial_addExpense"
	UPDATE_EXPENSE            api.APIOperation = "financial_updateExpense"
	DELETE_EXPENSE            api.APIOperation = "financial_deleteExpense"
	CLONE_EXPENSE             api.APIOperation = "financial_cloneExpense"
	CHANGE_EXPENSE_OWNER      api.APIOperation = "financial_changeExpenseOwner"
	BULK_ADD_EXPENSE          api.APIOperation = "financial_bulkAddExpense"
	BULK_UPDATE_EXPENSE       api.APIOperation = "financial_bulkUpdateExpense"
	BULK_DELETE_EXPENSE       api.APIOperation = "financial_bulkDeleteExpense"
	BULK_CHANGE_EXPENSE_OWNER api.APIOperation = "financial_bulkChangeExpenseOwner"

	// READ APIs.
	LIST_EXPENSES     api.APIOperation = "financial_listExpenses"
	COUNT_EXPENSES    api.APIOperation = "financial_countExpenses"
	GET_EXPENSE_BY_ID api.APIOperation = "cfinancial_getExpenseById"
)
