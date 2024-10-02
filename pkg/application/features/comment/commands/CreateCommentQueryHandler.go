package comment_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
	"database/sql"
)

type CreateCommentQueryHandler struct {
	ctx        context.Context
	repository repositories.ICommentRepository
}

func NewCreateCommentQueryHandler(ctx context.Context, repository repositories.ICommentRepository) *CreateCommentQueryHandler {
	return &CreateCommentQueryHandler{
		ctx,
		repository,
	}
}

func (c *CreateCommentQueryHandler) Handler(args CreateCommentQuery) (*CreateCommentQueryResponse, error) {
	item, error := c.repository.CreateComment(entities.Comment{
		ID: args.ID,
		OrderID: sql.NullInt32{
			Valid: true,
			Int32: args.OrderID,
		},
		UserID: sql.NullInt32{
			Valid: true,
			Int32: args.UserID,
		},
		Date: sql.NullTime{
			Valid: true,
			Time:  *args.Date,
		},
		Comment: sql.NullString{
			Valid:  true,
			String: *args.Comment,
		},
	})

	if error != nil {
		return nil, error
	}

	return item, nil
}
