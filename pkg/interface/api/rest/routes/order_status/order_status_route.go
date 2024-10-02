package orderstatus

import (
	order_status_commands "InfositelOR/pkg/application/features/order_status/commands"
	order_status_features "InfositelOR/pkg/application/features/order_status/queries"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderStatusRouter struct {
	getID  *order_status_features.GetOrderStatusByIDQueryHandler
	list   *order_status_features.ListOrderStatusQueryHandler
	create *order_status_commands.CreateOrderStatusQueryHandler
	update *order_status_commands.UpdateOrderStatusQueryHandler
	// delete *comment_commands.DeleteCommentQueryHandler
}

func RestOrderStatusRouter(
	getID *order_status_features.GetOrderStatusByIDQueryHandler,
	list *order_status_features.ListOrderStatusQueryHandler,
	create *order_status_commands.CreateOrderStatusQueryHandler,
	update *order_status_commands.UpdateOrderStatusQueryHandler,

) *OrderStatusRouter {

	return &OrderStatusRouter{
		getID,
		list,
		create,
		update,
		// delete,
	}
}

func (ar *OrderStatusRouter) ListOrdersStatusRoute(c echo.Context) error {
	query := new(ListOrdersStatusRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	response, err := ar.list.Handler(order_status_features.ListOrderStatusQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *OrderStatusRouter) GetOrderStatusRouteByUid(c echo.Context) error {
	query := new(GetOrderStatusByUidRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	item, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, "ID not Found")
	}
	response, err := ar.getID.Handler(order_status_features.GetOrderStatusByIDQuery{
		ID: int32(item),
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *OrderStatusRouter) CreateOrderStatusRouteByUid(c echo.Context) error {
	query := new(UpdateOrderStatusRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	item, err := ar.create.Handler(order_status_commands.CreateOrderStatusQuery{
		ID:   query.Id,
		Name: query.Name,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}
	return c.JSON(200, item)
}

func (ar *OrderStatusRouter) UpdateOrderStatusRouteByUid(c echo.Context) error {
	query := new(UpdateOrderStatusRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	err := ar.update.Handler(order_status_commands.UpdateOrderStatusQuery{
		ID:   query.Id,
		Name: query.Name,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}
	return c.JSON(200, nil)
}
