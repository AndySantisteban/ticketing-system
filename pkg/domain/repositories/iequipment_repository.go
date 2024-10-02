package repositories

import "InfositelOR/pkg/domain/entities"

type IEquipementRepository interface {
	ListAllEquipement(offset int32, limit int32) ([]entities.Equipment, error)
	GetEquipmentByID(id int32) (*entities.Equipment, error)
	CreateEquipment(args entities.Equipment) (*entities.Equipment, error)
	UpdateEquipmentByID(args entities.Equipment) error
	DeleteEquipmentByID(id int32) error
}
