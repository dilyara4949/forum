package usecase

import (
	"context"
	"forum/internal/domain"
	"time"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}

func (c *categoryUsecase) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.Create(ctx, category)
}

func (c *categoryUsecase) Get(ctx context.Context, name string) (*domain.Category, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.Get(ctx, name)
}

func (c *categoryUsecase) Delete(ctx context.Context, name string) error {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.Delete(ctx, name)
}

func (c *categoryUsecase) GetAll(ctx context.Context) ([]domain.Category, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.GetAll(ctx)
}
