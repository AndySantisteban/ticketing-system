package repositories

import (
	"InfositelOR/pkg/domain/entities"
)

type ICommentRepository interface {
	ListCommentsByOrderID(offset int32, limit int32, OrderID *int32) ([]entities.Comment, error)
	GetCommentByID(id int32) (*entities.Comment, error)
	CreateComment(args entities.Comment) (*entities.Comment, error)
	UpdateCommentByID(arg entities.Comment) error
	DeleteCommentByID(id int32) error
}
