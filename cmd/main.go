package main

import (
	"fmt"
	"forum/api/route"
	"forum/pkg"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := pkg.App()
	if err != nil {
		log.Fatal(err)
	}

	env := app.Env
	db := app.Pql
	defer app.CloseDBConnection()

	ginRouter := gin.Default()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	route.Setup(env, timeout, db, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", env.PORT))
}
