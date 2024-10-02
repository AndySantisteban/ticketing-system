package order_status_commands

import "InfositelOR/pkg/domain/entities"

type CreateOrderStatusQuery struct {
	ID   int32
	Name string
}

type CreateOrderStatusQueryResponse = entities.OrderStatus
