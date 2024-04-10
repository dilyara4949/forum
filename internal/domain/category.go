package domain

import "context"

type Category struct {
	Name       string `json:"name"`
}

type CategoryRepository interface {
	Create(c context.Context, category *Category) (*Category, error)
	Get(c context.Context, id string) (*Category, error)
	Delete(c context.Context, name string) error
	GetAll(c context.Context) ([]Category, error)
}

type CategoryUsecase interface {
	Create(c context.Context, category *Category) (*Category, error)
	Get(c context.Context, name string) (*Category, error)
	Delete(c context.Context, name string) error
	GetAll(c context.Context) ([]Category, error)
}
