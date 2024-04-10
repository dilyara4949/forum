package usecase

import (
	"context"
	"forum/internal/domain"
	"time"
)



type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}


func (uu *userUsecase) Update(c context.Context, user domain.User) (*domain.User, error) {
	_, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.Update(c, user)
}


func (su *userUsecase) GetByEmail(c context.Context, email string) (*domain.User, error) {
	_, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(c, email)
}


func (pu *userUsecase) GetAll(c context.Context) ([]domain.User, error) {
	_, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.GetAll(c)
}

func (pu *userUsecase) GetUserPassword(c context.Context, userid uint) (string, error) {
	_, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.GetUserPassword(c, userid)
}