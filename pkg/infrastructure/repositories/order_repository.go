package repositories

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
	"context"

	"github.com/devfeel/mapper"
)

type OrderRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// CreateOrder implements repositories.IOrderRepository.
func (o *OrderRepository) CreateOrder(args entities.Order) (*entities.Order, error) {
	CreateActivity, err := o.queries.CreateOrder(o.ctx, persistence.CreateOrderParams{
		ClientID:      args.ClientID,
		EquipmentID:   args.EquipmentID,
		OrderNumber:   args.OrderNumber,
		ReportedIssue: args.ReportedIssue,
		Diagnosis:     args.Diagnosis,
		Solution:      args.Solution,
		EstimatedTime: args.EstimatedTime,
		Budget:        args.Budget,
		StatusID:      args.StatusID,
		AssignedTo:    args.AssignedTo,
		Priority:      args.Priority,
	})
	RespActivity := entities.Order{}
	mapper.AutoMapper(&CreateActivity, &RespActivity)

	return &RespActivity, err
}

// DeleteOrderByID implements repositories.IOrderRepository.
func (o *OrderRepository) DeleteOrderByID(id int32) error {
	err := o.queries.DeleteOrderByID(o.ctx, id)
	return err
}

// GetOrderByID implements repositories.IOrderRepository.
func (o *OrderRepository) GetOrderByID(id int32) (*entities.Order, error) {
	order, err := o.queries.GetOrderByID(o.ctx, id)
	if err != nil {
		return nil, err
	}
	ResOrder := &entities.Order{}
	mapper.AutoMapper(&order, ResOrder)
	return ResOrder, nil
}

// ListAllOrders implements repositories.IOrderRepository.
func (o *OrderRepository) ListAllOrders(offset int32, limit int32) ([]entities.Order, error) {
	data, err := o.queries.ListAllOrders(o.ctx, persistence.ListAllOrdersParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	Response := []entities.Order{}
	mapper.MapperSlice(&data, &Response)

	return Response, nil
}

// UpdateOrderByID implements repositories.IOrderRepository.
func (o *OrderRepository) UpdateOrderByID(args entities.Order) error {
	err := o.queries.UpdateOrderByID(o.ctx, persistence.UpdateOrderByIDParams{
		ID:            args.ID,
		ClientID:      args.ClientID,
		EquipmentID:   args.EquipmentID,
		OrderNumber:   args.OrderNumber,
		ReportedIssue: args.ReportedIssue,
		Diagnosis:     args.Diagnosis,
		Solution:      args.Solution,
		EstimatedTime: args.EstimatedTime,
		Budget:        args.Budget,
		StatusID:      args.StatusID,
		AssignedTo:    args.AssignedTo,
		Priority:      args.Priority,
	})
	return err
}

func NewOrderRepository(ctx context.Context, queries *persistence.Queries) repositories.IOrderRepository {
	return &OrderRepository{
		ctx:     ctx,
		queries: queries,
	}
}
