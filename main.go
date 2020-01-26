package main

import (
	"CA-TechDojoGo/controllers"
	mysql "CA-TechDojoGo/db"
	_ "CA-TechDojoGo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	err := mysql.Initialize()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	//controllers
	//あるpackageで定義された識別子を外部参照させる場合は、大文字で始める必要がある
	r.POST("/user/create", controllers.UserCreate)
	r.Run(":8080")
}
