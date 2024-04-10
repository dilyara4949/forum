package controller

import (
	"forum/internal/domain"
	"forum/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



type SigninController struct {
	SigninUsecase domain.SigninUsecase
	Env          *pkg.Env
}


func (lc *SigninController) Signin(c *gin.Context) {

	var request domain.Signin

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			})
		return
	}

	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Required fields are missing or invalid",})
		return
	}

	user, err := lc.SigninUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: "Неверный логин или пароль",
			})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {

		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
				Message: "Неверный логин или пароль",
			})
		return
	}

	accessToken, err := lc.SigninUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {

		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Message: "Create access token error:"+err.Error(),})
		return
	}
	loginResponse := domain.SigninResponse{
		AccessToken:  accessToken,
	}
	c.JSON(http.StatusOK, loginResponse)
}
