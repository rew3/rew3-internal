package meta

import (
	tpe "github.com/rew3/rew3-internal/app/account"
	consts "github.com/rew3/rew3-internal/app/account/constants"
)

type MetaContext struct {
	ContextUser tpe.MiniUser
	MemberId    string
	Entity      string
	Module      string
	AccountType consts.AccountTypeAlias
}
