package entity

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntityInterface interface {
	GetId() int
	SetId(id int)
	GetCreated() time.Time
	SetCreated(created time.Time)
	GetModified() time.Time
	SetModified(modified time.Time)
}

type BaseEntity struct {
	BaseEntityWithoutId
	Id string `gorm:"type:uuid;primaryKey"`
}

func NewBaseEntity(id string) *BaseEntity {
	entity := &BaseEntity{}
	entity.BaseEntityWithoutId = *NewBaseEntityWithoutId()
	if id != "" {
		entity.Id = id
	} else {
		entity.Id = uuid.New().String()
	}
	return entity
}

func (e *BaseEntity) Clone() *BaseEntity {
	clone := &BaseEntity{}
	clone.BaseEntityWithoutId = *e.BaseEntityWithoutId.Clone()
	clone.Id = uuid.New().String()
	return clone
}

func (e *BaseEntity) GetId() string {
	return e.Id
}

func (e *BaseEntity) SetId(id string) *BaseEntity {
	e.Id = id
	return e
}
