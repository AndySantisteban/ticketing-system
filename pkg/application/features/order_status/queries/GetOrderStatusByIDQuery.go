package order_status_features

import "InfositelOR/pkg/domain/entities"

type GetOrderStatusByIDQuery struct {
	ID int32
}

type GetOrderStatusByIDQueryResponse = entities.OrderStatus
