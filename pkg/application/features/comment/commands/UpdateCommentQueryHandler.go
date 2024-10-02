package comment_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
	"database/sql"
)

type UpdateCommentQueryHandler struct {
	ctx        context.Context
	repository repositories.ICommentRepository
}

func NewUpdateCommentQueryHandler(ctx context.Context, repository repositories.ICommentRepository) *UpdateCommentQueryHandler {

	return &UpdateCommentQueryHandler{
		ctx,
		repository,
	}
}

func (u *UpdateCommentQueryHandler) Handler(args *UpdateCommentQuery) UpdateCommentQueryResponse {

	error := u.repository.UpdateCommentByID(entities.Comment{
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

	return error
}
