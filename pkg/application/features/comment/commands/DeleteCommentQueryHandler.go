package comment_commands

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type DeleteCommentQueryHandler struct {
	ctx        context.Context
	repository repositories.ICommentRepository
}

func NewDeleteCommentQueryHandler(ctx context.Context, repository repositories.ICommentRepository) *DeleteCommentQueryHandler {

	return &DeleteCommentQueryHandler{
		ctx,
		repository,
	}
}

func (d *DeleteCommentQueryHandler) Handler(args DeleteCommentQuery) DeleteCommentQueryResponse {

	err := d.repository.DeleteCommentByID(args.Id)

	return err
}
