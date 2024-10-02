package entities

import "database/sql"

type Comment struct {
	ID      int32
	OrderID sql.NullInt32
	UserID  sql.NullInt32
	Date    sql.NullTime
	Comment sql.NullString
}
