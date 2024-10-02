package cqrs

type ImplementsHandler[TParams any, QueryResponse any] interface {
	Handler(param TParams) QueryResponse
}
