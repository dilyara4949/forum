package domain

import "context"

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) (*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	Update(c context.Context, user User) (*User, error)
	GetAll(c context.Context) ([]User, error)
	GetUserPassword(c context.Context, userID uint) (string, error)
}

type UserUsecase interface {
	// Create(c context.Context, user *User) (*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	Update(c context.Context, user User) (*User, error)
	GetAll(c context.Context) ([]User, error)
	GetUserPassword(c context.Context, userID uint) (string, error)
}
