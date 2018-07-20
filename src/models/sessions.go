package models

// session保存的信息  用户基本信息（不包含保密信息） 用户权限
import (
	"github.com/gin-gonic/gin"
)



type Sessions interface{
	
}


type Session struct {
	User
}

var store = map[string]*User{}


func (self *Session) Save(c *gin.Context) {
	// 一边保存到服务器
	// 一边将唯一id写入cookies
	id := "3071497267061195602"   // 实际中需要根据当前时间生成唯一不同的id

	c.SetCookie("sessionid", id, 300, "/", "127.0.0.1", false, true)
	store[id] = &self.User
}


func (self *Session) Delete(id string) {
	// 删除session
	delete(store, id)
}


func (self *Session) Find(id string) (*User, bool) {
	// 删除session
	if value, ok := store[id]; ok {
		return value, true
	}else{
		return nil, false
	}
}