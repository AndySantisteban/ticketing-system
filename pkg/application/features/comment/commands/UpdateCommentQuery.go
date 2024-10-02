package comment_commands

import "time"

type UpdateCommentQuery struct {
	ID      int32
	OrderID int32
	UserID  int32
	Date    *time.Time
	Comment *string
}

type UpdateCommentQueryResponse = error
