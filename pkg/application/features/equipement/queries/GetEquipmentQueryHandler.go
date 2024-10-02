package equipment_features

import (
	"context"

	"InfositelOR/pkg/domain/repositories"
)

type GetEquipmentQueryHandler struct {
	_ctx        context.Context
	_repository repositories.IEquipementRepository
}

func NewGetEquipmentQueryHandler(ctx context.Context, repository repositories.IEquipementRepository) *GetEquipmentQueryHandler {

	constructor := &GetEquipmentQueryHandler{
		_ctx:        ctx,
		_repository: repository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *GetEquipmentQueryHandler) Handler(param GetEquipmentQuery) (*GetEquipmentQueryResponse, error) {

	response, err := r._repository.GetEquipmentByID(param.ID)

	if err != nil {
		return nil, err
	}

	return response, nil
}
