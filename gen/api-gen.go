package gen

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rew3/rew3-pkg/gen/schema"
	"github.com/rew3/rew3-pkg/gen/template"
	"github.com/rew3/rew3-pkg/gen/utils"
)

/**
 * API Code generator.
 * This api generator will generate Graphql Schema codes, Service API codes and Client GRPC Codes for given configuration.
 */
type APIGenerator struct {
	schemaGenerator *schema.SchemaTypeGenerator
	config          Config
	version         string
}

/**
 * Type overrides helps to define custom type for any particular type.
 * e.g. by default time.Time is override to String.
 * Version: version of code generation API.
 */
func NewGenerator(config Config, typeOverrides map[string]string, version string) *APIGenerator {
	schemaConfig := schema.SchemaConfig{JsonTagFieldName: true, EnableRequiredField: true}
	typeOverrides["Time"] = "String"
	sGen := schema.NewSchemaTypeGenerator(typeOverrides, schemaConfig, config.Enums.IsEnum)
	return &APIGenerator{sGen, config, version}
}

/**
 * Will load required templated to running project folder.
 * Note: project folder for generation must be in `gen` in base directory.
 */
func (gen *APIGenerator) LoadTemplates(internalVersion, genDir string) {
	utils.DeleteDirectory(genDir)
	utils.CopyModuleFiles("github.com/rew3/rew3-pkg", internalVersion, "gen/template", genDir)
	utils.DeleteFile("gen/template/template-gen.go")
}

/**
 * Generate Client GRPC API Codes.
 */
func (gen *APIGenerator) GenerateClientGrpcAPI() {
	config := gen.config
	for _, entity := range config.Entities {
		apiCodes := ClientGRPCAPICodes{Entity: entity.Entity}
		for _, rAPI := range entity.ReadAPIs {
			apiCodes.ServiceReadAPIs = append(apiCodes.ServiceReadAPIs, ServiceAPI{MethodName: utils.CapitalizeFirst(rAPI.Name), APIName: rAPI.Name})
		}
		for _, wAPI := range entity.WriteAPIs {
			apiCodes.ServiceReadAPIs = append(apiCodes.ServiceReadAPIs, ServiceAPI{MethodName: utils.CapitalizeFirst(wAPI.Name), APIName: wAPI.Name})
		}
		apiCodes.PackageName = entity.Directory.ClientGrpcAPIPackage
		if apiCodes.PackageName == "" {
			apiCodes.PackageName = "client"
		}
		outputPath := gen.config.ClientGrpcBaseDir + "/" + entity.Directory.ClientGrpcAPIDir + "/" + strings.ToLower(entity.Entity) + "_grpc_client.go"
		outputPath = utils.CleanDirPath(outputPath)
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/client-grpc-api.tmpl",
			OutputPath:    outputPath,
			Data:          apiCodes,
			DeleteIfExist: true,
			Version:       gen.version,
		})
	}
}

/**
 * Generate Service API Codes for CQRS.
 */
func (gen *APIGenerator) GenerateServiceAPI() {
	config := gen.config
	for _, entity := range config.Entities {
		apiCodes := ServiceCQRSAPICodes{Entity: entity.Entity}
		imports := make(map[string]string)
		generate := func(rAPI API) Code {
			inputTypeName := reflect.TypeOf(rAPI.Input.Data).Name()
			outputTypeName := reflect.TypeOf(rAPI.Output.Data).Name()
			if rAPI.Input.Import != nil {
				imports[rAPI.Input.Import.ImportUrl] = rAPI.Input.Import.ImportAlias
				inputTypeName = rAPI.Input.Import.ImportAlias + "." + inputTypeName
			}
			if rAPI.Output.Import != nil {
				imports[rAPI.Output.Import.ImportUrl] = rAPI.Output.Import.ImportAlias
				outputTypeName = rAPI.Output.Import.ImportAlias + "." + outputTypeName
			}
			if rAPI.Input.IsList {
				if rAPI.Input.IsListItemPointer {
					inputTypeName = "[]*" + inputTypeName
				} else {
					inputTypeName = "[]" + inputTypeName
				}
			}
			if rAPI.Input.IsPointer {
				inputTypeName = "*" + inputTypeName
			}
			if rAPI.Output.IsList {
				if rAPI.Output.IsListItemPointer {
					outputTypeName = "[]*" + outputTypeName
				} else {
					outputTypeName = "[]" + outputTypeName
				}
			}
			if rAPI.Output.IsPointer {
				outputTypeName = "*" + outputTypeName
			}
			code := fmt.Sprintf("[%s, %s](binder, api.ResolveEndpoint(%s))", inputTypeName, outputTypeName, "\""+rAPI.Name+"\"")
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
		if apiCodes.PackageName == "" {
			apiCodes.PackageName = "api"
		}

		outputPath := gen.config.APIBaseDir + "/" + entity.Directory.ServiceAPIDir + "/" + strings.ToLower(entity.Entity) + "_service_registry.go"
		outputPath = utils.CleanDirPath(outputPath)
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/service-api.tmpl",
			OutputPath:    outputPath,
			Data:          apiCodes,
			DeleteIfExist: true,
			Version:       gen.version,
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
				inputType := ""
				tpe := reflect.TypeOf(rAPI.Input.Data)
				if tpe.Kind() != reflect.Struct {
					inputType = gen.schemaGenerator.GetGraphQLType(tpe)
				} else {
					inputType = tpe.Name() + "Input" // this is input type.
				}
				code = code + "(data: " + inputType + ")"
			}
			outputTypeName := ""
			oTpe := reflect.TypeOf(rAPI.Output.Data)
			if oTpe.Kind() != reflect.Struct {
				outputTypeName = gen.schemaGenerator.GetGraphQLType(oTpe)
			} else {
				outputTypeName = oTpe.Name()
			}
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
			schemaAPICodes.SchemaMutations = append(schemaAPICodes.SchemaMutations, code)
		}

		outputPath := gen.config.SchemaBaseDir + "/" + entity.Directory.SchemaDir + "/" + config.Module + "_" + strings.ToLower(entity.Entity) + "_schema.graphql"
		outputPath = utils.CleanDirPath(outputPath)
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/schema-api.tmpl",
			OutputPath:    outputPath,
			Data:          schemaAPICodes,
			DeleteIfExist: true,
			Version:       gen.version,
		})
	}
}

