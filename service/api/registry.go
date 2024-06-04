package api

/**
 * Service API registry.
 * Contains list of all service methods. You can add methods as required/needed.
 */
type APIRegistry struct {
	registry map[Endpoint]*ServiceAPI
}

func NewAPIRegistry() *APIRegistry {
	registry := make(map[Endpoint]*ServiceAPI)
	return &APIRegistry{registry: registry}
}

func (registry *APIRegistry) AddServiceAPI(method *ServiceAPI) {
	registry.registry[method.api] = method
}

func (registry *APIRegistry) GetServiceAPI(api Endpoint) (bool, *ServiceAPI) {
	value, exists := registry.registry[api]
	return exists, value
}
