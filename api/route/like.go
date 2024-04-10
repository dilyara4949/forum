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

func NewLikeRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	cc := repo.NewLikeRepository(db)
	lc := &controller.LikeController{
		LikeUsecase: usecase.NewLikeUsecase(cc, timeout),
	}

	publicRouter := gin.Group("/api")

	publicRouter.GET("/posts/:post-id/likes", lc.Get)


	protectedRouter := gin.Group("/api")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	protectedRouter.POST("/posts/:post-id/likes", lc.Create)
	protectedRouter.DELETE("/posts/:post-id/likes", lc.Delete)

	
}
