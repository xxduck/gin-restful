package main

import (
	"gin-restful/src/databases"
	"gin-restful/src/models"
	"github.com/gin-gonic/gin"
	"gin-restful/src/middlewares"
)

func main() {
	databases.DbConnect()

	app := gin.Default()
	// 所有文章
	article := app.Group("article")
	{
		article.GET("/", models.GetAll)
		article.GET("/:id/", models.One)
	}

	// 用户
	user :=  app.Group("user")
	user.Use(middlewares.JwtMiddle())
	// user.Use(middlewares.UserMiddle())

	{
		user.GET("/", models.Index)
		// user.GET("/gettoken/", models.Login)
	}


	app.GET("/login/", models.Login)

	// 定义主页请求方法
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, app.Routes())
	})

	app.Run()
}

