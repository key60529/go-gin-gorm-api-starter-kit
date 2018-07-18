package main

import (
	"kayl-space_backend/app/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/", controller.GetUsers)
		v1.GET("/user/:id", controller.GetUser)
	}
	router.Run(":8080")
}
