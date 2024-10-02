package entities

import "database/sql"

type Activity struct {
	ID      int32
	OrderID sql.NullInt32
	UserID  sql.NullInt32
	Date    sql.NullTime
	Action  sql.NullString
	Details sql.NullString
}
