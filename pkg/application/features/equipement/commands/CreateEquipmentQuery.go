package equipment_commands

import (
	"InfositelOR/pkg/domain/entities"
)

type CreateEquipmentQuery struct {
	ID           int32
	TypeID       *int32
	Name         string
	SerialNumber string
	Notes        *string
}

type CreateEquipmentQueryResponse = entities.Equipment
