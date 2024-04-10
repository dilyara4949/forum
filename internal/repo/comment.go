package repo

import (
	"context"
	"errors"
	"forum/internal/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)



type commentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) domain.CommentRepository {
	return &commentRepository{db: db}
}

func (c *commentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	if _, err := c.db.Exec(ctx, createCommentQuery, &comment.Email, &comment.PostID, &comment.Message); err != nil {
		return nil, err
	}
	return comment, nil
}





func (c *commentRepository)	Delete(ctx context.Context, postID int, email string) error {
	r, err := c.db.Exec(ctx, deleteComment, email, postID)
	if err != nil {
		return err
	}
	if r.RowsAffected() != 1 {
		return errors.New("comment does not exist")
	}
	return nil
}

func (c *commentRepository)Get(ctx context.Context, postID int) ([]domain.Comment, error) {
	rows, err := c.db.Query(ctx, getComment, postID)
	if err != nil {
		return nil, err
	}

	comments := []domain.Comment{}

	for rows.Next() {
		comment := domain.Comment{}
		err = rows.Scan(&comment.Email, &comment.PostID, &comment.Message)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}


	
	