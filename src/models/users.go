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
	Role [3]Group
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

func (self *User) PermissionRead(url string) bool {
	
	xrange := len(self.Role)
	for i := 0; i < xrange; i++ {
		// 结构体不为空
		if self.Role[i].Name != "" && self.Role[i].Read(url) {
			return true
		}
	}
	return false
}

func (self *User) PermissionPut(url string) bool {
	
	xrange := len(self.Role)
	for i := 0; i < xrange; i++ {
		// 结构体不为空
		if self.Role[i].Name != "" && self.Role[i].Put(url) {
			return true
		}
	}
	return false
}

func (self *User) PermissionDelete(url string) bool {
	
	xrange := len(self.Role)
	for i := 0; i < xrange; i++ {
		// 结构体不为空
		if self.Role[i].Name != "" && self.Role[i].Delete(url) {
			return true
		}
	}
	return false
}

func (self *User) PermissionUpdate(url string) bool {
	
	xrange := len(self.Role)
	for i := 0; i < xrange; i++ {
		// 结构体不为空
		if self.Role[i].Name != "" && self.Role[i].Update(url) {
			return true
		}
	}
	return false
}

// view
func Index(c *gin.Context) {
	if v, ok := c.Get("user"); ok {
		if value, ok := v.(*User); ok {
			url := c.Request.URL.String()
			switch  {
			case value.PermissionRead(url):
				c.String(200, "可读")
				
				fallthrough
			
			case value.PermissionPut(url):
				c.String(200, "可增")
				fallthrough
			
			case value.PermissionUpdate(url):
				c.String(200, "可改")
				fallthrough
			
			case value.PermissionDelete(url):
				c.String(200, "可珊")
			default:
				c.String(200, "啥都干不了")
			}

		}else{
			c.String(200, "查询用户状态失败")
		}
	}else{
		c.String(200, "查询用户状态失败")
	}

	
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
	
