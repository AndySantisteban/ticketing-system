package entities

import "database/sql"

type UserStatus struct {
	ID             int32
	UserID         sql.NullInt32
	OrderID        sql.NullInt32
	AssignmentDate sql.NullTime
}
