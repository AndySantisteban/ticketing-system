package client_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type GetClientByIDQueryHandler struct {
	ctx              context.Context
	clientRepository repositories.IClientRepository
}

func NewGetClientByIDQueryHandler(ctx context.Context, clientRepository repositories.IClientRepository) *GetClientByIDQueryHandler {

	constructor := &GetClientByIDQueryHandler{
		ctx:              ctx,
		clientRepository: clientRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *GetClientByIDQueryHandler) Handler(param GetClientByIDQuery) (*GetCliendByIDResponse, error) {
	id := param.Id
	response, err := r.clientRepository.GetClientByID(id)

	if err != nil {
		return nil, err
	}

	return response, nil
}
