package models

import "encoding/json"

type RequestParam struct {
	Limit         *int                `bson:"limit,omitempty"`
	Expand        *bool               `bson:"expand,omitempty"`
	LimitFields   *string             `bson:"limit_fields,omitempty"`
	Offset        *int                `bson:"offset,omitempty"`
	Sort          *json.RawMessage    `bson:"sort,omitempty"`    //We can use *map[string]interface{} as well
	Filters       *json.RawMessage    `bson:"filters,omitempty"` //*map[string]interface{}
	RawParameters map[string][]string //This is not saved in DB. Rather used in following methods
}

func (r *RequestParam) WithRawParameters(params map[string][]string) *RequestParam {
	r.RawParameters = params
	return r
}

func (r *RequestParam) GetRawParameters() map[string][]string {
	return r.RawParameters
}
