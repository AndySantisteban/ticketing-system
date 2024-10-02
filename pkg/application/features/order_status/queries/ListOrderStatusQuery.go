package order_status_features

import "InfositelOR/pkg/domain/entities"

type ListOrderStatusQuery struct {
	Limit  int32
	Offset int32
}

type ListOrderStatusQueryResponse = []entities.OrderStatus
