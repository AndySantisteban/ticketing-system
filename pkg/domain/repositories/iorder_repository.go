package repositories

import (
	"InfositelOR/pkg/domain/entities"
)

type IOrderRepository interface {
	ListAllOrders(offset int32, limit int32) ([]entities.Order, error)
	GetOrderByID(id int32) (*entities.Order, error)
	CreateOrder(args entities.Order) (*entities.Order, error)
	UpdateOrderByID(args entities.Order) error
	DeleteOrderByID(id int32) error
}
