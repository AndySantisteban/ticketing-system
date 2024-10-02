package user_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type GetUserByIDQueryHandler struct {
	ctx          context.Context
	repositories repositories.IUserRepository
}

func NewGetUserByIDQueryHandler(
	ctx context.Context,
	repositories repositories.IUserRepository,
) *GetUserByIDQueryHandler {
	return &GetUserByIDQueryHandler{
		ctx,
		repositories,
	}
}

func (g *GetUserByIDQueryHandler) Handler(args GetUserByIDQuery) (*GetUserByIDQueryResponse, error) {

	item, err := g.repositories.GetUserByID(args.ID)

	if err != nil {
		return nil, err
	}

	return item, nil
}
