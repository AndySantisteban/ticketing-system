package order_status_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type CreateOrderStatusQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderStatusRepository
}

func NewCreateOrderStatusQueryHandler(ctx context.Context, repository repositories.IOrderStatusRepository) *CreateOrderStatusQueryHandler {
	return &CreateOrderStatusQueryHandler{
		ctx,
		repository,
	}
}

func (c *CreateOrderStatusQueryHandler) Handler(args CreateOrderStatusQuery) (*CreateOrderStatusQueryResponse, error) {
	item, err := c.repository.CreateOrderStatus(entities.OrderStatus{
		ID:   args.ID,
		Name: args.Name,
	})

	if err != nil {
		return nil, err
	}

	return item, nil
}
