package middlewares

import (
	"github.com/gin-gonic/gin"
	"gin-restful/src/models"
)

func UserMiddle() gin.HandlerFunc  {
	return func (c *gin.Context)  {
		user := &models.User{
			Id: 0,
			Name: "root",
			Group: models.Group{
				Name: "root",
			},
		}

		c.Set("user", user)

		c.Next()
	}
}