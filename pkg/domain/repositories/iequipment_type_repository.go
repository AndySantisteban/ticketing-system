package repositories

import "InfositelOR/pkg/domain/entities"

type IEquipmentTypeRepository interface {
	ListAllEquipmentType(offset int32, limit int32) ([]entities.EquipmentType, error)
	GetEquipmentTypeByID(id int32) (*entities.EquipmentType, error)
	CreateEquipmentType(args entities.EquipmentType) (*entities.EquipmentType, error)
	UpdateEquipmentTypeByID(args entities.EquipmentType) error
	DeleteEquipmentTypeByID(id int32) error
}
