package controllers

import (
	"CA-TechDojoGo/models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var user models.UserCreateRequest
	c.BindJSON(&user) //userに受け取ったJsonを代入
}
