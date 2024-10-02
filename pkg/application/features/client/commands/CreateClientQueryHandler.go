package client_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
	"database/sql"
)

type CreateClientQueryHandler struct {
	ctx              context.Context
	clientRepository repositories.IClientRepository
}

func NewCreateClientQueryHandler(ctx context.Context, clientRepository repositories.IClientRepository) *CreateClientQueryHandler {
	return &CreateClientQueryHandler{
		ctx,
		clientRepository,
	}
}

func (c *CreateClientQueryHandler) Handler(args CreateClientQuery) (*CreateClientQueryResponse, error) {

	response, err := c.clientRepository.CreateClient(entities.Client{
		ID:   args.ID,
		Name: args.Name,
		Address: sql.NullString{
			String: args.Address,
			Valid:  true,
		},
		District: sql.NullString{
			String: args.District,
			Valid:  true,
		},
		City: sql.NullString{
			String: args.City,
			Valid:  true,
		},
		Country: sql.NullString{
			String: args.Country,
			Valid:  true,
		},
		Phone: sql.NullString{
			String: args.Phone,
			Valid:  true,
		},
		Ruc: sql.NullString{
			String: args.Ruc,
			Valid:  true,
		},
		ContactPerson: sql.NullString{
			String: args.ContactPerson,
			Valid:  true,
		},
		Email: sql.NullString{
			String: args.Email,
			Valid:  true,
		},
		Website: sql.NullString{
			String: args.Website,
			Valid:  true,
		},
		AddressLine2: sql.NullString{
			String: args.AddressLine2,
			Valid:  true,
		},
		PostalCode: sql.NullString{
			String: args.PostalCode,
			Valid:  true,
		},
		Fax: sql.NullString{
			String: args.Fax,
			Valid:  true,
		},
		Notes: sql.NullString{
			String: args.Notes,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, err
	}

	return &response, nil

}
