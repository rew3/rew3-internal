package meta

import (
	consts "github.com/rew3/rew3-base/service/constants"
	tpe "github.com/rew3/rew3-base/types"
)

type MetaContext struct {
	ContextUser tpe.MiniUser
	MemberId    string
	Entity      string
	Module      string
	AccountType consts.AccountTypeAlias
}
