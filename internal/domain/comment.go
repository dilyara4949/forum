package domain

import "context"

type Comment struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	PostID  int    `json:"post_id"`
}

type CommentRepository interface {
	Create(c context.Context, Comment *Comment) (*Comment, error)
	Get(c context.Context, postID int) ([]Comment, error)
	Delete(c context.Context, postID int, email string) error
}

type CommentUsecase interface {
	Create(c context.Context, Comment *Comment) (*Comment, error)
	Get(c context.Context, postID int) ([]Comment, error)
	Delete(c context.Context, postID int, email string) error
}
