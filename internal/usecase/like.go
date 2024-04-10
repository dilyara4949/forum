package usecase

import (
	"context"
	"forum/internal/domain"
	"time"
)

type likeUsecase struct {
	likeRepository domain.LikeRepository
	contextTimeout time.Duration
}

func NewLikeUsecase(likeRepository domain.LikeRepository, timeout time.Duration) domain.LikeUsecase {
	return &likeUsecase {
		likeRepository: likeRepository,
		contextTimeout: timeout,
	}
}


	
	
func (c *likeUsecase) Create(ctx context.Context, like *domain.Like, email string) (*domain.Like, error){
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.likeRepository.Create(ctx, like, email)
}

func (c *likeUsecase) Get(ctx context.Context, postID int) (*domain.Like, error){
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.likeRepository.Get(ctx, postID)
}

func (c *likeUsecase) Delete(ctx context.Context, postID int, email string) (*domain.Like, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.likeRepository.Delete(ctx, postID, email)
}


