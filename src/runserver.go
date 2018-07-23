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
	app.Use(middlewares.Session())
	
	// 所有文章
	article := app.Group("article")
	article.Use(middlewares.CacheMiddle())
	{
		article.GET("/", models.GetAll)
		article.GET("/:id/", models.One)
	}

	// 用户
	user :=  app.Group("user")
	// user.Use(middlewares.JwtMiddle())
	// user.Use(middlewares.Session())

	{
		user.GET("/", models.UserInfo)
	}


	app.GET("/login/", models.Login)
	app.GET("/logon/", models.Logon)

	// 定义主页请求方法
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, app.Routes())
	})

	app.Run("127.0.0.1:8080")
}
