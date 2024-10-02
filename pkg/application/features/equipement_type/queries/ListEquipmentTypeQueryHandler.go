package equipment_type_features

import (
	"context"

	"InfositelOR/pkg/domain/repositories"
)

type ListEquipmentTypeQueryHandler struct {
	_ctx        context.Context
	_repository repositories.IEquipmentTypeRepository
}

func NewListEquipmentTypeQueryHandler(ctx context.Context, repository repositories.IEquipmentTypeRepository) *ListEquipmentTypeQueryHandler {

	constructor := &ListEquipmentTypeQueryHandler{
		_ctx:        ctx,
		_repository: repository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *ListEquipmentTypeQueryHandler) Handler(param ListAllEquipmentTypeQuery) (*ListAllEquipmentTypeResponse, error) {

	response, err := r._repository.ListAllEquipmentType(param.Offset, param.Limit)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
