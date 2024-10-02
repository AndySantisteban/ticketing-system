package client_features

import "InfositelOR/pkg/domain/entities"

type ListAllClientsQuery struct {
	Offset int32
	Limit  int32
}

type ListAllClientsResponse = []entities.Client
