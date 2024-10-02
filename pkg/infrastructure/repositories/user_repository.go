package repositories

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
	"context"

	"github.com/devfeel/mapper"
)

type UserRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// DeleteUserByID implements repositories.IUserRepository.
func (u *UserRepository) DeleteUserByID(id int32) error {
	err := u.queries.DeleteUserByID(u.ctx, id)
	return err
}

// CreateUser implements repositories.IUserRepository.
func (u *UserRepository) CreateUser(args entities.User) (*entities.User, error) {
	response, err := u.queries.CreateUser(u.ctx, persistence.CreateUserParams{
		Name:           args.Name,
		Email:          args.Email,
		PermissionType: args.PermissionType,
		InactiveStatus: args.InactiveStatus,
		Password:       args.Password,
	})
	if err != nil {
		return nil, err
	}
	entity := entities.User{}
	mapper.AutoMapper(&entity, &response)

	return &entity, nil
}

// GetUserByID implements repositories.IUserRepository.
func (u *UserRepository) GetUserByID(id int32) (*entities.User, error) {
	item, err := u.queries.GetUserByID(u.ctx, id)
	if err != nil {
		return nil, err
	}
	entity := entities.User{}
	mapper.AutoMapper(&item, &entity)

	return &entity, nil
}

// ListAllUser implements repositories.IUserRepository.
func (u *UserRepository) ListAllUser(offset int32, limit int32) ([]entities.User, error) {
	item, err := u.queries.ListAllUsers(u.ctx, persistence.ListAllUsersParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	Entity := []entities.User{}
	mapper.MapperSlice(&item, &Entity)
	return Entity, nil
}

// UpdateUserByID implements repositories.IUserRepository.
func (u *UserRepository) UpdateUserByID(args entities.User) error {
	err := u.queries.UpdateUserByID(u.ctx, persistence.UpdateUserByIDParams{
		Name:           args.Name,
		Email:          args.Email,
		PermissionType: args.PermissionType,
		ID:             args.ID,
		InactiveStatus: args.InactiveStatus,
		Password:       args.Password,
	})

	return err
}

func NewUserRepository(ctx context.Context, queries *persistence.Queries) repositories.IUserRepository {
	return &UserRepository{
		ctx:     ctx,
		queries: queries,
	}
}
