package features_activity

import (
	"InfositelOR/pkg/domain/entities"
)

type GetActivityByUidQuery struct {
	Id int32
}

type GetActivityByUidResponse = entities.Activity
