package equipment_type_commands

import (
	"context"
	"fmt"

	"InfositelOR/pkg/domain/repositories"
)

type DeleteEquipmentTypeQueryHandler struct {
	ctx        context.Context
	repository repositories.IEquipmentTypeRepository
}

func NewDeleteEquipementTypeQueryHandler(ctx context.Context, repository repositories.IEquipmentTypeRepository) *DeleteEquipmentTypeQueryHandler {

	return &DeleteEquipmentTypeQueryHandler{
		ctx,
		repository,
	}
}

func (d *DeleteEquipmentTypeQueryHandler) Handler(args DeleteEquipmentTypeQuery) DeleteEquipmentTypeQueryResponse {

	fmt.Println(args)
	err := d.repository.DeleteEquipmentTypeByID(args.Id)

	return err
}
