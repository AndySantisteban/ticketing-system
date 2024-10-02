package repositories

import "InfositelOR/pkg/domain/entities"

type IUserRepository interface {
	ListAllUser(offset int32, limit int32) ([]entities.User, error)
	GetUserByID(id int32) (*entities.User, error)
	DeleteUserByID(id int32) error
	CreateUser(args entities.User) (*entities.User, error)
	UpdateUserByID(args entities.User) error
}
