package common

import (
	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/common"
)

type Model interface{}

type ModelWrapper interface {
	GetId() string
	SetId(id string) Model
	GetMeta() common.MetaInfo
	SetMeta(meta common.MetaInfo) Model
	GetOwner() account.MiniUser
	SetOwner(owner account.MiniUser) Model
	GetVisibility() account.RecordVisibility
	SetVisibility(visibility account.RecordVisibility) Model
	GetReference() []common.Reference
	SetReference(reference []common.Reference) Model
	IsFavorite() bool
	SetFavorite(status bool) Model
}
