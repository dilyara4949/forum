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

func NewPostRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	cc := repo.NewPostRepository(db)
	lc := &controller.PostController{
		PostUsecase: usecase.NewPostUsecase(cc, timeout),
	}

	publicRouter := gin.Group("/api")
	publicRouter.GET("/posts", lc.GetAll)

	protectedRouter := gin.Group("/api")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	protectedRouter.GET("/user/posts", lc.GetOwn)
	// protectedRouter.GET("/posts/:post-id", lc.Get)
	protectedRouter.POST("/posts", lc.Create)
	protectedRouter.DELETE("/posts/:post-id", lc.Delete)

	
}
