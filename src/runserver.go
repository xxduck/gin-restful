package main

import (
	"gin-restful/src/databases"
	"gin-restful/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

	{
		user.GET("/", )
	}


	// 定义主页请求方法
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	app.Run()
}
