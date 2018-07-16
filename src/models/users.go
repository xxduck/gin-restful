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


	user := User{
		Name: c.DefaultQuery("name", "匿名"),
		Passwd: c.DefaultQuery("passwd", "匿名"),
		Email: c.DefaultQuery("email", "15732633601@163.com"),
		Tel: c.DefaultQuery("tel", "15732633601"),
	}
	
	if u, ok := user.MakeToken(); ok {
		c.JSON(200, u)
	}else{
		c.JSON(200, u)
	}
}