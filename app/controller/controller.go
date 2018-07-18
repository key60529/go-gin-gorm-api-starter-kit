package controller

import (
	"fmt"
	"kayl-space_backend/app/database"
	"kayl-space_backend/app/models"

	"github.com/gin-gonic/gin"
)

var tx = database.GetDB()

// GetUsers get all the user record form db
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := tx.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// GetUser get single user record form db
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
