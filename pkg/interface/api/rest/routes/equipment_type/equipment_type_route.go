package equipment_type

import (
	"net/http"

	"github.com/labstack/echo/v4"

	equipment_type_commands "InfositelOR/pkg/application/features/equipement_type/commands"
	equipment_type_features "InfositelOR/pkg/application/features/equipement_type/queries"
)

type EquipmentTypeRouter struct {
	getID  *equipment_type_features.GetEquipmentTypeQueryHandler
	list   *equipment_type_features.ListEquipmentTypeQueryHandler
	create *equipment_type_commands.CreateEquipmentTypeQueryHandler
	// update *comment_commands.UpdateCommentQueryHandler
	delete *equipment_type_commands.DeleteEquipmentTypeQueryHandler
}

func RestEquipmentRouter(
	getID *equipment_type_features.GetEquipmentTypeQueryHandler,
	list *equipment_type_features.ListEquipmentTypeQueryHandler,
	create *equipment_type_commands.CreateEquipmentTypeQueryHandler,
	// update *comment_commands.UpdateCommentQueryHandler,
	delete *equipment_type_commands.DeleteEquipmentTypeQueryHandler,

) *EquipmentTypeRouter {

	return &EquipmentTypeRouter{
		getID,
		list,
		create,
		// update,
		delete,
	}
}

func (ar *EquipmentTypeRouter) List(c echo.Context) error {
	query := new(ListAllEquipmentTypeQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.list.Handler(equipment_type_features.ListAllEquipmentTypeQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *EquipmentTypeRouter) GetByID(c echo.Context) error {
	query := new(GetEquipmentTypeQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.getID.Handler(equipment_type_features.GetEquipmentTypeQuery{
		ID: query.ID,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *EquipmentTypeRouter) Create(c echo.Context) error {
	query := new(CreateEquipmentTypeQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.create.Handler(equipment_type_commands.CreateEquipmentTypeQuery{
		ID:   query.ID,
		Name: query.Name,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *EquipmentTypeRouter) Delete(c echo.Context) error {
	query := new(DeleteEquipmentTypeRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := ar.delete.Handler(equipment_type_commands.DeleteEquipmentTypeQuery{
		Id: query.ID,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, nil)
}
