package middlewares

import (
	"bytes"
	"gin-restful/src/models"
	"github.com/gin-gonic/gin"
)

type Mycontext struct {
	gin.ResponseWriter
	Body bytes.Buffer
}

func (self *Mycontext) Write(body []byte) (int, error) {
	n, err := self.ResponseWriter.Write(body)
	if err != nil {
		return n, err
	}else{
		self.Body.Write(body)
	}
	return n, err
}

func CacheMiddle() gin.HandlerFunc {
	return func (c *gin.Context)  {
		url := c.Request.URL.String()
		caches := models.Caches

		if value, ok := caches.Find(url); ok {
			c.Writer.WriteString(value)
			c.Abort()
		}else{
				// 在之前替换c.Writer
			tmp := c.Writer

			rw := &Mycontext{
				ResponseWriter: c.Writer,
			}
			c.Writer = rw

			// 假函数去执行
			c.Next()

			caches.Save(url, rw.Body.String())
			// 执行完毕后换回原对象
			c.Writer = tmp

		}
		

		}
}