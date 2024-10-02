package equipment_commands

import (
	"context"
	"fmt"

	"InfositelOR/pkg/domain/repositories"
)

type DeleteEquipmentQueryHandler struct {
	ctx        context.Context
	repository repositories.IEquipementRepository
}

func NewDeleteEquipmentQueryHandler(ctx context.Context, repository repositories.IEquipementRepository) *DeleteEquipmentQueryHandler {

	return &DeleteEquipmentQueryHandler{
		ctx,
		repository,
	}
}

func (d *DeleteEquipmentQueryHandler) Handler(args DeleteEquipmentQuery) DeleteEquipmentQueryResponse {
	fmt.Println(args.Id)
	err := d.repository.DeleteEquipmentByID(args.Id)

	return err
}
