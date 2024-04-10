package route

import (
	"database/sql"
	"forum/pkg"
	"time"
)



func NewUserRouter(env *pkg.Env, timeout time.Duration, db *sql.DB) {
// 	ur := repo.NewUserRepository(db)
// 	lc := &controller.UserController{
// 		UserUsecase: usecase.NewUserUsecase(ur, timeout),
// 	}
// 	// group.GET("/user", lc.UserBasic)
// 	// group.POST("/user", lc.UserUpdate)
// 	// group.DELETE("/user", lc.UserDelete)

// 	group.GET("/user/profile", lc.UserProfile)
// 	group.POST("/user/profile", lc.UserUpdate)

}