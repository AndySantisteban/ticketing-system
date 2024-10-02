package features_activity

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"context"
	"database/sql"
)

type CreateActivityQueryHandler struct {
	ctx                context.Context
	activityRepository repositories.IActivityRepository
}

func NewCreateActivityQueryHandler(ctx context.Context, activityRepository repositories.IActivityRepository) *CreateActivityQueryHandler {

	constructor := &CreateActivityQueryHandler{
		ctx:                ctx,
		activityRepository: activityRepository,
	}

	return constructor
}

// Handler implements cqrs.ImplementsHandler.
func (r *CreateActivityQueryHandler) Handler(param CreateActivityQuery) (*CreateActivityResponse, error) {

	response, err := r.activityRepository.CreateActivity(entities.Activity{
		ID: param.ID,
		OrderID: sql.NullInt32{
			Int32: param.OrderID,
		},
		UserID: sql.NullInt32{
			Int32: param.UserID,
		},
		Date: sql.NullTime{
			Time: param.Date,
		},
		Action: sql.NullString{
			String: param.Action,
		},
		Details: sql.NullString{
			String: param.Details,
		},
	})
	if err != nil {
		return nil, err
	}
	return &response, err
}
