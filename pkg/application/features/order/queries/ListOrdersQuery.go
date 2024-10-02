package order_features

import "InfositelOR/pkg/domain/entities"

type ListOrdersQuery struct {
	Offset int32
	Limit  int32
}

type ListOrdersQueryResponse = []entities.Order
