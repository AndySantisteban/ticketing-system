package order_status_commands

type UpdateOrderStatusQuery struct {
	ID   int32
	Name string
}

type UpdateOrderStatusQueryResponse = error
