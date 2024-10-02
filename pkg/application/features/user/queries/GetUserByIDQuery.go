package user_features

import "InfositelOR/pkg/domain/entities"

type GetUserByIDQuery struct {
	ID int32
}

type GetUserByIDQueryResponse = entities.User
