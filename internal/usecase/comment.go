package usecase

import (
	"context"
	"forum/internal/domain"
	"time"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
	contextTimeout time.Duration
}

func NewCommentUsecase(commentRepository domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
	return &commentUsecase {
		commentRepository: commentRepository,
		contextTimeout: timeout,
	}
}


	
	
func (c *commentUsecase) Create(ctx context.Context, Comment *domain.Comment) (*domain.Comment, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.commentRepository.Create(ctx, Comment)
}

func (c *commentUsecase) Get(ctx context.Context, postID int) ([]domain.Comment, error) {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.commentRepository.Get(ctx, postID)
}

func (c *commentUsecase) Delete(ctx context.Context, postID int, email string) error {
	_, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.commentRepository.Delete(ctx, postID, email)
}

