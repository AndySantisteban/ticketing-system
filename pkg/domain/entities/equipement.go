package entities

import "database/sql"

type Equipment struct {
	ID           int32
	TypeID       sql.NullInt32
	Name         string
	SerialNumber string
	Notes        sql.NullString
}
