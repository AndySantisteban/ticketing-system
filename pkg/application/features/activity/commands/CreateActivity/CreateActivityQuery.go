package features_activity

import (
	"InfositelOR/pkg/domain/entities"
	"time"
)

type CreateActivityQuery struct {
	ID      int32
	OrderID int32
	UserID  int32
	Date    time.Time
	Action  string
	Details string
}

type CreateActivityResponse = entities.Activity
