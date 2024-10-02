package entities

import "database/sql"

type User struct {
	ID             int32
	Name           string
	Email          string
	PermissionType string
	CreationDate   sql.NullTime
	InactiveStatus sql.NullString
	Password       sql.NullString
}
