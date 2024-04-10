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

func NewCommentRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	cc := repo.NewCommentRepository(db)
	lc := &controller.CommentController{
		CommentUsecase: usecase.NewCommentUsecase(cc, timeout),
	}

	publicRouter := gin.Group("/api")

	publicRouter.GET("/posts/:post-id/comments", lc.Get)


	protectedRouter := gin.Group("/api")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	protectedRouter.POST("/posts/:post-id/comments", lc.Create)
	protectedRouter.DELETE("/posts/:post-id/comment", lc.Delete)

	
}
