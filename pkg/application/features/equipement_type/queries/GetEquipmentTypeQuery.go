package equipment_type_features

import "InfositelOR/pkg/domain/entities"

type GetEquipmentTypeQuery struct {
	ID int32
}

type GetEquipmentTypeQueryResponse = entities.EquipmentType
