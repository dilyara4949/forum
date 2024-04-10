package domain

import "context"

type Like struct {
	PostID int    `json:"post_id"`
	Like   int   `json:"like"`
}

type LikeRepository interface {
	Create(c context.Context, like *Like, email string) (*Like, error)
	Get(c context.Context, postID int) (*Like, error)
	Delete(c context.Context, postID int, email string) (*Like, error)
}

type LikeUsecase interface {
	Create(c context.Context, like *Like, email string) (*Like, error)
	Get(c context.Context, postID int) (*Like, error)
	Delete(c context.Context, postID int, email string) (*Like, error)
}