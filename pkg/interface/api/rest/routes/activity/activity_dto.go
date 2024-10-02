package activity

import "time"

type ListActivityByOrderIDRouteDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
	Id     int32 `query:"id"`
}

type GetActivityByUidRouteDTO struct {
	Id int32 `query:"id"`
}

type CreateActivityQuery struct {
	ID      int32     `json:"id"`
	OrderID int32     `json:"orderID,omitempty"`
	UserID  int32     `json:"userID,omitempty"`
	Date    time.Time `json:"date,omitempty"`
	Action  string    `json:"action,omitempty"`
	Details string    `json:"details,omitempty"`
}
