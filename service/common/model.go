package common

import (
	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/common"
)

type Model interface{}

type ModelWrapper interface {
	GetId() string
	SetId(id string)
	GetMeta() common.MetaInfo
	SetMeta(meta common.MetaInfo)
	GetOwner() account.MiniUser
	SetOwner(owner account.MiniUser)
	GetVisibility() account.RecordVisibility
	SetVisibility(visibility account.RecordVisibility)
	GetReference() []common.Reference
	SetReference(reference []common.Reference)
	IsFavorite() bool
	SetFavorite(status bool)
}
