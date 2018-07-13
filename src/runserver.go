package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义所有链接
var APIS = make(map[string]string)



func main() {
	app := gin.Default()

	// 定义所有链接
	APIS["index"] = "/"


	
	// 定义主页请求方法
	app.GET(APIS["index"], func (c *gin.Context)  {
		c.JSON(http.StatusOK, APIS)
	})

	
	app.Run()
}
