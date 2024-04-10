package repo

import (
	"context"
	"forum/internal/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) domain.PostRepository {
	return &postRepository{db: db}
}

func (c *postRepository) Create(ctx context.Context, Post *domain.PostRequest) (*domain.Post, error) {
	post := domain.Post{}
	if err := c.db.QueryRow(ctx, createPostQuery, &Post.Email, &Post.Title, &Post.Body, &Post.Category).Scan(&post.ID); err != nil {
		return nil, err
	}
	post.Email = Post.Email
	post.Body = Post.Body
	post.Title = Post.Title
	post.Category = Post.Category

	return &post, nil
}

func (c *postRepository) GetOwn(ctx context.Context, email string) ([]domain.Post, error) {
	rows, err := c.db.Query(ctx, getPost, email)
	if err != nil {
		return nil, err
	}

	posts := make([]domain.Post, 0)

	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.ID, &post.Email, &post.Title, &post.Body, &post.Category)
		if err != nil {
			return nil, err
		}

		if err := c.db.QueryRow(ctx, getOnePostLike,post.ID).Scan(&post.Likes); err != nil {
			return nil, err
		}

		post.Comments = make([]domain.Comment, 0)
		rows2, err := c.db.Query(ctx, getComment, post.ID)
		if err != nil {
			return nil, err
		}

		for rows2.Next() {
			comment := domain.Comment{}
			err = rows2.Scan(&comment.Email, &comment.PostID, &comment.Message)
			if err != nil {
				return nil, err
			}
			post.Comments = append(post.Comments, comment)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (c *postRepository) Delete(ctx context.Context, id int) error {
	if _, err := c.db.Exec(ctx, deletePost, id); err != nil {
		return err
	}
	return nil
}
func (c *postRepository) GetAll(ctx context.Context) ([]domain.Post, error) {
	rows, err := c.db.Query(ctx, getAllPost)
	if err != nil {
		return nil, err
	}

	posts := make([]domain.Post, 0)

	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.ID, &post.Email, &post.Title, &post.Body, &post.Category)
		if err != nil {
			return nil, err
		}
		if err := c.db.QueryRow(ctx, getOnePostLike, &post.Email, post.ID).Scan(&post.Likes); err != nil {
			return nil, err
		}

		post.Comments = make([]domain.Comment, 0)

		rows2, err := c.db.Query(ctx, getComment, post.ID)
		if err != nil {
			return nil, err
		}

		for rows2.Next() {
			comment := domain.Comment{}
			err = rows2.Scan(&comment.Email, &comment.PostID, &comment.Message)
			if err != nil {
				return nil, err
			}
			post.Comments = append(post.Comments, comment)
		}

		posts = append(posts, post)
	}
	return posts, nil
}
