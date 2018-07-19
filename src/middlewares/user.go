package middlewares

import (
	"github.com/gin-gonic/gin"
	"gin-restful/src/models"
)

func UserMiddle() gin.HandlerFunc  {
	return func (c *gin.Context)  {
		user := new(models.User)
		user = user.Init()
		c.Set("user", user)

		c.Next()
	}
}