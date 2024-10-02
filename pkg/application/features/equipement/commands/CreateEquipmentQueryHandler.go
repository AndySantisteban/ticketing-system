package equipment_commands

import (
	"context"
	"database/sql"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
)

type CreateEquipmentQueryHandler struct {
	ctx        context.Context
	repository repositories.IEquipementRepository
}

func NewCreateEquipmentHandler(ctx context.Context, repository repositories.IEquipementRepository) *CreateEquipmentQueryHandler {
	return &CreateEquipmentQueryHandler{
		ctx,
		repository,
	}
}

func (c *CreateEquipmentQueryHandler) Handler(args CreateEquipmentQuery) (*CreateEquipmentQueryResponse, error) {
	item, error := c.repository.CreateEquipment(entities.Equipment{
		ID:   args.ID,
		Name: args.Name,
		TypeID: sql.NullInt32{
			Int32: *args.TypeID,
			Valid: true,
		},
		SerialNumber: args.SerialNumber,
		Notes: sql.NullString{
			String: *args.Notes,
			Valid:  true,
		},
	})

	if error != nil {
		return nil, error
	}

	return item, nil
}
