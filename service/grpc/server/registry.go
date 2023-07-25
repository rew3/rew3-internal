package server

import "github.com/rew3/rew3-internal/service/grpc/api"

/**
 * Service method registry. 
 * Contains list of all service methods. You can add methods as required/needed. 
 */
type ServiceMethodRegistry struct {
	registry map[api.APIOperation]ServiceMethod
}

func NewServiceMethodRegistry() *ServiceMethodRegistry {
	registry := make(map[api.APIOperation]ServiceMethod)
	return &ServiceMethodRegistry{registry: registry}
}

func (registry *ServiceMethodRegistry) AddServiceMethod(method ServiceMethod) {
	registry.registry[method.api] = method
}

func (registry *ServiceMethodRegistry) GetServiceMethod(api api.APIOperation) (bool, *ServiceMethod) {
	value, exists := registry.registry[api]
	return exists, &value
}
