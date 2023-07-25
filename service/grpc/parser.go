package grpc

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/tidwall/gjson"
)

/**
 * Parser to parse payload data (request and response)
 */
type Parser struct {
	result gjson.Result
}

func PayloadParser(data json.RawMessage) *Parser {
	parsed := gjson.ParseBytes(data)
	return &Parser{result: parsed}
}

/**
 * Parse value on given path to provided type. 
 */
func ParseValue[T any](parser Parser, path string) (T, error) {
	var data T
	rawJson := parser.result.Get(path).Raw
	if rawJson == "" {
		return data, errors.New("parsing error: invalid data provided")
	}
	err := json.Unmarshal([]byte(rawJson), &data)
	if err != nil {
		log.Println("Parsing Error:", err.Error())
		return data, errors.New("Parsing Error: " + err.Error())
	}
	return data, nil
}
