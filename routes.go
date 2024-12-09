package main

import (
	"github.com/addione/New/middlewares"
	"github.com/addione/New/src"
	"github.com/gin-gonic/gin"
)

func handleHttp() {
	server := gin.Default()
	addRoutes(server)
	server.Run(":8091")

}

func addRoutes(server *gin.Engine) {
	sdi := src.NewSrcDI()
	sdb := sdi.GetBank()
	uc := sdi.GetUserController()

	loggedInUser := server.Group("/")
	loggedInUser.Use(middlewares.Authenticate)

	server.GET("/users/:id", uc.GetUserById)
	loggedInUser.GET("/users/list", uc.ListUsers)

	server.GET("/clean-db", sdb.CleanDb)

	server.POST("/users", uc.CreateUser)
	server.POST("/login", uc.Login)
	loggedInUser.PUT("/users/:id", uc.UpdateUser)

}
