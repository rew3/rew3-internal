package enums

import (
	"reflect"

	ac "github.com/rew3/rew3-internal/app/account/constants"
	"github.com/rew3/rew3-internal/app/common/constants"
	"github.com/rew3/rew3-internal/gen"
)

func GetCoreEnumTypeMapping() []gen.EnumType {
	return []gen.EnumType{
		{Type: reflect.TypeOf(ac.AccountTypeAlias("")), Items: []string{"TEAM", "ORGANIZATION"}},
		{Type: reflect.TypeOf(ac.PermissionType("")), Items: []string{"privateOpen", "publicClose", "privateClose", "publicOpen"}},
		{Type: reflect.TypeOf(ac.Rew3AccountType("")), Items: []string{"TEAM", "ORGANIZATION"}},
		{Type: reflect.TypeOf(ac.SecurityTypeAlias("")), Items: []string{"SIMPLE", "ADVANCE"}},
		{Type: reflect.TypeOf(ac.UserType("")), Items: []string{"INTERNAL", "EXTERNAL"}},
		{Type: reflect.TypeOf(ac.VisibilityType("")), Items: []string{"PRIVATE", "EVERYONE", "INDIVIDUAL", "TEAMS"}},

		{Type: reflect.TypeOf(constants.AccessType("")), Items: []string{"privateOpen", "publicClose", "privateClose", "publicOpen"}},
		{Type: reflect.TypeOf(constants.Currency("")), Items: []string{"USD", "CAD"}},
		{Type: reflect.TypeOf(constants.Entity("")), Items: []string{"crm_account", "crm_contact", "crm_lead", "crm_case", "crm_opportunity"}},
		{Type: reflect.TypeOf(constants.Module("")), Items: []string{"crm", "rms", "cms", "dms", "financial", "project"}},
		{Type: reflect.TypeOf(constants.TimeZone("")), Items: []string{"Canada/Eastern", "Canada/Central"}},
		// TODO add more later.
	}
}
