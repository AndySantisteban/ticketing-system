package equipment_type_features

import (
	"context"

	"InfositelOR/pkg/domain/repositories"
)

type GetEquipmentTypeQueryHandler struct {
	_ctx        context.Context
	_repository repositories.IEquipmentTypeRepository
}

func NewGetEquipmentQueryHandler(ctx context.Context, repository repositories.IEquipmentTypeRepository) *GetEquipmentTypeQueryHandler {

	constructor := &GetEquipmentTypeQueryHandler{
		_ctx:        ctx,
		_repository: repository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *GetEquipmentTypeQueryHandler) Handler(param GetEquipmentTypeQuery) (*GetEquipmentTypeQueryResponse, error) {

	response, err := r._repository.GetEquipmentTypeByID(param.ID)

	if err != nil {
		return nil, err
	}

	return response, nil
}
