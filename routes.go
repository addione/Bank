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
	sdb := src.NewSrcDI().GetBank()
	su := src.NewSrcDI().GetUserController()

	loggedInUser := server.Group("/")
	loggedInUser.Use(middlewares.Authenticate)

	server.GET("/users/:id", su.GetUserById)
	loggedInUser.GET("/users/list", su.ListUsers)

	server.GET("/clean-db", sdb.CleanDb)

	server.POST("/users", su.CreateUser)
	server.POST("/login", su.Login)
	loggedInUser.PUT("/users/:id", su.UpdateUser)

}
