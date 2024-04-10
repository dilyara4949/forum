package domain

import "context"

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type SigninResponse struct {
	AccessToken  string `json:"access_token"`
}


type SigninUsecase interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	GetUserPassword(c context.Context, userID uint) (string, error)
}