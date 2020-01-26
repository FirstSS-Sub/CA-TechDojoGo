package controllers

import (
	mysql "CA-TechDojoGo/db"
	"CA-TechDojoGo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var user models.UserCreateRequest
	var err error

	c.BindJSON(&user) //userに受け取ったJsonを代入

	_, err = mysql.DB.Exec("INSERT INTO i_user (name) VALUES (?)", user.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Inserted into i_user, name=" + user.UserName)
	return
}
