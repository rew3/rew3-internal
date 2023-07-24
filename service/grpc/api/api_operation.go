package api

type APIOperation string

// Resolve API Operation from String.
func ResolveOperation(api string) APIOperation {
	return APIOperation(api)
}
