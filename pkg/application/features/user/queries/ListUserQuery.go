package user_features

import "InfositelOR/pkg/domain/entities"

type ListUserQuery struct {
	Limit  int32
	Offset int32
}

type ListUserQueryResponse = []entities.User
