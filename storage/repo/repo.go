package repo

import (
	c "bay_store/comment_service/genproto/comment"
)

type CommentStoreI interface {
	WriteComment(*c.CommentRequest) (*c.CommentResponse, error)
	GetProductComments(*c.IdRequest) (*c.Comments, error)
	GetUserComments(*c.IdRequest) (*c.Comments, error)
	DeleteComment(*c.IdRequest) (*c.CommentResponse, error)
}