/**
 * Generate Schema Type Codes.
 */
func (gen *APIGenerator) generateSchemaTypes(config Config) {
	types := []SchemaType{}
	// extracting all schema types (input, output, model etc.)
	for _, entity := range config.Entities {
		// read input, output type of api, and append to tBase (schema typebase)
		generateTypeCode := func(api API) {
			gen.schemaGenerator.ClearResult()
			gen.schemaGenerator.GenerateGraphQLSchemaInputType(api.Input.Data)
			generatedInput := gen.schemaGenerator.GetGeneratedTypes()
			prepareSchemaType := func(isInputType bool, tCtx SchemaTypeContext) SchemaType {
				return SchemaType{
					Entity:      entity.Entity,
					BasePackage: entity.BasePackage,
					SchemaDir:   utils.CleanDirPath(gen.config.SchemaBaseDir + "/" + entity.Directory.SchemaDir),
					Type:        tCtx,
					IsInputType: isInputType,
				}
			}
			for _, code := range generatedInput {
				if reflect.TypeOf(api.Input.Data) == code.Type {
					types = append(types, prepareSchemaType(true, SchemaTypeContext{RawType: &api.Input, Type: code.Type, Code: code.Code}))
				} else {
					types = append(types, prepareSchemaType(true, SchemaTypeContext{Type: code.Type, Code: code.Code}))
				}
			}
			gen.schemaGenerator.ClearResult()
			if gen.schemaGenerator.IsPrimitiveType(reflect.TypeOf(api.Output.Data)) && api.WrapOutput {
				// seems like output data type is some primitive,
				types = append(types, prepareSchemaType(false,
					SchemaTypeContext{
						GenerateWrapper: api.WrapOutput,
						WrapperName:     api.WrapOutputName,
						RawType:         &api.Output,
						Type:            reflect.TypeOf(api.Output.Data),
						IsPrimitiveType: true,
						Code:            gen.schemaGenerator.GetGraphQLType(reflect.TypeOf(api.Output.Data))}))
			} else {
				gen.schemaGenerator.GenerateGraphQLSchemaType(api.Output.Data)
				generated := gen.schemaGenerator.GetGeneratedTypes()
				for _, code := range generated {
					if reflect.TypeOf(api.Output.Data) == code.Type {
						types = append(types, prepareSchemaType(false,
							SchemaTypeContext{GenerateWrapper: api.WrapOutput,
								WrapperName: api.WrapOutputName, RawType: &api.Output, Type: code.Type, Code: code.Code}))
					} else {
						types = append(types, prepareSchemaType(false, SchemaTypeContext{Type: code.Type, Code: code.Code}))
					}
				}
			}
			gen.schemaGenerator.ClearResult()
		}
		// generate schema type for types in all read apis of this entity
		for _, api := range entity.ReadAPIs {
			generateTypeCode(api)
		}
		// generate schema type for types in all write apis of this entity
		for _, api := range entity.WriteAPIs {
			generateTypeCode(api)
		}
	}
	distinct := gen.makeDistinctSchemaTypes(types)
	coreTypes, sharedTypes, entityTypeMaps := gen.filterSchemaTypes(distinct)

	coreTypeCodes := gen.prepareSchemaTypesCodes(coreTypes)
	coreTypeCodes.generateEnumSchemaType(gen.config.Enums.CoreEnums)
	sharedTypeCodes := gen.prepareSchemaTypesCodes(sharedTypes)
	sharedTypeCodes.generateEnumSchemaType(gen.config.Enums.SharedEnums)

	if !coreTypeCodes.IsEmpty() {
		outputPath := config.SchemaBaseDir + "/core/" + "core_types.graphql"
		outputPath = utils.CleanDirPath(outputPath)
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/schema-type.tmpl",
			OutputPath:    outputPath,
			Data:          coreTypeCodes,
			DeleteIfExist: true,
			Version:       gen.version,
		})
	}
	if !sharedTypeCodes.IsEmpty() {
		outputPath := config.SchemaBaseDir + "/shared/" + "shared_types.graphql"
		outputPath = utils.CleanDirPath(outputPath)
		template.GenerateFromTemplate(template.TemplateConfig{
			TemplatePath:  "gen/template/schema-type.tmpl",
			OutputPath:    outputPath,
			Data:          sharedTypeCodes,
			DeleteIfExist: true,
			Version:       gen.version,
		})
	}
	for entity, v := range entityTypeMaps {
		if len(v) != 0 {
			schemaDir := v[0].SchemaDir
			entityTypeCodes := gen.prepareSchemaTypesCodes(v)
			entityTypeCodes.generateEnumSchemaType(gen.config.Enums.ForEntity(entity))
			outputPath := schemaDir + "/" + config.Module + "_" + strings.ToLower(entity) + "_types.graphql"
			outputPath = utils.CleanDirPath(outputPath)
			template.GenerateFromTemplate(template.TemplateConfig{
				TemplatePath:  "gen/template/schema-type.tmpl",
				OutputPath:    outputPath,
				Data:          entityTypeCodes,
				DeleteIfExist: true,
				Version:       gen.version,
			})
		}
	}
}

