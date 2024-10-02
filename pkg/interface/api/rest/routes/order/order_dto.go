package order

import "time"

type ListOrdersRouteDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

type GetOrdersByUidRouteDTO struct {
	Id int32 `query:"id"`
}

type CreateRouteDTO struct {
	Id            int32      `json:"id"`
	StatusID      *int32     `json:"statusID"`
	AssignedTo    *int32     `json:"assignedTo"`
	CreationDate  *time.Time `json:"creationDate"`
	ClientID      *int32     `json:"clientID"`
	Priority      *string    `json:"priority"`
	Equipement    *int32     `json:"equipement"`
	OrderNumber   string     `json:"orderNumber"`
	ReportedIssue *string    `json:"reportedIssue"`
	Diagnosis     *string    `json:"diagnosis"`
	Solution      *string    `json:"solution"`
	EstimatedTime *int64     `json:"estimatedTime"`
	Budget        *string    `json:"budget"`
}

type UpdateRouteDTO struct {
	Id            int32      `json:"id"`
	StatusID      *int32     `json:"statusID"`
	AssignedTo    *int32     `json:"assignedTo"`
	CreationDate  *time.Time `json:"creationDate"`
	ClientID      *int32     `json:"clientID"`
	Priority      *string    `json:"priority"`
	Equipement    *int32     `json:"equipement"`
	OrderNumber   string     `json:"orderNumber"`
	ReportedIssue *string    `json:"reportedIssue"`
	Diagnosis     *string    `json:"diagnosis"`
	Solution      *string    `json:"solution"`
	EstimatedTime *int64     `json:"estimatedTime"`
	Budget        *string    `json:"budget"`
}
