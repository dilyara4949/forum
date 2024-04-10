package repo

import (
	"context"
	"forum/internal/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(c context.Context, user *domain.User) (*domain.User, error) {
	_, err := u.db.Exec(c, createUserQuery, &user.Email, &user.Password, user.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (u *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.QueryRow(c, getByEmailUserQuery, email).Scan(
		 &user.Username, &user.Email, &user.Password,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(c context.Context, user domain.User) (*domain.User, error) {

	if err := u.db.QueryRow(c, updateUserQuery, &user.Username,).Scan(&user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetAll(c context.Context) ([]domain.User, error) {
	users := []domain.User{}

	rows, err := u.db.Query(c, getAllUsersQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan( &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u *userRepository) GetUserPassword(c context.Context, userID uint) (string, error) {
	var pass string
	if err := u.db.QueryRow(c, getUserPassword, userID).Scan(&pass); err != nil {
		return "", err
	}
	return pass, nil
}