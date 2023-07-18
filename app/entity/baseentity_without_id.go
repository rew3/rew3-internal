package entity

import "time"

type BaseEntityWithoutId struct {
	Created  time.Time `gorm:"column:created;type:datetime"`
	Modified time.Time `gorm:"column:modified;type:datetime"`
}

func NewBaseEntityWithoutId() *BaseEntityWithoutId {
	return &BaseEntityWithoutId{
		Created:  time.Now(),
		Modified: time.Now(),
	}
}

func (b *BaseEntityWithoutId) Clone() *BaseEntityWithoutId {
	return &BaseEntityWithoutId{
		Created:  b.Created,
		Modified: b.Modified,
	}
}

func (b *BaseEntityWithoutId) SetCreated(created time.Time) *BaseEntityWithoutId {
	b.Created = created
	return b
}

func (b *BaseEntityWithoutId) GetCreated() time.Time {
	return b.Created
}

func (b *BaseEntityWithoutId) SetModified(modified time.Time) *BaseEntityWithoutId {
	b.Modified = modified
	return b
}

func (b *BaseEntityWithoutId) GetModified() time.Time {
	return b.Modified
}
