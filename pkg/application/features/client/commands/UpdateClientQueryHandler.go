package client_commands

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
	"database/sql"
)

type UpdateClientQueryHandler struct {
	ctx              context.Context
	clientRepository repositories.IClientRepository
}

func NewUpdateClientQueryHandler(ctx context.Context, clientRepository repositories.IClientRepository) *UpdateClientQueryHandler {

	return &UpdateClientQueryHandler{
		ctx,
		clientRepository,
	}
}

func (u *UpdateClientQueryHandler) Handler(args UpdateClientQuery) UpdateClientResponse {

	err := u.clientRepository.UpdateClientByID(entities.Client{
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

	return err
}
