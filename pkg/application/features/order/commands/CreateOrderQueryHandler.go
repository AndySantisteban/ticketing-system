package order_commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
)

type CreateOrderQueryHandler struct {
	ctx         context.Context
	respository repositories.IOrderRepository
}

func NewCreateOrderQueryHandler(ctx context.Context, respository repositories.IOrderRepository) *CreateOrderQueryHandler {

	return &CreateOrderQueryHandler{
		ctx,
		respository,
	}
}

func (c *CreateOrderQueryHandler) Handler(args CreateOrderQuery) (*CreateOrderQueryResponse, error) {
	fmt.Println(*args.ClientID)
	item, err := c.respository.CreateOrder(entities.Order{
		ID: args.ID,
		ClientID: sql.NullInt32{
			Valid: true,
			Int32: *args.ClientID,
		},
		EquipmentID: sql.NullInt32{
			Int32: *args.Equipement,
			Valid: true,
		},
		OrderNumber: args.OrderNumber,
		ReportedIssue: sql.NullString{
			Valid:  true,
			String: *args.ReportedIssue,
		},
		Diagnosis: sql.NullString{
			Valid:  true,
			String: *args.Diagnosis,
		},
		Solution: sql.NullString{
			Valid:  true,
			String: *args.Solution,
		},
		EstimatedTime: sql.NullInt64{
			Valid: true,
			Int64: *args.EstimatedTime,
		},
		Budget: sql.NullString{
			Valid:  args.Budget != nil,
			String: *args.Budget,
		},
		StatusID: sql.NullInt32{
			Valid: true,
			Int32: *args.StatusID,
		},
		AssignedTo: sql.NullInt32{
			Valid: true,
			Int32: *args.AssignedTo,
		},
		CreationDate: sql.NullTime{
			Valid: true,
			Time:  time.Time{},
		},
		Priority: sql.NullString{
			Valid:  true,
			String: *args.Priority,
		},
	})

	if err != nil {
		return nil, err
	}

	return item, nil
}
