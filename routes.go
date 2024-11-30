package main

import (
	"github.com/addione/New/src"
	"github.com/gin-gonic/gin"
)

func handleHttp() {
	server := gin.Default()
	sdb := src.NewSrcDI().GetBank()
	su := src.NewSrcDI().GetUserController()

	server.POST("/new-user", su.CreateUser)
	server.GET("/user/:id", su.GetUserById)
	server.GET("/clean-db", sdb.CleanDb)
	server.GET("/list-users", su.ListUsers)

	server.PUT("/user/:id", su.UpdateUser)

	server.Run(":8091")

}
