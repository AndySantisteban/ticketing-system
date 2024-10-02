package order_status_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type ListOrderStatusQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderStatusRepository
}

func NewListOrderStatusQueryHandler(ctx context.Context, repository repositories.IOrderStatusRepository) *ListOrderStatusQueryHandler {

	return &ListOrderStatusQueryHandler{
		ctx,
		repository,
	}
}

func (l *ListOrderStatusQueryHandler) Handler(args ListOrderStatusQuery) (ListOrderStatusQueryResponse, error) {

	list, err := l.repository.ListAllOrderStatus(args.Offset, args.Limit)

	if err != nil {
		return nil, err
	}

	return list, nil
}
