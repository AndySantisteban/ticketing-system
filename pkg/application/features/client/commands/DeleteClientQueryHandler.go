package client_commands

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type DeleteClientByIDQueryHandler struct {
	ctx              context.Context
	clientRepository repositories.IClientRepository
}

func NewDeleteClientByIDQueryHandler(ctx context.Context, clientRepository repositories.IClientRepository) *DeleteClientByIDQueryHandler {

	return &DeleteClientByIDQueryHandler{
		ctx,
		clientRepository,
	}
}

func (d *DeleteClientByIDQueryHandler) Handler(args DeleteClientByIDQuery) DeleteClientByIDResponse {

	err := d.clientRepository.DeleteClientByID(args.Id)

	return err
}
