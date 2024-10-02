package client

import (
	client_commands "InfositelOR/pkg/application/features/client/commands"
	client_features "InfositelOR/pkg/application/features/client/queries"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClientRouter struct {
	getID  *client_features.GetClientByIDQueryHandler
	list   *client_features.ListClientsQueryHandler
	create *client_commands.CreateClientQueryHandler
	update *client_commands.UpdateClientQueryHandler
	delete *client_commands.DeleteClientByIDQueryHandler
}

func RestClientRouter(getID *client_features.GetClientByIDQueryHandler, list *client_features.ListClientsQueryHandler, create *client_commands.CreateClientQueryHandler, update *client_commands.UpdateClientQueryHandler, delete *client_commands.DeleteClientByIDQueryHandler) *ClientRouter {

	return &ClientRouter{
		getID,
		list,
		create,
		update,
		delete,
	}
}

func (ar *ClientRouter) ListClientRoute(c echo.Context) error {
	query := new(ListClientRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.list.Handler(client_features.ListAllClientsQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *ClientRouter) GetClientByUid(c echo.Context) error {
	query := new(GetClientByUidRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	item, _ := strconv.Atoi(c.QueryParam("id"))
	response, err := ar.getID.Handler(client_features.GetClientByIDQuery{
		Id: int32(item),
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *ClientRouter) CreateClient(c echo.Context) error {
	query := new(CreateClientRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.create.Handler(client_commands.CreateClientQuery{
		ID:            query.ID,
		Name:          query.Name,
		Address:       query.Address,
		District:      query.District,
		City:          query.City,
		Country:       query.Country,
		Phone:         query.Phone,
		Ruc:           query.Ruc,
		ContactPerson: query.ContactPerson,
		Email:         query.Email,
		Website:       query.Website,
		AddressLine2:  query.AddressLine2,
		PostalCode:    query.PostalCode,
		Fax:           query.Fax,
		Notes:         query.Notes,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *ClientRouter) UpdateClient(c echo.Context) error {
	query := new(UpdateClientRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	err := ar.update.Handler(client_commands.UpdateClientQuery{
		ID:            query.ID,
		Name:          query.Name,
		Address:       query.Address,
		District:      query.District,
		City:          query.City,
		Country:       query.Country,
		Phone:         query.Phone,
		Ruc:           query.Ruc,
		ContactPerson: query.ContactPerson,
		Email:         query.Email,
		Website:       query.Website,
		AddressLine2:  query.AddressLine2,
		PostalCode:    query.PostalCode,
		Fax:           query.Fax,
		Notes:         query.Notes,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, err)
}

func (ar *ClientRouter) DeleteClient(c echo.Context) error {
	query := new(DeleteClientRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	err := ar.delete.Handler(client_commands.DeleteClientByIDQuery{
		Id: query.ID,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, err)
}
