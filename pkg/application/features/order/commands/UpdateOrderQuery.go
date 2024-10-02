package order_commands

import "time"

type UpdateOrderQuery struct {
	ID            int32
	ClientID      *int32
	Equipement    *int32
	OrderNumber   string
	ReportedIssue *string
	Diagnosis     *string
	Solution      *string
	EstimatedTime *int64
	Budget        *string
	StatusID      *int32
	AssignedTo    *int32
	CreationDate  *time.Time
	Priority      *string
}

type UpdateOrderQueryResponse = error
