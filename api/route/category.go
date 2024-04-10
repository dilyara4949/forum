package route

import (
	"forum/api/controller"
	"forum/internal/repo"
	"forum/internal/usecase"
	"forum/pkg"
	"forum/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewCategoryRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	cc := repo.NewCategoryRepository(db)
	lc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(cc, timeout),
	}

	publicRouter := gin.Group("/api")

	publicRouter.GET("/categories", lc.GetAll)
	publicRouter.GET("/categories/:name", lc.Get)


	protectedRouter := gin.Group("/api")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	protectedRouter.POST("/categories", lc.Create)
	protectedRouter.DELETE("/categories/:name", lc.Delete)

	
}
