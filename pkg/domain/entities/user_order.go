package entities

import "database/sql"

type UserOrder struct {
	ID             int32
	UserID         sql.NullInt32
	OrderID        sql.NullInt32
	AssignmentDate sql.NullTime
}
