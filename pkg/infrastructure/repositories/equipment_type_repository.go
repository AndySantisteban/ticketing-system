package repositories

import (
	"context"
	"errors"

	"github.com/devfeel/mapper"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
)

type EquipmentTypeRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// CreateEquipmentType implements repositories.IEquipmentTypeRepository.
func (e *EquipmentTypeRepository) CreateEquipmentType(args entities.EquipmentType) (*entities.EquipmentType, error) {
	create, err := e.queries.CreateEquipmentType(e.ctx, args.Name)
	res := entities.EquipmentType{}
	mapper.AutoMapper(&create, &res)

	return &res, err
}

// DeleteEquipmentTypeByID implements repositories.IEquipmentTypeRepository.
func (e *EquipmentTypeRepository) DeleteEquipmentTypeByID(id int32) error {
	// err := e.queries.DeleteEquipmentByID(e.ctx, id)
	return errors.New("the equipments can't be delete")
}

// GetEquipmentTypeByID implements repositories.IEquipmentTypeRepository.
func (e *EquipmentTypeRepository) GetEquipmentTypeByID(id int32) (*entities.EquipmentType, error) {
	data, err := e.queries.GetEquipmentTypeByID(e.ctx, id)
	if err != nil {
		return nil, err
	}
	res := &entities.EquipmentType{}
	mapper.AutoMapper(&data, res)
	return res, nil
}

// ListAllEquipmentType implements repositories.IEquipmentTypeRepository.
func (e *EquipmentTypeRepository) ListAllEquipmentType(offset int32, limit int32) ([]entities.EquipmentType, error) {
	data, err := e.queries.ListAllEquipmentTypes(e.ctx, persistence.ListAllEquipmentTypesParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	res := []entities.EquipmentType{}
	mapper.MapperSlice(&data, &res)

	return res, nil
}

// UpdateEquipmentTypeByID implements repositories.IEquipmentTypeRepository.
func (e *EquipmentTypeRepository) UpdateEquipmentTypeByID(args entities.EquipmentType) error {
	err := e.queries.UpdateEquipmentTypeByID(e.ctx, persistence.UpdateEquipmentTypeByIDParams{
		ID:   args.ID,
		Name: args.Name,
	})
	return err
}

func NewEquipmentTypeRepository(ctx context.Context, queries *persistence.Queries) repositories.IEquipmentTypeRepository {
	return &EquipmentTypeRepository{
		ctx:     ctx,
		queries: queries,
	}
}
