package equipment

import (
	"net/http"

	"github.com/labstack/echo/v4"

	equipment_commands "InfositelOR/pkg/application/features/equipement/commands"
	equipment_features "InfositelOR/pkg/application/features/equipement/queries"
)

type EquipmentRouter struct {
	getID  *equipment_features.GetEquipmentQueryHandler
	list   *equipment_features.ListEquipmentQueryHandler
	create *equipment_commands.CreateEquipmentQueryHandler
	// update *comment_commands.UpdateCommentQueryHandler
	delete *equipment_commands.DeleteEquipmentQueryHandler
}

func RestEquipmentRouter(
	getID *equipment_features.GetEquipmentQueryHandler,
	list *equipment_features.ListEquipmentQueryHandler,
	create *equipment_commands.CreateEquipmentQueryHandler,
	// update *comment_commands.UpdateCommentQueryHandler,
	delete *equipment_commands.DeleteEquipmentQueryHandler,

) *EquipmentRouter {

	return &EquipmentRouter{
		getID,
		list,
		create,
		// update,
		delete,
	}
}

func (ar *EquipmentRouter) List(c echo.Context) error {
	query := new(ListAllEquipmentQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	response, err := ar.list.Handler(equipment_features.ListAllEquipmentQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
	})

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *EquipmentRouter) GetByID(c echo.Context) error {
	query := new(GetEquipmentQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.getID.Handler(equipment_features.GetEquipmentQuery{
		ID: query.ID,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *EquipmentRouter) Create(c echo.Context) error {
	query := new(CreateEquipmentQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := ar.create.Handler(equipment_commands.CreateEquipmentQuery{
		ID:           query.ID,
		Name:         query.Name,
		TypeID:       query.TypeID,
		Notes:        query.Notes,
		SerialNumber: query.SerialNumber,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, response)
}

func (ar *EquipmentRouter) Delete(c echo.Context) error {
	query := new(DeleteEquipmentTypeQueryDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err := ar.delete.Handler(equipment_commands.DeleteEquipmentQuery{
		Id: query.Id,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, nil)
}
