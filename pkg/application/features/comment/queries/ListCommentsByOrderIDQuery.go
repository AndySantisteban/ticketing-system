package comment_features

import "InfositelOR/pkg/domain/entities"

type ListCommentsByOrderIDQuery struct {
	Offset  int32
	Limit   int32
	OrderID int32
}

type ListCommentsByOrderIDQueryResponse = []entities.Comment
