package request

type RequestParam struct {
	Limit       int    `json:"limit,omitempty" bson:"limit,omitempty"`
	Expand      bool   `json:"expand,omitempty" bson:"expand,omitempty"`
	LimitFields string `json:"limit_fields,omitempty" bson:"limit_fields,omitempty"`
	Offset      int    `json:"offset,omitempty" bson:"offset,omitempty"`
	Sort        string `json:"sort,omitempty" bson:"sort,omitempty"`       //We can use *map[string]interface{} as well
	Filters     string `json:"filters,omitempty" bson:"filters,omitempty"` //*map[string]interface{}
}
