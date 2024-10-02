package comment_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type ListCommentsByOrderIDQueryHandler struct {
	ctx        context.Context
	repository repositories.ICommentRepository
}

func NewListCommentsByOrderIDQueryHandler(ctx context.Context, clientRepository repositories.ICommentRepository) *ListCommentsByOrderIDQueryHandler {

	constructor := &ListCommentsByOrderIDQueryHandler{
		ctx:        ctx,
		repository: clientRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *ListCommentsByOrderIDQueryHandler) Handler(param ListCommentsByOrderIDQuery) (*ListCommentsByOrderIDQueryResponse, error) {

	response, err := r.repository.ListCommentsByOrderID(param.Offset, param.Limit, &param.OrderID)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
