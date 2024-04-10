package route

import (
	"forum/api/controller"
	"forum/internal/repo"
	"forum/internal/usecase"
	"forum/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)



func NewSigninRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	ur := repo.NewUserRepository(db)
	lc := &controller.SigninController{
		SigninUsecase: usecase.NewSigninUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/signin", lc.Signin)
}