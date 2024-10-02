package user_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type ListUserQueryHandler struct {
	ctx        context.Context
	repository repositories.IUserRepository
}

func NewListUserQueryHanlder(
	ctx context.Context,
	repository repositories.IUserRepository,
) *ListUserQueryHandler {

	return &ListUserQueryHandler{
		ctx,
		repository,
	}
}

func (l *ListUserQueryHandler) Handler(args ListUserQuery) (ListUserQueryResponse, error) {

	data, err := l.repository.ListAllUser(args.Offset, args.Limit)
	if err != nil {
		return nil, err
	}
	return data, nil

}
