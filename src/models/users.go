package models

import (
	"github.com/gin-gonic/gin"
)


type User struct {
	Id int
	Name  string
	Passwd string
	Email string
	Tel string
	Token string
}


func (self *User) MakeToken() (*User, bool) {
	if self.check() {
		self.Token = "hello"
		return self, true
	}else{
		return self, false
	}
}


func (self *User) check() bool {
	if self.Name == "xiaofang" && self.Passwd == "123456" {
		return true
	}else{
		return false
	}
}

// view
func Index(c *gin.Context) {
	if value, ok := c.Get("JWT"); ok {
		if v, ok := value.(Jwt); ok{
			c.JSON(200, gin.H{
				"yourid": v.UserId,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"status": 200,
	})
}


func Login(c *gin.Context)  {
	
	
	c.JSON(200, gin.H{
		"status": 200,
	})
}
	
