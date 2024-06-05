package gen

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rew3/rew3-internal/gen/schema"
	"github.com/rew3/rew3-internal/gen/template"
	"github.com/rew3/rew3-internal/gen/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/**
 * API Code generator.
 * This api generator will generate Graphql Schema codes, Service API codes and Client GRPC Codes for given configuration.
 */
type APIGenerator struct {
	schemaGenerator *schema.SchemaTypeGenerator
	config          Config
}

func NewGenerator(config Config) *APIGenerator {
	schemaConfig := schema.SchemaConfig{JsonTagFieldName: true, EnableRequiredField: true}
	overrides := make(map[string]string)
	overrides["Time"] = "String"
	sGen := schema.NewSchemaTypeGenerator(overrides, schemaConfig)
	return &APIGenerator{sGen, config}
}

/**
 * Will load required templated to running project folder.
 * Note: project folder for generation must be in `gen` in base directory.
 */
func (gen *APIGenerator) LoadTemplates(internalVersion string) {
	utils.CopyModuleFiles("github.com/rew3/rew3-internal@"+internalVersion+"/gen/templates", "gen/template")
	utils.DeleteFile("gen/template/template-gen.go")
}

/**
 * Generate Client GRPC API Codes.
 */
func (gen *APIGenerator) GenerateClientGrpcAPI() {
	config := gen.config
	for _, entity := range config.Entities {
		apiCodes := ClientGRPCAPICodes{}
		c := cases.Title(language.English) // capitalize.
		for _, rAPI := range entity.ReadAPIs {
			apiCodes.ServiceReadAPIs = append(apiCodes.ServiceReadAPIs, ServiceAPI{APIName: c.String(rAPI.Name)})
		}
		for _, wAPI := range entity.WriteAPIs {
			apiCodes.ServiceReadAPIs = append(apiCodes.ServiceReadAPIs, ServiceAPI{APIName: c.String(wAPI.Name)})
		}
		apiCodes.PackageName = entity.Directory.ClientGrpcAPIPackage
		outputPath := entity.Directory.ClientGrpcAPIDir + "/" + strings.ToLower(entity.Entity) + "_client_api.go"
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/client-grpc-api.tmpl",
			OutputPath:    outputPath,
			Data:          apiCodes,
			DeleteIfExist: true,
		})
	}
}

/**
 * Generate Service API Codes for CQRS.
 */
func (gen *APIGenerator) GenerateServiceAPI() {
	config := gen.config
	for _, entity := range config.Entities {
		apiCodes := ClientCQRSAPICodes{}
		imports := make(map[string]string)
		generate := func(rAPI API) Code {
			if rAPI.Input.ImportUrl != "" {
				imports[rAPI.Input.ImportUrl] = rAPI.Input.ImportAlias
			}
			if rAPI.Output.ImportUrl != "" {
				imports[rAPI.Output.ImportUrl] = rAPI.Output.ImportAlias
			}
			inputTypeName := rAPI.Input.ImportAlias + "." + reflect.TypeOf(rAPI.Input.Data).Name()
			outputTypeName := rAPI.Output.ImportAlias + "." + reflect.TypeOf(rAPI.Output.Data).Name()
			if rAPI.Input.IsList {
				inputTypeName = "[]" + inputTypeName
			}
			if rAPI.Output.IsList {
				outputTypeName = "[]" + outputTypeName
			}
			code := fmt.Sprintf("[%s, %s](binder, api.ResolveEndpoint(%s))", inputTypeName, outputTypeName, rAPI.Name)
			return Code{Code: code}
		}
		for _, rAPI := range entity.ReadAPIs {
			code := generate(rAPI)
			code.Code = "api.BindQueryAPI" + code.Code
			apiCodes.QueryAPIs = append(apiCodes.QueryAPIs, code)
		}
		for _, wAPI := range entity.WriteAPIs {
			code := generate(wAPI)
			code.Code = "api.BindCommandAPI" + code.Code
			apiCodes.CommandAPIs = append(apiCodes.CommandAPIs, code)
		}
		for im, al := range imports {
			code := Code{Code: al + " " + "\"" + im + "\""}
			apiCodes.Imports = append(apiCodes.Imports, code)
		}
		apiCodes.PackageName = entity.Directory.ServiceAPIPackage

		outputPath := entity.Directory.ServiceAPIDir + "/" + strings.ToLower(entity.Entity) + "_apis.go"
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/service-api.tmpl",
			OutputPath:    outputPath,
			Data:          apiCodes,
			DeleteIfExist: true,
		})
	}
}

/**
 * Generate Schema Codes.
 */
func (gen *APIGenerator) GenerateSchema() {
	gen.generateSchemaAPI(gen.config)
	gen.generateSchemaTypes(gen.config)
}

/**
 * Generate Schema API Codes.
 */
