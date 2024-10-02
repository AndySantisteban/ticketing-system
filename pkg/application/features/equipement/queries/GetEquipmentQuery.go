package equipment_features

import "InfositelOR/pkg/domain/entities"

type GetEquipmentQuery struct {
	ID int32
}

type GetEquipmentQueryResponse = entities.Equipment
