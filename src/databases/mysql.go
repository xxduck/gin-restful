package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/gin-gonic/gin"
	"os"
)

func DbConnect() *sql.DB {
	// export MYSQLPWD=""
	// export MYSQLIP=""
	// export MYSQLPORT="3306"
	pwd, _ := os.LookupEnv("MYSQLPWD")
	ip, _ := os.LookupEnv("MYSQLIP")
	port, _ := os.LookupEnv("MYSQLPORT")

	addr := fmt.Sprintf("dev:%s@tcp(%s:%s)/?charset=utf8mb4", pwd, ip, port)
	conn, err := sql.Open("mysql", addr)

	if err != nil {
		fmt.Println(err)
		return nil
	}else{
		return conn
	}

	// value, err := conn.Query("show databases;")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	var Database string
	// 	for value.Next() {
	// 		value.Scan(&Database)
	// 		fmt.Println(Database)

	// 	}
	// }
	// return conn
}
