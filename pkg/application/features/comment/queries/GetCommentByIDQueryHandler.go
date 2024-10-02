package comment_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type GetCommentByIDQueryHandler struct {
	ctx        context.Context
	repository repositories.ICommentRepository
}

func NewGetCommentByIDQueryHandler(ctx context.Context, clientRepository repositories.ICommentRepository) *GetCommentByIDQueryHandler {

	constructor := &GetCommentByIDQueryHandler{
		ctx:        ctx,
		repository: clientRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *GetCommentByIDQueryHandler) Handler(param GetCommentByIDQuery) (*GetCommentByIDQueryResponse, error) {
	id := param.Id
	response, err := r.repository.GetCommentByID(id)

	if err != nil {
		return nil, err
	}

	return response, nil
}
