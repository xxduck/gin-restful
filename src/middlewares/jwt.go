package middlewares

import (

	"github.com/gin-gonic/gin"
	"gin-restful/src/models"
)

func JwtMiddle() gin.HandlerFunc  {
	return func (c *gin.Context)  {

		// 头部有jwt
		auth := c.Request.Header.Get("Authorization")
		if auth != "" {
			// 有token
			m := models.Jwt{}

			if m.Checktoken(auth) {
				c.Set("JWT", m)
				c.Next()
			}

		}else{
			// 无token
			c.AbortWithStatusJSON(200, gin.H{
				"reson": "权限不允许",
			})
		}

		
	}
}