package comment

import "time"

type ListCommentRouteDTO struct {
	Offset  int32 `query:"offset"`
	Limit   int32 `query:"limit"`
	OrderID int32 `query:"orderID"`
}

type GetCommentByUidRouteDTO struct {
	Id int32 `query:"id"`
}

type CreateCommentRouteDTO struct {
	ID      int32      `json:"id"`
	OrderID int32      `json:"orderID"`
	UserID  int32      `json:"userID"`
	Date    *time.Time `json:"date"`
	Comment *string    `json:"comment"`
}
type UpdateCommentRouteDTO struct {
	ID      int32      `json:"id"`
	OrderID int32      `json:"orderID"`
	Comment *string    `json:"comment"`
	UserID  int32      `json:"userID"`
	Date    *time.Time `json:"date"`
}
type DeleteCommentRouteDTO struct {
	ID int32 `query:"id"`
}
