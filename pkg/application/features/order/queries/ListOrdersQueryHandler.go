package order_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type ListOrdersQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderRepository
}

func NewListOrdersQueryHandler(ctx context.Context, repository repositories.IOrderRepository) *ListOrdersQueryHandler {
	return &ListOrdersQueryHandler{
		ctx,
		repository,
	}
}

func (l *ListOrdersQueryHandler) Handler(args ListOrdersQuery) (ListOrdersQueryResponse, error) {

	list, err := l.repository.ListAllOrders(args.Offset, args.Limit)
	if err != nil {
		return nil, err
	}
	return list, err
}
