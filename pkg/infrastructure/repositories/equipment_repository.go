package repositories

import (
	"context"

	"github.com/devfeel/mapper"
	// "github.com/devfeel/mapper"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
)

type EquipmentRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// CreateEquipment implements repositories.IEquipementRepository.
func (e *EquipmentRepository) CreateEquipment(args entities.Equipment) (*entities.Equipment, error) {
	create, err := e.queries.CreateEquipment(e.ctx, persistence.CreateEquipmentParams{
		TypeID:       args.TypeID,
		Name:         args.Name,
		SerialNumber: args.SerialNumber,
		Notes:        args.Notes,
	})

	res := entities.Equipment{}
	mapper.AutoMapper(&create, &res)

	return &res, err
}

// DeleteEquipmentByID implements repositories.IEquipementRepository.
func (e *EquipmentRepository) DeleteEquipmentByID(id int32) error {
	err := e.queries.DeleteEquipmentByID(e.ctx, id)
	return err
}

// GetEquipmentByID implements repositories.IEquipementRepository.
func (e *EquipmentRepository) GetEquipmentByID(id int32) (*entities.Equipment, error) {
	data, err := e.queries.GetEquipmentByID(e.ctx, id)
	if err != nil {
		return nil, err
	}
	res := &entities.Equipment{}
	mapper.AutoMapper(&data, res)
	return res, nil
}

// ListAllEquipement implements repositories.IEquipementRepository.
func (e *EquipmentRepository) ListAllEquipement(offset int32, limit int32) ([]entities.Equipment, error) {
	data, err := e.queries.ListAllEquipments(e.ctx, persistence.ListAllEquipmentsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	res := []entities.Equipment{}
	mapper.MapperSlice(&data, &res)

	return res, nil
}

// UpdateEquipmentByID implements repositories.IEquipementRepository.
func (e *EquipmentRepository) UpdateEquipmentByID(args entities.Equipment) error {
	err := e.queries.UpdateEquipmentByID(e.ctx, persistence.UpdateEquipmentByIDParams{
		ID:           args.ID,
		TypeID:       args.TypeID,
		Name:         args.Name,
		SerialNumber: args.SerialNumber,
		Notes:        args.Notes,
	})
	return err
}

func NewEquipmentRepository(ctx context.Context, queries *persistence.Queries) repositories.IEquipementRepository {
	return &EquipmentRepository{
		ctx:     ctx,
		queries: queries,
	}
}
