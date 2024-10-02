package repositories

import (
	"InfositelOR/pkg/domain/entities"
)

type IActivityRepository interface {
	ListActivityByOrderID(offset int32, limit int32, OrderID *int32) ([]entities.Activity, error)
	GetActivityByID(id int32) (entities.Activity, error)
	CreateActivity(args entities.Activity) (entities.Activity, error)
}
