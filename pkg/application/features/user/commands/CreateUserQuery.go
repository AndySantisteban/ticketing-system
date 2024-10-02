package user_commands

import (
	"InfositelOR/pkg/domain/entities"
	"time"
)

type CreateUserQuery struct {
	ID             int32
	Name           string
	Email          string
	PermissionType string
	CreationDate   *time.Time
	InactiveStatus *string
	Password       string
}

type CreateUserQueryResponse = entities.User
