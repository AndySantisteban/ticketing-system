// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package persistence

import (
	"database/sql"
)

type Activity struct {
	ID      int32
	OrderID sql.NullInt32
	UserID  sql.NullInt32
	Date    sql.NullTime
	Action  sql.NullString
	Details sql.NullString
}

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

type Comment struct {
	ID      int32
	OrderID sql.NullInt32
	UserID  sql.NullInt32
	Date    sql.NullTime
	Comment sql.NullString
}

type Equipment struct {
	ID           int32
	TypeID       sql.NullInt32
	Name         string
	SerialNumber string
	Notes        sql.NullString
}

type EquipmentType struct {
	ID   int32
	Name string
}

type Order struct {
	ID            int32
	ClientID      sql.NullInt32
	EquipmentID   sql.NullInt32
	OrderNumber   string
	ReportedIssue sql.NullString
	Diagnosis     sql.NullString
	Solution      sql.NullString
	EstimatedTime sql.NullString
	Budget        sql.NullString
	StatusID      sql.NullInt32
	AssignedTo    sql.NullInt32
	CreationDate  sql.NullTime
	Priority      sql.NullString
}

type OrderStatus struct {
	ID   int32
	Name string
}

type User struct {
	ID             int32
	Name           string
	Email          string
	Password       sql.NullString
	PermissionType string
	CreationDate   sql.NullTime
	InactiveStatus sql.NullString
}

type UserOrder struct {
	ID             int32
	UserID         sql.NullInt32
	OrderID        sql.NullInt32
	AssignmentDate sql.NullTime
}
