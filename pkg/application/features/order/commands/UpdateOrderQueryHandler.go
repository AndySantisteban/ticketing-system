package order_commands

import (
	"context"
	"database/sql"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
)

type UpdateOrderQueryHandler struct {
	ctx        context.Context
	repository repositories.IOrderRepository
}

func NewUpdateOrderQQueryHandler(ctx context.Context, respository repositories.IOrderRepository) *UpdateOrderQueryHandler {

	return &UpdateOrderQueryHandler{
		ctx,
		respository,
	}
}

func (c *UpdateOrderQueryHandler) Handler(args UpdateOrderQuery) UpdateOrderQueryResponse {

	err := c.repository.UpdateOrderByID(entities.Order{
		ID: args.ID,
		ClientID: sql.NullInt32{
			Valid: args.ClientID != nil,
			Int32: *args.ClientID,
		},
		EquipmentID: sql.NullInt32{
			Int32: *args.Equipement,
			Valid: args.Equipement != nil,
		},
		OrderNumber: args.OrderNumber,
		ReportedIssue: sql.NullString{
			Valid:  args.ReportedIssue != nil,
			String: *args.ReportedIssue,
		},
		Diagnosis: sql.NullString{
			Valid:  args.Diagnosis != nil,
			String: *args.Diagnosis,
		},
		Solution: sql.NullString{
			Valid:  args.Solution != nil,
			String: *args.Solution,
		},
		EstimatedTime: sql.NullInt64{
			Valid: args.EstimatedTime != nil,
			Int64: *args.EstimatedTime,
		},
		Budget: sql.NullString{
			Valid:  args.Budget != nil,
			String: *args.Budget,
		},
		StatusID: sql.NullInt32{
			Valid: args.StatusID != nil,
			Int32: *args.StatusID,
		},
		AssignedTo: sql.NullInt32{
			Valid: args.AssignedTo != nil,
			Int32: *args.AssignedTo,
		},
		CreationDate: sql.NullTime{
			Valid: args.CreationDate != nil,
			Time:  *args.CreationDate,
		},
		Priority: sql.NullString{
			Valid:  args.Priority != nil,
			String: *args.Priority,
		},
	})

	return err
}
