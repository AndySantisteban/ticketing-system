package equipment_type_features

import "InfositelOR/pkg/domain/entities"

type ListAllEquipmentTypeQuery struct {
	Offset int32
	Limit  int32
}

type ListAllEquipmentTypeResponse = []entities.EquipmentType