/**
 * Filter and make given list of schema types unique/distinct.
 */
func (gen *APIGenerator) makeDistinctSchemaTypes(types []SchemaType) []SchemaType {
	seen := make(map[string]bool)
	unique := []SchemaType{}
	for _, tc := range types {
		typeName := tc.Type.Type.PkgPath() + "/" + tc.Type.Type.Name()
		if tc.IsInputType {
			typeName = typeName + "Input"
		}
		if tc.Type.GenerateWrapper {
			typeName = tc.Type.WrapperName
		}
		if !seen[typeName] {
			seen[typeName] = true
			unique = append(unique, tc)
		}
	}
	return unique
}

/**
 * From given list of schema types, seperate input and non input types.
 * return - (input types, nonInput types)
 */
func (gen *APIGenerator) seperateInputSchemaTypes(types []SchemaType) ([]SchemaType, []SchemaType) {
	nonInputs := []SchemaType{}
	inputs := []SchemaType{}
	for _, i := range types {
		if i.IsInputType {
			inputs = append(inputs, i)
		} else {
			nonInputs = append(nonInputs, i)
		}
	}
	return inputs, nonInputs
}

/**
 * Filter schema types.
 * Seperate core, shared and entity specific schema types.
 * Note: core schema types are - rew3-internal/platform core types, shared are types that dont belongs to any entity base packages.
 * Return - (core schemastypes, shared schemastypes, entity/Schematype-list Map  )
 */
func (gen *APIGenerator) filterSchemaTypes(types []SchemaType) ([]SchemaType, []SchemaType, map[string][]SchemaType) {
	coreTypes := []SchemaType{}
	sharedTypes := []SchemaType{}
	entityTypes := []SchemaType{}

	allEntityBasePackages := []string{}
	for _, i := range gen.config.Entities {
		allEntityBasePackages = append(allEntityBasePackages, i.BasePackage)
	}
	for _, i := range types {
		isEntityType := false
		for _, bp := range allEntityBasePackages {
			if strings.Contains(i.Type.Type.PkgPath(), bp) {
				isEntityType = true
			}
			// if wrapper type and primitive type (package path for primitive type is empty)
			// and this belongs to respective entity (not shared, or core)
			if i.Type.Type.PkgPath() == "" && i.Type.GenerateWrapper {
				isEntityType = true
			}
		}
		if strings.Contains(i.Type.Type.PkgPath(), "github.com/rew3/rew3-pkg") {
			coreTypes = append(coreTypes, i)
		} else if isEntityType {
			entityTypes = append(entityTypes, i)
		} else {
			sharedTypes = append(sharedTypes, i)
		}
	}
	entityTypeMap := make(map[string][]SchemaType)
	for _, i := range entityTypes {
		list := entityTypeMap[i.Entity]
		if list == nil {
			entityTypeMap[i.Entity] = []SchemaType{i}
		} else {
			list = append(list, i)
			entityTypeMap[i.Entity] = list
		}
	}
	return coreTypes, sharedTypes, entityTypeMap
}

