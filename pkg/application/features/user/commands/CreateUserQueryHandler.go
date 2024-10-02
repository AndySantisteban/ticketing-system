package user_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
	"database/sql"
)

type CreateUserQueryHandler struct {
	ctx        context.Context
	repository repositories.IUserRepository
}

func NewCreateUserQueryHandler(
	ctx context.Context,
	repository repositories.IUserRepository,
) *CreateUserQueryHandler {

	return &CreateUserQueryHandler{
		ctx,
		repository,
	}
}

func (h *CreateUserQueryHandler) Handler(args CreateUserQuery) (*CreateUserQueryResponse, error) {

	item, err := h.repository.CreateUser(
		entities.User{
			ID:             args.ID,
			Name:           args.Name,
			Email:          args.Email,
			PermissionType: args.PermissionType,
			Password: sql.NullString{
				Valid:  true,
				String: args.Password,
			},
			CreationDate: sql.NullTime{
				Valid: true,
				Time:  *args.CreationDate,
			},
			InactiveStatus: sql.NullString{
				Valid:  true,
				String: *args.InactiveStatus,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return item, err
}
