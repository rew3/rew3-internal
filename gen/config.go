package gen

type Config struct {
	Module        string
	BaseSchemaDir string
	Entities      []EntityConfig
}

type EntityConfig struct {
	Directory   Directory
	Entity      string // Note entity must be unique (no duplicate within same application, otherwise undesired things can happen)
	BasePackage string // base package for this entity. in / notation for directory.
	ReadAPIs    []API[interface{}, interface{}]
	WriteAPIs   []API[interface{}, interface{}]
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
	WrapOutput     bool // Wrap to execution result.
	WrapOutputName string
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
