package equipment_type_commands

import (
	"InfositelOR/pkg/domain/entities"
)

type CreateEquipmentTypeQuery struct {
	ID   int32
	Name string
}

type CreateEquipmentTypeQueryResponse = entities.EquipmentType
