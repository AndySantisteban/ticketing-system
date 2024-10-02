package repositories

import (
	"context"
	"database/sql"

	"github.com/devfeel/mapper"

	"InfositelOR/pkg/domain/entities"
	"InfositelOR/pkg/domain/repositories"
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
)

type CommentRepository struct {
	ctx     context.Context
	queries *persistence.Queries
}

// CreateComment implements repositories.ICommentRepository.
func (c *CommentRepository) CreateComment(args entities.Comment) (*entities.Comment, error) {
	data, err := c.queries.CreateComment(c.ctx, persistence.CreateCommentParams{
		OrderID: args.OrderID,
		UserID:  args.UserID,
		Comment: args.Comment,
	})
	if err != nil {
		return nil, err
	}
	ResComment := entities.Comment{}
	mapper.AutoMapper(&data, &ResComment)
	return &ResComment, nil
}

// DeleteCommentByID implements repositories.ICommentRepository.
func (c *CommentRepository) DeleteCommentByID(id int32) error {
	err := c.queries.DeleteCommentByID(c.ctx, id)
	return err
}

// GetCommentByID implements repositories.ICommentRepository.
func (c *CommentRepository) GetCommentByID(id int32) (*entities.Comment, error) {
	data, err := c.queries.GetCommentByID(c.ctx, id)
	if err != nil {
		return nil, err
	}
	ResComment := &entities.Comment{}
	mapper.AutoMapper(&data, ResComment)
	return ResComment, nil
}

// ListCommentsByOrderID implements repositories.ICommentRepository.
func (c *CommentRepository) ListCommentsByOrderID(offset int32, limit int32, OrderID *int32) ([]entities.Comment, error) {
	data, err := c.queries.ListCommentsByOrderID(c.ctx, persistence.ListCommentsByOrderIDParams{
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
	RespComment := []entities.Comment{}
	mapper.MapperSlice(&data, &RespComment)

	return RespComment, nil
}

// UpdateCommentByID implements repositories.ICommentRepository.
func (c *CommentRepository) UpdateCommentByID(arg entities.Comment) error {
	err := c.queries.UpdateCommentByID(c.ctx, persistence.UpdateCommentByIDParams{
		ID:      arg.ID,
		Comment: arg.Comment,
	})
	return err
}

func NewCommentRepository(ctx context.Context, queries *persistence.Queries) repositories.ICommentRepository {
	return &CommentRepository{
		ctx:     ctx,
		queries: queries,
	}
}
