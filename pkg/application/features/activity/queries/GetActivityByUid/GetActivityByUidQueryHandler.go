package features_activity

import (
	"InfositelOR/pkg/domain/repositories"
	"context"
)

type GetActivityByUidQueryHandler struct {
	ctx                context.Context
	activityRepository repositories.IActivityRepository
}

func NewLGetActivityByUidQueryHandler(ctx context.Context, activityRepository repositories.IActivityRepository) *GetActivityByUidQueryHandler {

	constructor := &GetActivityByUidQueryHandler{
		ctx:                ctx,
		activityRepository: activityRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *GetActivityByUidQueryHandler) Handler(param GetActivityByUidQuery) (*GetActivityByUidResponse, error) {
	id := param.Id
	response, err := r.activityRepository.GetActivityByID(id)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
