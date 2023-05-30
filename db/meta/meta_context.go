package meta

import (
	tpe "github.com/rew3/rew3-base/app/model"
	consts "github.com/rew3/rew3-base/app/service/constants"
)

type MetaContext struct {
	ContextUser tpe.MiniUser
	MemberId    string
	Entity      string
	Module      string
	AccountType consts.AccountTypeAlias
}
