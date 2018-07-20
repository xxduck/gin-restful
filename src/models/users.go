package models

import (
	"strings"

	"github.com/gin-gonic/gin"
)


type User struct {
	Id int `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Passwd string	`json:"passwd,omitempty" form:"passwd" binding:"required"`
	Email string	`json:"email,omitempty" form:"email" binding:"required"`
	Tel string		`json:"tel,omitempty" form:"tel" binding:"required"`
	Role [3]Group
}

// 初始化默认user（即匿名用户）
func (self *User) Init() *User {

	self.Name = "anonymous"
	self.Id = 0
	self.Passwd = ""
	self.Email = ""
	self.Tel = ""
	self.Role = [3]Group{
		Group{
			Name: "anonymous",
		},
	}
	
	return self
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
func UserInfo(c *gin.Context)  {
	user, _ := c.Get("user")
	if u, ok := user.(*User); ok {
		c.JSON(200, u)
	}else{
		c.String(200, "你好")
	}
}


func Login(c *gin.Context)  {
	name := c.Query("name")
	passwd := c.Query("passwd")
	infos := read()
	index := 0
	for K, info := range infos {
		for _, i := range strings.Split(info, " ") {
			if i == name {
				index = K
				goto LITE
			}
		}
	}

	LITE:
		name_pd := infos[index]
		np := strings.Split(name_pd, " ")
		if name == np[0]  && passwd == np[1] {
			c.JSON(200, gin.H{
				"status": "登录成功",
			})
		}else{
			c.JSON(200, gin.H{
				"status": "用户名或密码错误",
			})
		}
	
	
}


func Logon(c *gin.Context)  {
	// 注册
	if user, ok :=  c.Get("user"); ok {
		// 已经;
		if value, ok := user.(*User); ok && value.Name == "anonymous" {
			u := &User{}
			err := c.Bind(u)
			// g := new(Group)
			// g.Name = "custom"
			// u.Role = [3]Group{*g}
			if err != nil {
				c.JSON(200, gin.H{
					"error": err,
				})
			}else{
				write(u)
				c.JSON(200, u)
			}
			

		}else{
			// 已经登录了，无需再注册
			c.Redirect(302, "http://127.0.0.1:8080")

		}
	}
}
	
