package api

type Endpoint string

// Resolve API Endpoint from String.
func ResolveEndpoint(api string) Endpoint {
	return Endpoint(api)
}
