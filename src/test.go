package main

import (
	// "time"
	"fmt"
	"gin-restful/src/models"
)


func main() {
	m := models.Jwt{}
// }
	// fmt.Println(m.Token())
	fmt.Println(m.Checktoken(m.Token()))

	fmt.Println(m.UserId)
}