package route

import (
	"forum/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Setup(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	
	NewSignupRouter(env, timeout, db, publicRouter)
	NewSigninRouter(env, timeout, db, publicRouter)

	NewCategoryRouter(env, timeout, db, gin)
	NewPostRouter(env, timeout, db, gin)
	NewCommentRouter(env, timeout, db, gin)
	NewLikeRouter(env, timeout, db, gin)
}
