package main

import (
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

	server.GET("/user/:id", su.GetUserById)
	server.GET("/list-users", su.ListUsers)

	server.GET("/clean-db", sdb.CleanDb)

	server.POST("/new-user", su.CreateUser)
	server.POST("/login", su.Login)

	server.PUT("/user/:id", su.UpdateUser)

}
