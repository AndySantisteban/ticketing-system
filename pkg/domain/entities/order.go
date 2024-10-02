package entities

import "database/sql"

type Order struct {
	ID            int32
	ClientID      sql.NullInt32
	EquipmentID   sql.NullInt32
	OrderNumber   string
	ReportedIssue sql.NullString
	Diagnosis     sql.NullString
	Solution      sql.NullString
	EstimatedTime sql.NullInt64
	Budget        sql.NullString
	StatusID      sql.NullInt32
	AssignedTo    sql.NullInt32
	CreationDate  sql.NullTime
	Priority      sql.NullString
}
