package meta

import (
	tpe "github.com/rew3/rew3-pkg/common/account"
	consts "github.com/rew3/rew3-pkg/common/account/constants"
)

type MetaContext struct {
	ContextUser tpe.MiniUser
	MemberId    string
	Entity      string
	Module      string
	AccountType consts.AccountTypeAlias
}
