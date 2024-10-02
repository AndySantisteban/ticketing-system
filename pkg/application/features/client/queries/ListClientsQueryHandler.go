package client_features

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type ListClientsQueryHandler struct {
	ctx              context.Context
	clientRepository repositories.IClientRepository
}

func NewListClientsQueryHandler(ctx context.Context, clientRepository repositories.IClientRepository) *ListClientsQueryHandler {

	constructor := &ListClientsQueryHandler{
		ctx:              ctx,
		clientRepository: clientRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *ListClientsQueryHandler) Handler(param ListAllClientsQuery) (*ListAllClientsResponse, error) {

	response, err := r.clientRepository.ListAllClients(param.Offset, param.Limit)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
