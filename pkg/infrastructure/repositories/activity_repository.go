package repositories

import (
	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
	"context"
	"database/sql"

	"github.com/devfeel/mapper"
)

type ActivityRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

func NewActivityRepository(ctx context.Context, queries *persistence.Queries) repositories.IActivityRepository {
	return &ActivityRepository{
		ctx:     ctx,
		queries: queries,
	}
}

// CreateActivity implements repositories.IActivityRepository.
func (r *ActivityRepository) CreateActivity(args entities.Activity) (entities.Activity, error) {
	CreateActivity, err := r.queries.CreateActivity(r.ctx, persistence.CreateActivityParams{
		OrderID: args.OrderID,
		UserID:  args.UserID,
		Action:  args.Action,
		Details: args.Details,
	})
	RespActivity := &entities.Activity{}
	mapper.AutoMapper(&CreateActivity, RespActivity)

	return *RespActivity, err
}

// GetActivityByID implements repositories.IActivityRepository.
func (r *ActivityRepository) GetActivityByID(id int32) (entities.Activity, error) {
	ActivityByID, err := r.queries.GetActivityByID(r.ctx, id)
	RespActivity := &entities.Activity{}
	mapper.AutoMapper(&ActivityByID, RespActivity)

	return *RespActivity, err
}

// ListActivityByOrderID implements repositories.IActivityRepository.
func (r *ActivityRepository) ListActivityByOrderID(offset int32, limit int32, OrderID *int32) ([]entities.Activity, error) {
	ListActivityByID, err := r.queries.ListActivityByOrderID(r.ctx, persistence.ListActivityByOrderIDParams{
		Offset: offset,
		Limit:  limit,
		OrderID: sql.NullInt32{
			Valid: true,
			Int32: *OrderID,
		},
	})
	if err != nil {
		return nil, err
	}
	RespActivity := []entities.Activity{}
	mapper.MapperSlice(&ListActivityByID, &RespActivity)

	return RespActivity, err
}