/**
 * From given list of schema types, generate schema type codes.
 */
func (gen *APIGenerator) prepareSchemaTypesCodes(types []SchemaType) SchemaTypeCodes {
	inputs, nonInputs := gen.seperateInputSchemaTypes(types)
	typeCodes := SchemaTypeCodes{}
	inputCtx := []SchemaTypeContext{}
	nonInputCtx := []SchemaTypeContext{}
	for _, i := range inputs {
		inputCtx = append(inputCtx, i.Type)
	}
	for _, i := range nonInputs {
		nonInputCtx = append(nonInputCtx, i.Type)
	}
	typeCodes.generate(nonInputCtx, inputCtx)
	return typeCodes
}

/**
 * Client GRPC API codes.
 */
type ClientGRPCAPICodes struct {
	Entity           string
	PackageName      string
	ServiceReadAPIs  []ServiceAPI
	ServiceWriteAPIs []ServiceAPI
}

type ServiceAPI struct {
	MethodName string
	APIName    string
}

/**
 * Service API codes for CQRS.
 */
type ServiceCQRSAPICodes struct {
	Entity      string
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

func (c SchemaAPICodes) HasQueries() bool {
	return len(c.SchemaQueries) > 0
}

func (c SchemaAPICodes) HasMutations() bool {
	return len(c.SchemaMutations) > 0
}

/**
 * Schema type codes for all type (input, models, output etc)
 */
type SchemaTypeCodes struct {
	Models  []Code
	Inputs  []Code
	Outputs []Code
	Enums   []Code
}

func (c SchemaTypeCodes) HasModels() bool {
	return len(c.Models) > 0
}

func (c SchemaTypeCodes) HasInputs() bool {
	return len(c.Inputs) > 0
}

func (c SchemaTypeCodes) HasOutputs() bool {
	return len(c.Outputs) > 0
}

func (c SchemaTypeCodes) HasEnums() bool {
	return len(c.Enums) > 0
}

func (sc *SchemaTypeCodes) IsEmpty() bool {
	return len(sc.Inputs) == 0 && len(sc.Models) == 0 && len(sc.Outputs) == 0 && len(sc.Enums) == 0
}

/**
 * From given input and non input types, generate code for Inputs, Models and Output types.
 */
func (sc *SchemaTypeCodes) generate(types []SchemaTypeContext, inputTypes []SchemaTypeContext) {
	for _, i := range inputTypes {
		sc.Inputs = append(sc.Inputs, Code{Code: i.Code})
	}
	for _, i := range types {
		if i.GenerateWrapper {
			sc.generateWrapperSchemaType(i)
		} else {
			sc.Models = append(sc.Models, Code{Code: i.Code})
		}
	}
}

/**
 * Generate the wrapped schema type (i.e. Execution Result type - which is also called output e.g. AddContactOutput)
 */
func (sc *SchemaTypeCodes) generateWrapperSchemaType(typeContext SchemaTypeContext) {
	tName := typeContext.WrapperName
	dtName := typeContext.Type.Name()
	if typeContext.IsPrimitiveType {
		dtName = typeContext.Code
	}
	if typeContext.RawType != nil && typeContext.RawType.IsList {
		dtName = "[" + dtName + "]"
	}
	if typeContext.RawType != nil && typeContext.RawType.IsRequired {
		dtName = dtName + "!"
	}
	str := fmt.Sprintf(`type %s {
	action: String!
	message: String!
	status: StatusTypeEnum!
	data: %s
}`, tName, dtName)
	sc.Outputs = append(sc.Outputs, Code{Code: str})
}

/**
 * Generate enum schema type.
 */
func (sc *SchemaTypeCodes) generateEnumSchemaType(enums []EnumType) {
	for _, i := range enums {
		tName := i.Type.Name() + "Enum"
		str := fmt.Sprintf(`enum %s {`, tName)
		for _, e := range i.Items {
			str = str + "\n\t" + e
		}
		str = str + "\n}\n"
		sc.Enums = append(sc.Enums, Code{Code: str})
	}
}

type SchemaType struct {
	Module      string
	Entity      string
	BasePackage string
	SchemaDir   string
	Type        SchemaTypeContext
	IsInputType bool
}

type SchemaTypeContext struct {
	GenerateWrapper bool
	WrapperName     string
	RawType         *DataType
	Type            reflect.Type
	IsPrimitiveType bool
	Code            string
}

type Code struct {
	Code string
}