func (gen *APIGenerator) generateSchemaAPI(config Config) {
	for _, entity := range config.Entities {
		schemaAPICodes := SchemaAPICodes{}
		generate := func(rAPI API) Code {
			code := rAPI.Name
			if !rAPI.Input.IsNull {
				code = "(data: " + reflect.TypeOf(rAPI.Input.Data).Name() + "Input)"
			}
			outputTypeName := reflect.TypeOf(rAPI.Output.Data).Name()
			if rAPI.Output.IsList {
				outputTypeName = "[" + outputTypeName + "]"
			}
			if rAPI.Output.IsRequired {
				outputTypeName = outputTypeName + "!"
			}
			if rAPI.WrapOutput {
				outputTypeName = rAPI.WrapOutputName
			}
			code = code + ": " + outputTypeName
			return Code{Code: code}
		}
		for _, rAPI := range entity.ReadAPIs {
			code := generate(rAPI)
			schemaAPICodes.SchemaQueries = append(schemaAPICodes.SchemaQueries, code)
		}
		for _, wAPI := range entity.WriteAPIs {
			code := generate(wAPI)
			schemaAPICodes.SchemaMutations = append(schemaAPICodes.SchemaQueries, code)
		}

		outputPath := entity.Directory.SchemaDir + "/" + config.Module + "_" + strings.ToLower(entity.Entity) + "_schema.graphql"
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/schema-api.tmpl",
			OutputPath:    outputPath,
			Data:          schemaAPICodes,
			DeleteIfExist: true,
		})
	}
}

/**
 * Generate Schema Type Codes.
 */
func (gen *APIGenerator) generateSchemaTypes(config Config) {
	typeBases := []SchemaTypeBase{}
	for _, entity := range config.Entities {
		tBase := SchemaTypeBase{
			Directory:   entity.Directory,
			Module:      config.Module,
			Entity:      entity.Entity,
			BasePackage: entity.BasePackage,
		}
		typeBases = append(typeBases, tBase)
		generateTypeCode := func(api API) {
			gen.schemaGenerator.ClearResult()
			gen.schemaGenerator.GenerateGraphQLSchemaInputType(api.Input.Data)
			generatedInput := gen.schemaGenerator.GetGeneratedTypes()
			for gtype, code := range generatedInput {
				tBase.InputTypes = append(tBase.InputTypes, SchemaTypeContext{
					GenerateWrapper: false,
					RawType:         api.Input,
					Type:            gtype,
					Code:            code,
				})
			}
			gen.schemaGenerator.GenerateGraphQLSchemaType(api.Output.Data)
			generated := gen.schemaGenerator.GetGeneratedTypes()
			for gtype, code := range generated {
				tBase.Types = append(tBase.Types, SchemaTypeContext{
					GenerateWrapper: api.WrapOutput,
					WrapperName:     api.WrapOutputName,
					RawType:         api.Output,
					Type:            gtype,
					Code:            code,
				})
			}
		}
		for _, api := range entity.ReadAPIs {
			generateTypeCode(api)
		}
		for _, api := range entity.WriteAPIs {
			generateTypeCode(api)
		}
	}
	coreTypeCodes := SchemaTypeCodes{}
	for _, tb := range typeBases {
		entityTypeCodes, sharedTypeCodes := gen.generateSchemaTypesCodes(tb, coreTypeCodes)
		outputPath := tb.Directory.SchemaDir + "/" + tb.Module + "_" + strings.ToLower(tb.Entity) + "_types.graphql"
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/schema-type.tmpl",
			OutputPath:    outputPath,
			Data:          entityTypeCodes,
			DeleteIfExist: true,
		})
		if !sharedTypeCodes.IsEmpty() {
			outputPath := config.BaseSchemaDir + "/shared/" + "shared_types.graphql"
			template.GenerateFromTemplate(template.TemplateConfig{
				TemplatePath:  "gen/template/schema-type.tmpl",
				OutputPath:    outputPath,
				Data:          entityTypeCodes,
				DeleteIfExist: true,
			})
		}
	}
	if !coreTypeCodes.IsEmpty() {
		outputPath := config.BaseSchemaDir + "/core/" + "core_types.graphql"
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/schema-type.tmpl",
			OutputPath:    outputPath,
			Data:          coreTypeCodes,
			DeleteIfExist: true,
		})
	}
}

