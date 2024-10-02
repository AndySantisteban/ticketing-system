package order_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type GetOrderQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderRepository
}

func NewGetOrderQueryHandler(ctx context.Context, repository repositories.IOrderRepository) *GetOrderQueryHandler {
	return &GetOrderQueryHandler{
		ctx,
		repository,
	}
}

func (g *GetOrderQueryHandler) Handler(args GetOrderQuery) (*GetOrderResponse, error) {
	// fmt.Println(args)
	item, err := g.repository.GetOrderByID(args.Id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
