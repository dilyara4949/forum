package domain

import "context"

type PostRequest struct {
	Email    string `json:"email"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category string `json:"category"`
}

type Post struct {
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments"`
	Category string    `json:"category"`
}

type PostRepository interface {
	Create(c context.Context, Post *PostRequest) (*Post, error)
	GetOwn(c context.Context, email string) ([]Post, error)
	Delete(c context.Context, id int) error
	GetAll(c context.Context) ([]Post, error)
}

type PostUsecase interface {
	Create(c context.Context, Post *PostRequest) (*Post, error)
	GetOwn(c context.Context, email string) ([]Post, error)
	Delete(c context.Context, id int) error
	GetAll(c context.Context) ([]Post, error)
}
