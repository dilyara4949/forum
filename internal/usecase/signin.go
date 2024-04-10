package usecase

import (
	"context"
	"forum/internal/domain"
	"forum/internal/tokenutil"

	"time"
)

type signinUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSigninUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SigninUsecase {
	return &signinUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *signinUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	_, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.userRepository.GetByEmail(c, email)
}

func (lu *signinUsecase) CreateAccessToken(user *domain.User,  secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

// func (lu *signinUsecase) CreateRefreshToken(user *domain.UserProfile, secret string, expiry int) (refreshToken string, err error) {
// 	return tokenutil.CreateRefreshToken(user, secret, expiry)
// }

// func (lu *signinUsecase) ValidateRefreshToken(refreshToken string, secret string) (uint, error) {
// 	return tokenutil.ValidateRefreshToken(refreshToken, secret)
// }


func (lu *signinUsecase) GetUserPassword(c context.Context, userID uint) (string, error) {
	_, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.userRepository.GetUserPassword(c, userID)
}
