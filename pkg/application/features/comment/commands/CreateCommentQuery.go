package comment_commands

import (
	"InfositelOR/pkg/domain/entities"
	"time"
)

type CreateCommentQuery struct {
	ID      int32
	OrderID int32
	UserID  int32
	Date    *time.Time
	Comment *string
}

type CreateCommentQueryResponse = entities.Comment
