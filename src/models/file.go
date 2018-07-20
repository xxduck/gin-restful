package models

import (
	"io/ioutil"
	"strings"
	"os"
)

func write(u *User) {
	tmp := make([]string, 4)
	tmp[0] = u.Name
	tmp[1] = u.Passwd
	tmp[2] = u.Email
	tmp[3] = u.Tel
	s := strings.Join(tmp, " ")
	s = string(append([]byte(s), '\n'))
	f, _ := os.OpenFile("user.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	f.WriteString(s)
}

func read() []string {
	value, _ := ioutil.ReadFile("user.txt")
	list := strings.Split(string(value), "\n")
	return list
}