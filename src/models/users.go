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
	Group
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
	// if value, ok := c.Get("JWT"); ok {
	// 	if v, ok := value.(Jwt); ok{
	// 		c.JSON(200, v)
	// 		return
	// 	}
	// }
	// c.JSON(200, gin.H{
	// 	"status": 200,
	// })

	if value, ok := c.Get("user"); ok {
		if v, ok := value.(*User); ok {
			
			con := [4]string{}
			func (per Permission)  {

				
				p := "增"
				d := "删"
				u := "改"
				r := "查"

				if per.Read(c.Request.URL.String()) {
					con[0] = p
				}
				if per.Put(c.Request.URL.String()) {
					con[1] = d
				}
				if per.Delete(c.Request.URL.String()) {
					con[2] = u
				}
				if per.Update(c.Request.URL.String()) {
					con[3] = r
				}
			}(v)
			c.JSON(200, con)
			return
			}}

	c.JSON(200, gin.H{
		"status": "不可以读取",
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
	