// return pair of (entityTypeCodes, sharedTypeCodes)
func (gen *APIGenerator) generateSchemaTypesCodes(schemaTypeBase SchemaTypeBase, coreTypeCodes SchemaTypeCodes) (SchemaTypeCodes, SchemaTypeCodes) {
	// Seperating core platform/internal types.
	coreTypes := []SchemaTypeContext{}
	coreInputTypes := []SchemaTypeContext{}
	remainingTypes := []SchemaTypeContext{}
	remainingInputTypes := []SchemaTypeContext{}
	for _, i := range schemaTypeBase.Types {
		if strings.Contains(i.Type.PkgPath(), "github.com/rew3/rew3-internal") {
			coreTypes = append(coreTypes, i)
		} else {
			remainingTypes = append(remainingTypes, i)
		}
	}
	for _, i := range schemaTypeBase.InputTypes {
		if strings.Contains(i.Type.PkgPath(), "github.com/rew3/rew3-internal") {
			coreInputTypes = append(coreInputTypes, i)
		} else {
			remainingInputTypes = append(remainingInputTypes, i)
		}
	}
	// Removing duplicate types.
	seen := make(map[string]bool)
	unique := []SchemaTypeContext{}
	uniqueInputs := []SchemaTypeContext{}
	for _, tc := range remainingTypes {
		typeName := tc.Type.PkgPath() + "/" + tc.Type.Name()
		if !seen[typeName] {
			seen[typeName] = true
			unique = append(unique, tc)
		}
	}
	for _, tc := range remainingInputTypes {
		typeName := tc.Type.PkgPath() + "/" + tc.Type.Name()
		if !seen[typeName] {
			seen[typeName] = true
			uniqueInputs = append(uniqueInputs, tc)
		}
	}
	// Separating types.
	sharedTypes := []SchemaTypeContext{}
	sharedInputTypes := []SchemaTypeContext{}
	entityTypes := []SchemaTypeContext{}
	entityInputTypes := []SchemaTypeContext{}
	for _, t := range unique {
		if strings.Contains(t.Type.PkgPath(), schemaTypeBase.BasePackage) {
			entityTypes = append(entityTypes, t)
		} else {
			sharedTypes = append(sharedTypes, t)
		}
	}
	for _, t := range uniqueInputs {
		if strings.Contains(t.Type.PkgPath(), schemaTypeBase.BasePackage) {
			entityInputTypes = append(entityInputTypes, t)
		} else {
			sharedInputTypes = append(sharedInputTypes, t)
		}
	}
	entityTypeCodes := SchemaTypeCodes{}
	sharedTypeCodes := SchemaTypeCodes{}
	entityTypeCodes.generate(entityTypes, entityInputTypes)
	sharedTypeCodes.generate(sharedTypes, sharedInputTypes)
	coreTypeCodes.generate(coreTypes, coreInputTypes) // for core.

	return entityTypeCodes, sharedTypeCodes
}

/**
 * Client GRPC API codes.
 */
type ClientGRPCAPICodes struct {
	PackageName      string
	ServiceReadAPIs  []ServiceAPI
	ServiceWriteAPIs []ServiceAPI
}

type ServiceAPI struct {
	APIName string
}

/**
 * Service API codes for CQRS.
 */
type ClientCQRSAPICodes struct {
	PackageName string
	Imports     []Code
	CommandAPIs []Code
	QueryAPIs   []Code
}

/**
 * Schema API codes.
 */
type SchemaAPICodes struct {
	SchemaQueries   []Code
	SchemaMutations []Code
}

/**
 * Schema type codes for all type (input, models, output etc)
 */
type SchemaTypeCodes struct {
	Inputs  []Code
	Models  []Code
	Outputs []Code
}

func (sc *SchemaTypeCodes) IsEmpty() bool {
	return len(sc.Inputs) == 0 && len(sc.Models) == 0 && len(sc.Outputs) == 0
}

func (sc *SchemaTypeCodes) generate(types []SchemaTypeContext, inputTypes []SchemaTypeContext) {
	for _, i := range inputTypes {
		sc.Inputs = append(sc.Inputs, Code{Code: i.Code})
	}
	for _, i := range types {
		sc.Inputs = append(sc.Models, Code{Code: i.Code})
		if i.GenerateWrapper {
			sc.generateWrapperSchemaType(i)
		}
	}
}

func (sc *SchemaTypeCodes) generateWrapperSchemaType(typeContext SchemaTypeContext) {
	tName := typeContext.WrapperName
	dtName := typeContext.Type.Name()
	if typeContext.RawType.IsList {
		dtName = "[" + dtName + "]"
	}
	if typeContext.RawType.IsRequired {
		dtName = dtName + "!"
	}
	str := fmt.Sprintf(`type %s {
		action: String!
		message: String!
		status: PlatformStatusEnum!
		data: %s
	}`, tName, dtName)
	sc.Outputs = append(sc.Outputs, Code{Code: str})
}

type SchemaTypeBase struct {
	Directory   Directory
	Module      string
	Entity      string
	BasePackage string
	Types       []SchemaTypeContext
	InputTypes  []SchemaTypeContext
}

type SchemaTypeContext struct {
	GenerateWrapper bool
	WrapperName     string
	RawType         DataType
	Type            reflect.Type
	Code            string
}

type Code struct {
	Code string
}
