package features_activity

import (
	"InfositelOR/pkg/domain/entities"
)

type ListActivityByOrderIDQuery struct {
	Offset int32
	Limit  int32
	Id     int32
}

type ListActivityByOrderIDResponse = []entities.Activity
