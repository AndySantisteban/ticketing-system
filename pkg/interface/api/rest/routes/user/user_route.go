package user

import (
	user_commands "InfositelOR/pkg/application/features/user/commands"
	user_features "InfositelOR/pkg/application/features/user/queries"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	getID  *user_features.GetUserByIDQueryHandler
	list   *user_features.ListUserQueryHandler
	create *user_commands.CreateUserQueryHandler
	// update *order_status_commands.UpdateOrderStatusQueryHandler
	// delete *comment_commands.DeleteCommentQueryHandler
}

func RestUserRouter(
	getID *user_features.GetUserByIDQueryHandler,
	list *user_features.ListUserQueryHandler,
	create *user_commands.CreateUserQueryHandler,
	// update *order_status_commands.UpdateOrderStatusQueryHandler,

) *UserRouter {

	return &UserRouter{
		getID,
		list,
		create,
		// update,
		// delete,
	}
}

func (ar *UserRouter) ListUsersRoute(c echo.Context) error {
	query := new(ListUsersRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.list.Handler(user_features.ListUserQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *UserRouter) GetUserRouteByUid(c echo.Context) error {
	query := new(GetUserByUidRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	item, _ := strconv.Atoi(c.QueryParam("id"))

	response, err := ar.getID.Handler(user_features.GetUserByIDQuery{
		ID: int32(item),
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *UserRouter) CreateUserRoute(c echo.Context) error {
	query := new(CreateUserDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.create.Handler(user_commands.CreateUserQuery{
		ID:             query.ID,
		Name:           query.Name,
		Email:          query.Email,
		PermissionType: query.PermissionType,
		CreationDate:   query.CreationDate,
		InactiveStatus: query.InactiveStatus,
		Password:       query.Password,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}
