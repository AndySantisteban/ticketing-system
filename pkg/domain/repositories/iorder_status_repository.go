package repositories

import (
	"InfositelOR/pkg/domain/entities"
)

type IOrderStatusRepository interface {
	ListAllOrderStatus(offset int32, limit int32) ([]entities.OrderStatus, error)
	GetOrderStatusByID(id int32) (*entities.OrderStatus, error)
	CreateOrderStatus(args entities.OrderStatus) (*entities.OrderStatus, error)
	UpdateOrderStatusByID(args entities.OrderStatus) error
}
