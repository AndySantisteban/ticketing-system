package repositories

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
	"context"

	"github.com/devfeel/mapper"
)

type ClientRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// CreateClient implements repositories.IClientRepository.
func (c *ClientRepository) CreateClient(args entities.Client) (*entities.Client, error) {
	CreateActivity, err := c.queries.CreateClient(c.ctx, persistence.CreateClientParams{
		Name:          args.Name,
		Address:       args.Address,
		District:      args.District,
		City:          args.City,
		Country:       args.Country,
		Phone:         args.Phone,
		Ruc:           args.Ruc,
		ContactPerson: args.ContactPerson,
		Email:         args.Email,
		Website:       args.Website,
		AddressLine2:  args.AddressLine2,
		PostalCode:    args.PostalCode,
		Fax:           args.Fax,
		Notes:         args.Notes,
	})
	RespActivity := &entities.Client{}
	mapper.Mapper(&CreateActivity, RespActivity)

	return RespActivity, err
}

// DeleteClientByID implements repositories.IClientRepository.
func (c *ClientRepository) DeleteClientByID(id int32) error {
	err := c.queries.DeleteClientByID(c.ctx, id)
	return err
}

// GetClientByID implements repositories.IClientRepository.
func (c *ClientRepository) GetClientByID(id int32) (*entities.Client, error) {
	data, err := c.queries.GetClientByID(c.ctx, id)
	if err != nil {
		return nil, err
	}
	RespActivity := &entities.Client{}
	mapper.AutoMapper(&data, RespActivity)
	return RespActivity, nil
}

// ListAllClients implements repositories.IClientRepository.
func (c *ClientRepository) ListAllClients(offset int32, limit int32) ([]entities.Client, error) {
	data, err := c.queries.ListAllClients(c.ctx, persistence.ListAllClientsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	RespActivity := []entities.Client{}
	mapper.MapperSlice(&data, &RespActivity)

	return RespActivity, nil
}

// UpdateClientByID implements repositories.IClientRepository.
func (c *ClientRepository) UpdateClientByID(arg entities.Client) error {
	err := c.queries.UpdateClientByID(c.ctx, persistence.UpdateClientByIDParams{
		ID:            arg.ID,
		Name:          arg.Name,
		Address:       arg.Address,
		District:      arg.District,
		City:          arg.City,
		Country:       arg.Country,
		Phone:         arg.Phone,
		Ruc:           arg.Ruc,
		ContactPerson: arg.ContactPerson,
		Email:         arg.Email,
		Website:       arg.Website,
		AddressLine2:  arg.AddressLine2,
		PostalCode:    arg.PostalCode,
		Fax:           arg.Fax,
		Notes:         arg.Notes,
	})
	return err
}

func NewClientRepository(ctx context.Context, queries *persistence.Queries) repositories.IClientRepository {
	return &ClientRepository{
		ctx:     ctx,
		queries: queries,
	}
}
