package model

import (
	"encoding/json"

	c "github.com/rew3/rew3-base/service/constants"
)

type Reference struct {
	Data     *json.RawMessage `bson:"data,omitempty"`
	Module   c.Module         `bson:"module"`
	Entity   *c.Entity        `bson:"entity,omitempty"`
	EntityID *string          `bson:"entity_id,omitempty"`
	Title    *string          `bson:"title,omitempty"`
}

/*
ToDo this method should be create in go later.
Need to figure how enum can be called same like withName method in scala enum
func NewReference(entityName string, id string) *Reference {
	return &Reference{
		Module:   Module.GetModule(Entity(entityName).String()),
		Entity:   Entity.WithName(Entity(entityName).String()),
		EntityID: &id,
	}
}

func NewReferenceWithModuleName(moduleName, entityName, id string) *Reference {
	return &Reference{
		Module:   Module.WithName(moduleName),
		Entity:   Entity.WithName(Entity(entityName).String()),
		EntityID: &id,
	}
}

func NewReferenceWithModuleNameAndTitle(moduleName, entityName, id, title string) *Reference {
	return &Reference{
		Module:   Module.WithName(moduleName),
		Entity:   Entity.WithName(Entity(entityName).String()),
		EntityID: &id,
		Title:    &title,
	}
}

func NewReferenceWithModuleNameOnly(moduleName string) *Reference {
	return &Reference{
		Module: Module.WithName(moduleName),
	}
}

func NewReferenceWithModuleNameAndEntityName(moduleName string, entityName *string, id *string) *Reference {
	var entity *Entity = nil
	if entityName != nil {
		entity = Entity.WithName(Entity(*entityName).String())
	}
	return &Reference{
		Module:   Module.WithName(moduleName),
		Entity:   entity,
		EntityID: id,
	}
}
*/