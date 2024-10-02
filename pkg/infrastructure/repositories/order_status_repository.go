package repositories

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
	"context"

	"github.com/devfeel/mapper"
)

type OrderStatusRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// CreateOrderStatus implements repositories.IOrderStatusRepository.
func (o *OrderStatusRepository) CreateOrderStatus(args entities.OrderStatus) (*entities.OrderStatus, error) {
	CreateOrderStatus, err := o.queries.CreateOrderStatus(o.ctx, args.Name)
	RespOrderstatus := entities.OrderStatus{}
	mapper.AutoMapper(&CreateOrderStatus, &RespOrderstatus)

	return &RespOrderstatus, err
}

// GetOrderStatusByID implements repositories.IOrderStatusRepository.
func (o *OrderStatusRepository) GetOrderStatusByID(id int32) (*entities.OrderStatus, error) {
	item, err := o.queries.GetOrderStatusByID(o.ctx, id)
	if err != nil {
		return nil, err
	}
	RespOrderstatus := entities.OrderStatus{}
	mapper.AutoMapper(&item, &RespOrderstatus)

	return &RespOrderstatus, nil
}

// ListAllOrderStatus implements repositories.IOrderStatusRepository.
func (o *OrderStatusRepository) ListAllOrderStatus(offset int32, limit int32) ([]entities.OrderStatus, error) {
	data, err := o.queries.ListAllOrderStatus(o.ctx, persistence.ListAllOrderStatusParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	Response := []entities.OrderStatus{}
	mapper.MapperSlice(&data, &Response)

	return Response, nil
}

// UpdateOrderStatusByID implements repositories.IOrderStatusRepository.
func (o *OrderStatusRepository) UpdateOrderStatusByID(args entities.OrderStatus) error {
	err := o.queries.UpdateOrderStatusByID(o.ctx, persistence.UpdateOrderStatusByIDParams{
		ID:   args.ID,
		Name: args.Name,
	})

	return err
}

func NewOrderStatusRepository(ctx context.Context, queries *persistence.Queries) repositories.IOrderStatusRepository {
	return &OrderStatusRepository{
		ctx:     ctx,
		queries: queries,
	}
}
