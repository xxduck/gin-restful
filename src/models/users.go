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
			c.JSON(200, v)
			return
		}
	}
	c.JSON(200, gin.H{
		"status": 200,
	})
}


func Login(c *gin.Context)  {
	name := c.Query("name")
	passwd := c.Query("passwd")

	if name == "xiaofang" && passwd == "123456" {
		jwt := new(Jwt)
		jwt = jwt.Init()
		jwt.UserId = 10

		token := jwt.Token()
		c.Header("Authorization", token)
		c.JSON(200, gin.H{
			"token": token,
		})
		return
	}else{
		c.JSON(200, gin.H{
			"status": name,
			"pwd": passwd,
		})
	}
	
	
}
	
