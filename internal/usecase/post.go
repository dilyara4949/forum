package usecase

import (
	"context"
	"forum/internal/domain"
	"time"
)

type postUsecase struct {
	postRepository domain.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(postRepository domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
		contextTimeout: timeout,
	}
}
func (c *postUsecase) Create(ctx context.Context, post *domain.PostRequest) (*domain.Post, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.postRepository.Create(ctx, post)
}

func (c *postUsecase) GetOwn(ctx context.Context, email string) ([]domain.Post, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.postRepository.GetOwn(ctx, email)
}

func (c *postUsecase) Delete(ctx context.Context, id int) error {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.postRepository.Delete(ctx, id)
}

func (c *postUsecase) GetAll(ctx context.Context) ([]domain.Post, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.postRepository.GetAll(ctx)
}
