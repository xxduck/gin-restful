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
				c.Set("user", &m.User)
				c.Next()
			}else{
				c.AbortWithStatusJSON(200, gin.H{
				"reson": "权限不允许",
			})
			}

		}else{
			// 无token
			user := &models.User{
				Id: 0,
				Name: "root",
				Role: [3]models.Group{
					// 默认状态
					models.Group{Name: "custom"},
					// models.Group{Name: "root"},
				}}
			
	
			c.Set("user", user)
	
			c.Next()

		}

		
	}
}