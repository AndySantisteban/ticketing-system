package order

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	order_commands "InfositelOR/pkg/application/features/order/commands"
	order_features "InfositelOR/pkg/application/features/order/queries"
)

type OrderRouter struct {
	getID  *order_features.GetOrderQueryHandler
	list   *order_features.ListOrdersQueryHandler
	create *order_commands.CreateOrderQueryHandler
	update *order_commands.UpdateOrderQueryHandler
	// delete *comment_commands.DeleteCommentQueryHandler
}

func RestOrderRouter(
	getID *order_features.GetOrderQueryHandler,
	list *order_features.ListOrdersQueryHandler,
	create *order_commands.CreateOrderQueryHandler,
	update *order_commands.UpdateOrderQueryHandler,

) *OrderRouter {

	return &OrderRouter{
		getID,
		list,
		create,
		update,
		// delete,
	}
}

func (ar *OrderRouter) ListOrdersRoute(c echo.Context) error {
	query := new(ListOrdersRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.list.Handler(order_features.ListOrdersQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *OrderRouter) GetOrderRouteByUid(c echo.Context) error {
	query := new(GetOrdersByUidRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	item, _ := strconv.Atoi(c.QueryParam("id"))

	response, err := ar.getID.Handler(order_features.GetOrderQuery{
		Id: int32(item),
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *OrderRouter) GetOrderWithTotalnfo(c echo.Context) error {
	query := new(GetOrdersByUidRouteDTO)
	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	item, _ := strconv.Atoi(c.QueryParam("id"))

	response, err := ar.getID.Handler(order_features.GetOrderQuery{
		Id: int32(item),
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)

}

func (ar *OrderRouter) CreateRouteByUid(c echo.Context) error {
	query := new(CreateRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(query)

	item, err := ar.create.Handler(order_commands.CreateOrderQuery{
		ID:            query.Id,
		ClientID:      query.ClientID,
		Equipement:    query.Equipement,
		OrderNumber:   query.OrderNumber,
		ReportedIssue: query.ReportedIssue,
		Diagnosis:     query.Diagnosis,
		Solution:      query.Solution,
		EstimatedTime: query.EstimatedTime,
		Budget:        query.Budget,
		StatusID:      query.StatusID,
		AssignedTo:    query.AssignedTo,
		CreationDate:  query.CreationDate,
		Priority:      query.Priority,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}
	return c.JSON(200, item)
}

func (ar *OrderRouter) UpdateRouteByUid(c echo.Context) error {
	query := new(UpdateRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	err := ar.update.Handler(order_commands.UpdateOrderQuery{
		ID:            query.Id,
		ClientID:      query.ClientID,
		Equipement:    query.Equipement,
		OrderNumber:   query.OrderNumber,
		ReportedIssue: query.ReportedIssue,
		Diagnosis:     query.Diagnosis,
		Solution:      query.Solution,
		EstimatedTime: query.EstimatedTime,
		Budget:        query.Budget,
		StatusID:      query.StatusID,
		AssignedTo:    query.AssignedTo,
		CreationDate:  query.CreationDate,
		Priority:      query.Priority,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}
	return c.JSON(200, nil)
}
