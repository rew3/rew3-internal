package enums

import (
	"reflect"

	"github.com/rew3/rew3-internal/app/common/constants"
	"github.com/rew3/rew3-internal/gen"
)

func GetCoreEnumTypeMapping() []gen.EnumType {
	return []gen.EnumType{
		{Type: reflect.TypeOf(constants.Entity("")), Items: []string{"crm_account", "crm_contact", "crm_lead", "crm_case", "crm_opportunity"}},
		{Type: reflect.TypeOf(constants.Module("")), Items: []string{"crm", "rms", "cms", "dms", "financial", "project"}},
		// TODO add more later.
	}
}
