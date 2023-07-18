package entity

import (
	"time"
)

type EntityEditionListener struct{}

func (e *EntityEditionListener) PreUpdate(args *ORM.LifecycleEventArgs) {
	entity := args.Object()
	if entity, ok := entity.(Entity.BaseEntityInterface); ok {
		entity.SetModified(time.Now())
	}
}
