package models


import (
	"path"
	
)
// 模拟一个数据库表

// group path permission
var GroupTable = map[string]map[string]int{
	// 普通用户
	"custom":	{
		"/": 15,
		"/user/*" : 0,

	},

	// 管理员
	"root": {
		"/user/*": 15,
	},
}


type Permission interface {
	Read(url string) bool
	Put(url string) bool
	Update(url string) bool
	Delete(url string) bool
}


type Group struct {
	Name string `json:"role"`
}


func (self *Group) Read(url string) bool {
	code := self.find(url)
	if code & 8 == 8 {
		return true
	}else{
		return false
	}
}


func (self *Group) Update(url string) bool {
	code := self.find(url)
	if code & 4 == 4 {
		return true
	}else{
		return false
	}
}

func (self *Group) Delete(url string) bool {
	code := self.find(url)
	if code & 2 == 2 {
		return true
	}else{
		return false
	}
}

func (self *Group) Put(url string) bool {
	code := self.find(url)
	if code & 1 == 1 {
		return true
	}else{
		return false
	}
}


func (self *Group) find(url string) int {
	if value, ok := GroupTable[self.Name]; ok {
		for k, v := range value {
			if ok, _ := path.Match(k, url); ok {
				return v	
			}
		}
	}else{
		return 0
	}

	return 0
}