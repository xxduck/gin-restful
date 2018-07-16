package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gin-restful/src/databases"
	"strconv"
)


type Article struct {
	Id int
	Article_id int
	Author_id interface{}
	Author interface{}
	Title interface{}
	Context interface{}
	Ptime interface{}
	Ptype interface{}
	Views interface{}
	Replys interface{}
	Status_Code int
}


func GetAll(c *gin.Context)  {
	// c.JSON(200, gin.H{
	// 	"status": "ok",
	// })
	page := c.DefaultQuery("page", "1")
	page_num, _ := strconv.Atoi(page)
	
	if page_num <= 0 {
		page_num = 1
	}

	page_num = (page_num - 1) * 10

	docker := []*Article{}

	conn := databases.DbConnect()
	sql := fmt.Sprintf("select * from autohome_article.article limit 10 offset %d", page_num)

	value, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for value.Next() {
		a := &Article{}
		err := value.Scan(&a.Id, &a.Article_id, &a.Author_id, &a.Author, &a.Title, &a.Context, &a.Ptime, &a.Ptype, &a.Views, &a.Replys, &a.Status_Code)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, err)
		}
		docker = append(docker, a)
		
	}
	c.JSON(200, docker)


}


func One(c *gin.Context)  {
	id := c.Param("id")
	conn := databases.DbConnect()
	sql := fmt.Sprintf("select * from autohome_article.article where id = %s", id)

	value, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for value.Next() {
		a := &Article{}
		err := value.Scan(&a.Id, &a.Article_id, &a.Author_id, &a.Author, &a.Title, &a.Context, &a.Ptime, &a.Ptype, &a.Views, &a.Replys, &a.Status_Code)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, err)
		}
		c.JSON(200, a)
		
	}
}
