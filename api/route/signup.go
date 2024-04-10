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

// func NewSignupRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
// 	ur := repo.NewUserRepository(db)
// 	sc := controller.SignupController{
// 		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
// 		Env:           env,
// 	}
// 	group.POST("/signup", sc.Signup)
// }

func NewSignupRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {

	ur := repo.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}

	group.POST("/signup", sc.Signup)

}