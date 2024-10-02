package equipment_type_commands

import (
	"context"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
)

type CreateEquipmentTypeQueryHandler struct {
	ctx        context.Context
	repository repositories.IEquipmentTypeRepository
}

func NewCreateEquipmentTypeQueryHandler(ctx context.Context, repository repositories.IEquipmentTypeRepository) *CreateEquipmentTypeQueryHandler {
	return &CreateEquipmentTypeQueryHandler{
		ctx,
		repository,
	}
}

func (c *CreateEquipmentTypeQueryHandler) Handler(args CreateEquipmentTypeQuery) (*CreateEquipmentTypeQueryResponse, error) {
	item, error := c.repository.CreateEquipmentType(entities.EquipmentType{
		ID:   args.ID,
		Name: args.Name,
	})

	if error != nil {
		return nil, error
	}

	return item, nil
}
