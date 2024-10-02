package comment

import (
	comment_commands "InfositelOR/pkg/application/features/comment/commands"
	comment_features "InfositelOR/pkg/application/features/comment/queries"
	"InfositelOR/pkg/interface/api/rest/routes/login"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentRouter struct {
	getID  *comment_features.GetCommentByIDQueryHandler
	list   *comment_features.ListCommentsByOrderIDQueryHandler
	create *comment_commands.CreateCommentQueryHandler
	update *comment_commands.UpdateCommentQueryHandler
	delete *comment_commands.DeleteCommentQueryHandler
}

func RestCommentRouter(getID *comment_features.GetCommentByIDQueryHandler, list *comment_features.ListCommentsByOrderIDQueryHandler, create *comment_commands.CreateCommentQueryHandler, update *comment_commands.UpdateCommentQueryHandler, delete *comment_commands.DeleteCommentQueryHandler) *CommentRouter {

	return &CommentRouter{
		getID,
		list,
		create,
		update,
		delete,
	}
}

func (ar *CommentRouter) ListCommentRoute(c echo.Context) error {
	query := new(ListCommentRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.list.Handler(comment_features.ListCommentsByOrderIDQuery{
		Limit:   query.Limit,
		Offset:  query.Offset,
		OrderID: query.OrderID,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *CommentRouter) GetCommentRouteByUid(c echo.Context) error {
	query := new(GetCommentByUidRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.getID.Handler(comment_features.GetCommentByIDQuery{
		Id: query.Id,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *CommentRouter) CreateComment(c echo.Context) error {
	query := new(CreateCommentRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	tokenString := c.Request().Header.Get("Authorization")

	tokenString = tokenString[len("Bearer "):]
	item, _ := login.GetUserInfoFromToken(tokenString)

	ID := item["ID"].(float64)

	response, err := ar.create.Handler(comment_commands.CreateCommentQuery{
		ID:      query.ID,
		OrderID: query.OrderID,
		UserID:  int32(ID),
		Date:    query.Date,
		Comment: query.Comment,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *CommentRouter) UpdateComment(c echo.Context) error {
	query := new(UpdateCommentRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := ar.update.Handler(&comment_commands.UpdateCommentQuery{
		ID:      query.ID,
		OrderID: query.OrderID,
		UserID:  query.UserID,
		Date:    query.Date,
		Comment: query.Comment,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, nil)
}

func (ar *CommentRouter) DeleteComment(c echo.Context) error {
	query := new(DeleteCommentRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := ar.delete.Handler(comment_commands.DeleteCommentQuery{
		Id: query.ID,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, nil)
}
