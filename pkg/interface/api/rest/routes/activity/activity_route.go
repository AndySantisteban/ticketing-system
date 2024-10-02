package activity

import (
	Update "InfositelOR/pkg/application/features/activity/commands/CreateActivity"
	Get "InfositelOR/pkg/application/features/activity/queries/GetActivityByUid"
	GetList "InfositelOR/pkg/application/features/activity/queries/ListActivityByOrderID"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityRouter struct {
	GetList *GetList.ListActivityByOrderIDQueryHandler
	Get     *Get.GetActivityByUidQueryHandler
	Update  *Update.CreateActivityQueryHandler
}

func RestActivityRouter(GetList *GetList.ListActivityByOrderIDQueryHandler, Get *Get.GetActivityByUidQueryHandler, Update *Update.CreateActivityQueryHandler) *ActivityRouter {

	router := ActivityRouter{
		GetList: GetList,
		Get:     Get,
		Update:  Update,
	}

	return &router
}

func (ar *ActivityRouter) ListActivityByOrderIDRoute(c echo.Context) error {
	query := new(ListActivityByOrderIDRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.GetList.Handler(GetList.ListActivityByOrderIDQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
		Id:     query.Id,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *ActivityRouter) GetActivityByUid(c echo.Context) error {
	query := new(GetActivityByUidRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.Get.Handler(Get.GetActivityByUidQuery{
		Id: query.Id,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *ActivityRouter) CreateActivity(c echo.Context) error {

	query := new(CreateActivityQuery)
	if err := c.Bind(query); err != nil {
		return c.String(http.StatusInternalServerError, "Bad request")
	}

	response, err := ar.Update.Handler(Update.CreateActivityQuery{
		ID:      query.ID,
		OrderID: query.OrderID,
		UserID:  query.UserID,
		Date:    query.Date,
		Action:  query.Action,
		Details: query.Details,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}
