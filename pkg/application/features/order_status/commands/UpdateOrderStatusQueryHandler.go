package order_status_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type UpdateOrderStatusQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderStatusRepository
}

func NewUpdateOrderStatusQueryHandler(ctx context.Context, repository repositories.IOrderStatusRepository) *UpdateOrderStatusQueryHandler {

	return &UpdateOrderStatusQueryHandler{
		ctx,
		repository,
	}
}

func (c *UpdateOrderStatusQueryHandler) Handler(args UpdateOrderStatusQuery) UpdateOrderStatusQueryResponse {

	err := c.repository.UpdateOrderStatusByID(entities.OrderStatus{
		ID:   args.ID,
		Name: args.Name,
	})

	return err
}
