package middlewares

import (
	"gin-restful/src/models"
	"github.com/gin-gonic/gin"
)


func Session() gin.HandlerFunc {
	return func (c *gin.Context)  {
		value, err := c.Cookie("sessionid")
		if err !=  nil {
			// 没有sessionid
			user := new(models.User)
			user = user.Init()
			c.Set("user", user)
		}else{
			session := new(models.Session)
			if user, ok := session.Find(value); ok {
				// 找到session
				c.Set("user", user)
			}else{
				user := new(models.User)
				user = user.Init()
				c.Set("user", user)
			}
		}

		c.Next()
	}
}


