package meta

import (
	tpe "github.com/rew3/rew3-base/common/types"
	consts "github.com/rew3/rew3-base/service/constants"
)

type MetaContext struct {
	ContextUser tpe.MiniUser
	MemberId    string
	Entity      string
	Module      string
	AccountType consts.AccountTypeAlias
}
