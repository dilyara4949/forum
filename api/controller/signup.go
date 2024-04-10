package controller

import (
	"fmt"
	"forum/internal/domain"
	"forum/pkg"
	"log"
	"net/http"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *pkg.Env
}

var (
	verifier = emailverifier.NewVerifier()
)

func (sc *SignupController) Signup(c *gin.Context) {
	
		var request domain.Signup

		err := c.ShouldBind(&request)

		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error to bind request"})
			return
		}
		log.Println("----SignupRequest----: ", request)

		if request.Email == "" || request.Password == "" || request.Username == "" {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "not all fields provided"})
			return
		}

		_, err = verifier.Verify(request.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: fmt.Sprintf("verify email address failed, error is: %v", err)})
			return
		}

		u, _ := sc.SignupUsecase.GetUserByEmail(c, request.Email)

		if u != nil {
			if u.Email == request.Email && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(request.Password)) == nil &&  u.Username == request.Username  {
				accessToken, err := sc.SignupUsecase.CreateAccessToken(u, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
				if err != nil {
					c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "create access token error"})
					return
				}
				signupResponse := domain.SignupResponse{
					AccessToken: accessToken,
				}
				c.JSON(http.StatusOK, signupResponse)
				return
			} else {
				c.JSON(http.StatusConflict, domain.ErrorResponse{
					Message: fmt.Sprintf("User with email %s already exists", request.Email),
				})
				return
			}
		}

		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(request.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		request.Password = string(encryptedPassword)

		user := domain.User{
			Username: request.Username,
			Email:    request.Email,
			Password: request.Password,
		}
		us, err := sc.SignupUsecase.Create(c, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "create user error8:" + err.Error()})
			return
		}

		accessToken, err := sc.SignupUsecase.CreateAccessToken(us, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "create access token error"})
			return
		}

		signupResponse := domain.SignupResponse{
			AccessToken: accessToken,
		}
		c.JSON(http.StatusOK, signupResponse)
	
}
