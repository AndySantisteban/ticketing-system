package repositories

import (
	"InfositelOR/pkg/domain/entities"
)

type IClientRepository interface {
	ListAllClients(offset int32, limit int32) ([]entities.Client, error)
	GetClientByID(id int32) (*entities.Client, error)
	CreateClient(args entities.Client) (*entities.Client, error)
	UpdateClientByID(arg entities.Client) error
	DeleteClientByID(id int32) error
}
