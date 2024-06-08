package schema

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	fUtil "github.com/rew3/rew3-internal/gen/utils"
)

/**
 * A Utility to generate graphql schema type code.
 */
type SchemaTypeGenerator struct {
	generated      []string
	overrides      map[string]string
	processedTypes map[reflect.Type]bool
	typeAndCodes   []GeneratedTypeCode
	config         SchemaConfig
	enumChecker    func(reflect.Type) bool
}

type GeneratedTypeCode struct {
	Index int
	Code  string
	Type  reflect.Type
}

type SchemaConfig struct {
	JsonTagFieldName    bool // if true, field name will be read from json tag.
	EnableRequiredField bool // if true, will check for requried and option field (omitempty tag)
}

func NewSchemaTypeGenerator(overrides map[string]string, config SchemaConfig, enumChecker func(reflect.Type) bool) *SchemaTypeGenerator {
	v1 := []string{}
	v2 := make(map[reflect.Type]bool)
	return &SchemaTypeGenerator{v1, overrides, v2, []GeneratedTypeCode{}, config, enumChecker}
}

func (gen *SchemaTypeGenerator) GetGeneratedTypes() []GeneratedTypeCode {
	sort.Slice(gen.typeAndCodes, func(i, j int) bool {
		return gen.typeAndCodes[i].Index < gen.typeAndCodes[j].Index
	})
	return gen.typeAndCodes
}

// Deprecated-------------------------------------
func (gen *SchemaTypeGenerator) GetResult() []string {
	sort.Slice(gen.typeAndCodes, func(i, j int) bool {
		return gen.typeAndCodes[i].Index < gen.typeAndCodes[j].Index
	})
	list := []string{}
	for _, i := range gen.typeAndCodes {
		list = append(list, i.Code)
	}
	gen.reverseList(list)
	return list
}

func (gen *SchemaTypeGenerator) ClearResult() {
	gen.generated = []string{}
	gen.processedTypes = make(map[reflect.Type]bool)
	gen.typeAndCodes = []GeneratedTypeCode{}
}

// GenerateGraphQLSchema generates GraphQL schema types from a Go type.
func (gen *SchemaTypeGenerator) GenerateGraphQLSchemaType(model interface{}) string {
	typ := reflect.TypeOf(model)
	return gen.generateType(0, typ, false)
}

// GenerateGraphQLSchema generates GraphQL schema input types from a Go type.
func (gen *SchemaTypeGenerator) GenerateGraphQLSchemaInputType(model interface{}) string {
	typ := reflect.TypeOf(model)
	return gen.generateType(0, typ, true)
}

func (gen *SchemaTypeGenerator) generateType(index int, typ reflect.Type, isInputType bool) string {
	// if type is pointer, convert it to non pointer type.
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	// for non struct type, e.g. primitive type, return.
	if typ.Kind() != reflect.Struct {
		if gen.enumChecker(typ) {
			return typ.Name() + "Enum"
		}
		return gen.GetGraphQLType(typ)
	}
	schemaType := "type"
	schemaName := typ.Name()
	if isInputType {
		schemaType = "input"
		schemaName = schemaName + "Input"
	}

	// check if graphql type is already generated for this type
	if _, ok := gen.processedTypes[typ]; ok {
		return schemaName
	}

	// mark this type as generated type.
	gen.processedTypes[typ] = true

	var sbb strings.Builder
	sb := &sbb
	sb.WriteString(fmt.Sprintf("%s %s {\n", schemaType, schemaName))
	gen.readFields(index, sb, typ, isInputType)
	sb.WriteString("}\n")
	gen.generated = append(gen.generated, sb.String())
	gen.typeAndCodes = append(gen.typeAndCodes, GeneratedTypeCode{index, sb.String(), typ})
	return schemaName
}

func (gen *SchemaTypeGenerator) readFields(index int, sb *strings.Builder, typ reflect.Type, isInputType bool) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if gen.isInlineField(field) {
			gen.readFields(index, sb, field.Type, isInputType)
			continue
		}
		if !gen.hasJsonTag(field) || gen.isIgnoredField(field) { // ignore fields that dont have json tag.
			continue
		}
		fieldName := field.Name
		fieldType := field.Type

		jsonFieldName, isRequired := gen.readFieldJsonTag(field)
		if gen.config.JsonTagFieldName {
			fieldName = jsonFieldName
		}

		requiredString := ""
		if isRequired && gen.config.EnableRequiredField {
			requiredString = "!"
		}

		var fieldTypeStr string

		fieldTypeName := fieldType.String()
		if fieldType.Kind() == reflect.Struct {
			fieldTypeName = fieldType.Name()
		} else if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct {
			fieldTypeName = fieldType.Elem().Name()
		}
		if override, ok := gen.overrides[fieldTypeName]; ok {
			fieldTypeStr = override + requiredString
		} else {
			if fieldType.Kind() == reflect.Ptr { // Pointer field type.
				fieldType = fieldType.Elem()
			}
			if fieldType.Kind() == reflect.Slice {
				if fieldType.Elem().Kind() == reflect.Struct { // List value type is Non pointer type.
					fieldTypeStr = "[" + gen.generateType(index+1, fieldType.Elem(), isInputType) + requiredString + "]"
				} else if fieldType.Elem().Kind() == reflect.Ptr && fieldType.Elem().Elem().Kind() == reflect.Struct { // List value type is pointer type.
					fieldTypeStr = "[" + gen.generateType(index+1, fieldType.Elem().Elem(), isInputType) + requiredString + "]"
				} else if fieldType.Kind() == reflect.String && fieldType.Name() != "" { // List value type is scalar type.
					fieldTypeStr = fieldTypeName + requiredString
				} else { // List value type is scalar type.
					fieldTypeStr = "[" + gen.GetGraphQLType(fieldType) + requiredString + "]"
				}
			} else if fieldType.Kind() == reflect.Struct {
				fieldTypeStr = gen.generateType(index+1, fieldType, isInputType) + requiredString
			} else {
				fieldTypeStr = gen.GetGraphQLType(fieldType) + requiredString
			}
		}
		sb.WriteString(fmt.Sprintf("\t%s: %s\n", fieldName, fieldTypeStr))
	}
}

