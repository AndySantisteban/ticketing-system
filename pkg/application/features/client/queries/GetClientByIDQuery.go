package client_features

import "InfositelOR/pkg/domain/entities"

type GetClientByIDQuery struct {
	Id int32
}

type GetCliendByIDResponse = entities.Client
