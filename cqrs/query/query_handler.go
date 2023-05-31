package query

type QueryHandler interface {
	Handle(query Query)
}
