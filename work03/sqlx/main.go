package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id   int
	Name string
	Age  int
}

var db *sqlx.DB

func initDB() (err error) {
	// 连数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn) //sqlx会帮助ping一下
	if err != nil {                      //dsn格式不正确时报错
		// fmt.Printf("dsn invalid: %v\n", err)
		return
	}
	if err != nil {
		// fmt.Printf("open database: %v\n", err)
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB error: %v", err)
	}
	sqlStr1 := `select id,name,age from user where id=1`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select err: %v", err)
	}
	fmt.Printf("userlist:%#v\n", userList)
}
