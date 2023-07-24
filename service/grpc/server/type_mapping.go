package server

import "github.com/rew3/rew3-internal/service/grpc/api"

type RequestTypeMapping struct {
	Mapping map[api.APIOperation]interface{}
}

func NewRequestTypeMapping(mapping map[api.APIOperation]interface{}) *RequestTypeMapping {
	return &RequestTypeMapping{Mapping: mapping}
}
