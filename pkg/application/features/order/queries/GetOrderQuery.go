package order_features

import "InfositelOR/pkg/domain/entities"

type GetOrderQuery struct {
	Id int32
}

type GetOrderResponse = entities.Order
