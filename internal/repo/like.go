package repo

import (
	"context"
	"forum/internal/domain"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type likeRepository struct {
	db *pgxpool.Pool
}

func NewLikeRepository(db *pgxpool.Pool) domain.LikeRepository {
	return &likeRepository{db: db}
}

func (c *likeRepository) Create(ctx context.Context, like *domain.Like, email string) (*domain.Like, error) {

	_, err := c.db.Exec(ctx, createLikeQuery, email, like.PostID)
	if err != nil {

		if strings.Contains(err.Error(), "duplicate") {
			return like, nil
		}
		return nil, err
	}
	return like, nil
}

func (c *likeRepository) Delete(ctx context.Context, postID int, email string) (*domain.Like, error) {
	_, err := c.db.Exec(ctx, deleteLike, email, postID)
	if err != nil {
		return nil, err
	}
	like := &domain.Like{PostID: postID}
	if err := c.db.QueryRow(ctx, getOnePostLike, postID).Scan(&like.Like); err != nil {
		return nil, err
	}
	
	return like, nil
}

func (c *likeRepository) Get(ctx context.Context, postID int) (*domain.Like, error) {
	like := &domain.Like{PostID: postID}
	if err := c.db.QueryRow(ctx, getOnePostLike, postID).Scan(&like.Like); err != nil {
		return nil, err
	}
	
	return like, nil
}
