package gen

type Config struct {
	Module        string
	BaseSchemaDir string
	Entities      []EntityConfig
}

type EntityConfig struct {
	Directory   Directory
	Entity      string        // Note entity must be unique (no duplicate within same application, otherwise undesired things can happen)
	BasePackage string        // base package for this entity/model in the project. in / notation for directory.
	ReadAPIs    []interface{} // must be list of API[interface{}, interface{}]
	WriteAPIs   []interface{} // must be list of API[interface{}, interface{}]
}

func (c *EntityConfig) GetReadAPIs() []API[interface{}, interface{}] {
	list := []API[interface{}, interface{}]{}
	for _, item := range c.ReadAPIs {
		switch v := item.(type) {
		case API[interface{}, interface{}]:
			list = append(list, v)
		}
	}
	return list
}

func (c *EntityConfig) GetWriteAPIs() []API[interface{}, interface{}] {
	list := []API[interface{}, interface{}]{}
	for _, item := range c.WriteAPIs {
		switch v := item.(type) {
		case API[interface{}, interface{}]:
			list = append(list, v)
		}
	}
	return list
}

type Directory struct {
	SchemaDir            string // can be nested.
	ServiceAPIDir        string // can be nested.
	ClientGrpcAPIDir     string // can be nested.
	ServiceAPIPackage    string
	ClientGrpcAPIPackage string
}

type API[I any, O any] struct {
	Name           string
	Input          DataType[I]
	Output         DataType[O]
	WrapOutput     bool   // Wrap to execution result.
	WrapOutputName string // this is for schema output type name.
}

type DataType[T any] struct {
	IsNull      bool // use null if no input need to define. but this should not happen, every command/query has type.
	Data        T
	IsRequired  bool
	IsList      bool
	ImportUrl   string
	ImportAlias string
}

// input type generation
// output type generation

// mutation - execution result conveert?
// query - direct model response.

// required, non required.

// module, parentEntity, entity.
// in config, import package, input type, output type?

// generate Service API.
// generate Schema API
// generate schema types.

// isolate, base core types, and project specific shared types.
