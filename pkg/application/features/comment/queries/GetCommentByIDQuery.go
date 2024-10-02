package comment_features

import "InfositelOR/pkg/domain/entities"

type GetCommentByIDQuery struct {
	Id int32
}

type GetCommentByIDQueryResponse = entities.Comment
