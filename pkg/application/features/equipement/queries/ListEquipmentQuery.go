package equipment_features

import "InfositelOR/pkg/domain/entities"

type ListAllEquipmentQuery struct {
	Offset int32
	Limit  int32
}

type ListAllEquipmentResponse = []entities.Equipment
