package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

var db *sql.DB //是一个连接池对象

func initDB() (err error) {
	// 连数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码
	if err != nil {                  //dsn格式不正确时报错
		// fmt.Printf("dsn invalid: %v\n", err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		// fmt.Printf("open database: %v\n", err)
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询
func queryOne(id int) {
	var u1 user
	// 1.查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?`
	// 2.执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age) //从连接池中拿出一个连接去数据库查询单条记录
	// 3.拿到结果赋给结构体
	//必须调用Scan方法，该方法中会释放连接
	fmt.Printf("u1:%#v", u1)
}

func queryMore(n int) {
	// 1.写查询sql语句
	sqlStr := `select id,name,age from user where id>?;`
	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed,err:%v\n", sqlStr, err)
		return
	}
	// 3.一定要关闭rows
	defer rows.Close()
	// 4.循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Println("scan failed", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入数据
func insert() {
	// 1.写SQL语句
	sqlStr := `insert into user(name,age) values("hahaGe",89);`
	// 2.exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	// 如果是插入数据，能拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId error:", err)
		return
	}
	fmt.Println(id)
}

func updateRow(newAge, id int) {
	sqlStr := `update user set age=? where id>?;`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get Id error:", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

func delectRow(id int) {
	sqlStr := `delete from user where id=?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete Id error:", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name,age) values (?,?)`
	stmt, err := db.Prepare(sqlStr) //把sql语句发给mysql先预处理
	if err != nil {
		fmt.Printf("prepare failed", err)
		return

	}
	defer stmt.Close()
	// 后续拿到stmt去执行一些操作
	var m = map[string]int{
		"qiqi":    3000,
		"hutao":   17,
		"yunjing": 18,
		"ganyu":   2000,
	}

	for k, v := range m {
		stmt.Exec(k, v)
	}
}

// 事务操作
func transactionDemo() {
	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Begin failed err:", err)
		return
	}
	// 2.执行多个SQL操作
	sqlStr1 := `update user set age=age-2 where id=1;`
	sqlStr2 := `update user set age=age+2 where id=2;`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 回滚
		tx.Rollback()
		fmt.Println("执行sql1出错，回滚")
		return
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 回滚
		tx.Rollback()
		fmt.Println("执行sql2出错，回滚")
		return
	}
	// 上面两步SQL都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 回滚
		tx.Rollback()
		fmt.Println("提交出错，回滚")
		return
	}
	fmt.Println("事务成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB error: %v", err)
	}
	// queryMore(0)
	// insert()
	// updateRow(829, 0)
	// delectRow(2)
	// prepareInsert()
	transactionDemo()
}
