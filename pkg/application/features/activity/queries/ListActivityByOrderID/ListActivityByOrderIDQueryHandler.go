package features_activity

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type ListActivityByOrderIDQueryHandler struct {
	ctx                context.Context
	activityRepository repositories.IActivityRepository
}

func NewListActivityByOrderIDQueryHandler(ctx context.Context, activityRepository repositories.IActivityRepository) *ListActivityByOrderIDQueryHandler {

	constructor := &ListActivityByOrderIDQueryHandler{
		ctx:                ctx,
		activityRepository: activityRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *ListActivityByOrderIDQueryHandler) Handler(c ListActivityByOrderIDQuery) (*ListActivityByOrderIDResponse, error) {
	offset := c.Offset
	limit := c.Limit
	id := &c.Id

	response, err := r.activityRepository.ListActivityByOrderID(offset, limit, id)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
