package response

type Response[T any] struct {
	Action string
}

func NewResponse[T any](action string) Response[T] {
	return Response[T]{
		Action: action,
	}
}