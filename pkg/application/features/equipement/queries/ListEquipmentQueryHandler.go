package equipment_features

import (
	"context"

	"InfositelOR/pkg/domain/repositories"
)

type ListEquipmentQueryHandler struct {
	_ctx        context.Context
	_repository repositories.IEquipementRepository
}

func NewListEquipmentQueryHandler(ctx context.Context, repository repositories.IEquipementRepository) *ListEquipmentQueryHandler {

	constructor := &ListEquipmentQueryHandler{
		_ctx:        ctx,
		_repository: repository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *ListEquipmentQueryHandler) Handler(param ListAllEquipmentQuery) (*ListAllEquipmentResponse, error) {

	response, err := r._repository.ListAllEquipement(param.Offset, param.Limit)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
