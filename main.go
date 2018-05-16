package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomoyamatsuyama/db-sample/handler"
)

func main() {
	server := gin.Default()

	server.POST("/users", handler.SignUp)

	server.POST("/login", handler.Login)

	server.Run(":3000")
}
