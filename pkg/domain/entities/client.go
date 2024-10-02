package entities

import "database/sql"

type Client struct {
	ID            int32
	Name          string
	Address       sql.NullString
	District      sql.NullString
	City          sql.NullString
	Country       sql.NullString
	Phone         sql.NullString
	Ruc           sql.NullString
	ContactPerson sql.NullString
	Email         sql.NullString
	Website       sql.NullString
	AddressLine2  sql.NullString
	PostalCode    sql.NullString
	Fax           sql.NullString
	Notes         sql.NullString
}
