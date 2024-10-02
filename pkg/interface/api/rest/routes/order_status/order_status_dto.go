package orderstatus

type ListOrdersStatusRouteDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

type GetOrderStatusByUidRouteDTO struct {
	ID int32 `param:"ID"`
}

type CreateOrderStatusRouteDTO struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type UpdateOrderStatusRouteDTO struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
