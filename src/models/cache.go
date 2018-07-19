package models

import (
)

// 单例
var m = new(MemCache)
var Caches = m.Init()


type Cache interface{
	Update()
	Delete()
	Get()
}


type MemCache struct {
	Caches map[string]string
}

func (self *MemCache) Init() *MemCache  {
	self.Caches = make(map[string]string)
	return self
}


func (self *MemCache) Find(key string) (string, bool) {
	if value, ok := self.Caches[key]; ok {
		return value, true
	}else{
		return "", false
	}
}

func (self *MemCache) Save(key, value string) {
	self.Caches[key] = value
}
