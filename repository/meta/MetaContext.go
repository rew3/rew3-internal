package meta

import (
	tpe "github.com/rew3/rew3-base/data/entity"
	consts "github.com/rew3/rew3-base/data/service/constants"
)

type MetaContext struct {
	ContextUser tpe.MiniUser
	MemberId    string
	Entity      string
	Module      string
	AccountType consts.AccountTypeAlias
}