// Check if given field has `json` tag name or not.
func (gen *SchemaTypeGenerator) hasJsonTag(field reflect.StructField) bool {
	tags := field.Tag.Get("json")
	return tags != ""
}

// Check if given field has tag has ignored json field.
func (gen *SchemaTypeGenerator) isIgnoredField(field reflect.StructField) bool {
	tags := field.Tag.Get("json")
	return tags == "-"
}

// Check if given field has tag has inline json field.
func (gen *SchemaTypeGenerator) isInlineField(field reflect.StructField) bool {
	tags := field.Tag.Get("json")
	return strings.Contains(tags, "inline")
}

func (gen *SchemaTypeGenerator) readFieldJsonTag(field reflect.StructField) (string, bool) {
	// Get the field name from the JSON tag, or use the struct field name as fallback
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return field.Name, false
	}
	// Parse JSON tag to get the field name and options
	tagOptions := strings.Split(jsonTag, ",")
	fieldName := tagOptions[0]

	// Check if the field is required based on the presence of omitempty option
	isRequired := true
	for _, option := range tagOptions[1:] {
		if option == "omitempty" {
			isRequired = false
			break
		}
	}
	return fieldName, isRequired
}

func (gen *SchemaTypeGenerator) GetGraphQLType(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.String:
		return "String"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "Int"
	case reflect.Float32, reflect.Float64:
		return "Float"
	case reflect.Bool:
		return "Boolean"
	default:
		return "String"
	}
}

func (gen *SchemaTypeGenerator) reverseList(list []string) {
	left, right := 0, len(list)-1
	for left < right {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
}

// Deprecated-------------------------------------
type SchemaTypesContext struct {
	Types []SchemaType
}

func NewSchemaTypesContext() *SchemaTypesContext {
	return &SchemaTypesContext{Types: []SchemaType{}}
}

// Deprecated-------------------------------------
type SchemaType struct {
	Type        interface{}
	IsInputType bool
	FileName    string
}

func (m *SchemaTypesContext) GenType(model interface{}, fileName string) {
	m.Types = append(m.Types, SchemaType{Type: model, IsInputType: false, FileName: fileName})
}

func (m *SchemaTypesContext) GenInputType(model interface{}, fileName string) {
	m.Types = append(m.Types, SchemaType{Type: model, IsInputType: true, FileName: fileName})
}

/**
 * Main method.
 * Generate Schems types.
 */
// Deprecated-------------------------------------
func GenerateSchemaTypes(context *SchemaTypesContext, generator *SchemaTypeGenerator) {
	for _, tpe := range context.Types {
		if !tpe.IsInputType {
			generator.GenerateGraphQLSchemaType(tpe.Type)
		} else {
			generator.GenerateGraphQLSchemaInputType(tpe.Type)
		}
		data := ""
		for _, v := range generator.GetResult() {
			fmt.Println(v)
			data = data + v + "\n"
		}
		generator.ClearResult()
		fUtil.WriteToFile(tpe.FileName, data)
	}
}

// Example usage:
// Deprecated-------------------------------------
type Address struct {
	Street  string
	City    string
	Country string
}
type User struct {
	ID       int        `json:"_id"`
	Name     string     `json:"name,omitempty"`
	Email    string     `json:"email,omitempty"`
	Birthday string     `json:"birth_day,omitempty"`
	Address  Address    `json:"address,omitempty"`
	Created  time.Time  `json:"created"`
	Tags     []Address  `json:"tags,omitempty"`
	UserType EntityType `json:"user_type"`
}
type EntityType string

const (
	CONTACT EntityType = "contact"
)

// Example code generation.
// Deprecated-------------------------------------
func RunExample() {
	// You can provide override to replace types present in model.
	// e.g. in below, Time type will be replaced by String.
	overrides := map[string]string{
		"Time": "String",
	}
	config := SchemaConfig{true, true}
	generator := NewSchemaTypeGenerator(overrides, config, func(t reflect.Type) bool { return false })
	typeContext := NewSchemaTypesContext()

	typeContext.GenInputType(User{}, "user_input_type.graphql")
	typeContext.GenType(User{}, "user_type.graphql")

	GenerateSchemaTypes(typeContext, generator)
}

/**
 * From given file, read schema types present in that file.
 */
func ReadSchemaTypes(filePath string) map[string]string {
	// Read the contents of the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Define a regular expression that matches the type blocks
	re := regexp.MustCompile(`(?s)(type|input)\s+(\w+)\s*{([^}]*)}`)

	// Find all matches in the file contents
	matches := re.FindAllStringSubmatch(string(data), -1)

	// Create a map to store the type blocks
	types := make(map[string]string)

	// Iterate over the matches and add them to the map
	for _, match := range matches {
		name := match[2]
		block := fmt.Sprintf("%s %s {%s}\n", match[1], match[2], match[3])
		types[name] = block
	}
	return types
}
