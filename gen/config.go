package gen

import "reflect"

type Config struct {
	Module        string
	BaseSchemaDir string
	Entities      []EntityConfig
	Enums         EnumTypeMapping
}

type EntityConfig struct {
	Directory   Directory
	Entity      string // Note entity must be unique (no duplicate within same application, otherwise undesired things can happen)
	BasePackage string // base package for this entity/model in the project. in / notation for directory. this is to make sure, all the sub entities are within this package or note.
	ReadAPIs    []API
	WriteAPIs   []API
}

type Directory struct {
	SchemaDir            string // can be nested.
	ServiceAPIDir        string // can be nested.
	ClientGrpcAPIDir     string // can be nested.
	ServiceAPIPackage    string
	ClientGrpcAPIPackage string
}

type API struct {
	Name           string
	Input          DataType
	Output         DataType
	WrapOutput     bool   // Wrap to execution result.
	WrapOutputName string // this is for schema output type name.
}

type DataType struct {
	Data       interface{}
	IsNull     bool // use null if no input need to define. but this should not happen, every command/query has type.
	IsRequired bool
	IsList     bool
	Import     *Import
}

type Import struct {
	ImportUrl   string
	ImportAlias string
}

/**
 * Enum type mapping.
 */
type EnumTypeMapping struct {
	CoreEnums   []EnumType
	SharedEnums []EnumType
	EntityEnums []EntityEnums
}

// Check if given type name is enum or not (defined in enum mapping)
func (em *EnumTypeMapping) IsEnum(enum reflect.Type) bool {
	isEnum := false
	for _, i := range em.CoreEnums {
		if i.Type == enum {
			isEnum = true
		}
	}
	for _, i := range em.SharedEnums {
		if i.Type == enum {
			isEnum = true
		}
	}
	for _, i := range em.EntityEnums {
		for _, ii := range i.Enums {
			if ii.Type == enum {
				isEnum = true
			}
		}
	}
	return isEnum
}

// Check the enums for given entity type / empty is not exists.
func (e *EnumTypeMapping) ForEntity(entity string) []EnumType {
	list := []EnumType{}
	for _, i := range e.EntityEnums {
		if i.Entity == entity {
			list = i.Enums
		}
	}
	return list
}

type EnumType struct {
	Type  reflect.Type
	Items []string
}
type EntityEnums struct {
	Entity string
	Enums  []EnumType
}
