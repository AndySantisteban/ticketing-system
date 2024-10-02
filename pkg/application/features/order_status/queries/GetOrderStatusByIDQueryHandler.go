package order_status_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type GetOrderStatusByIDQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderStatusRepository
}

func NewGetOrderStatusByIDQueryHandler(ctx context.Context, repository repositories.IOrderStatusRepository) *GetOrderStatusByIDQueryHandler {

	return &GetOrderStatusByIDQueryHandler{
		ctx,
		repository,
	}
}

func (g *GetOrderStatusByIDQueryHandler) Handler(args GetOrderStatusByIDQuery) (*GetOrderStatusByIDQueryResponse, error) {

	response, err := g.repository.GetOrderStatusByID(args.ID)
	if err != nil {
		return nil, err
	}
	return response, nil
}
