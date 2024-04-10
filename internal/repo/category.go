package repo

import (
	"context"
	"errors"
	"forum/internal/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)



type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	if _, err := c.db.Exec(ctx, createCategoryQuery, &category.Name); err != nil {
		return nil, err
	}
	return category, nil
}


func (c *categoryRepository) Get(ctx context.Context, name string) (*domain.Category, error) {
	category := domain.Category{}
	if err := c.db.QueryRow(ctx, getCategory, name).Scan(&category.Name); err != nil {
		return nil, err
	}
	return &category, nil
}


func (c *categoryRepository)Delete(ctx context.Context, name string) error {
	r, err := c.db.Exec(ctx, deleteCategory, name)
	if err != nil {
		return err
	}
	if r.RowsAffected() != 1 {
		return errors.New("category does not exist")
	}
	return nil
}
func (c *categoryRepository)GetAll(ctx context.Context) ([]domain.Category, error) {
	rows, err := c.db.Query(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}

	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}
		err = rows.Scan(&category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}





