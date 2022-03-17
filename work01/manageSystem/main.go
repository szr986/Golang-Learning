package main

import (
	"fmt"
	"os"
)

/*
函数版学生管理系统
能够查看 新增 删除
*/

type student struct {
	id   int64
	name string
}

var allStudent map[int64]*student //变量声明

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)

	fmt.Print("输入学号：")
	fmt.Scanln(&id)
	fmt.Print("输入姓名：")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	allStudent[id] = stu

}

func deleteStudent() {
	var delId int64
	fmt.Println("输入学号")
	fmt.Scanln(&delId)
	delete(allStudent, delId)

}

func showAllStudent(all map[int64]*student) {
	for k, v := range all {
		fmt.Printf("学号：%d  姓名：%s\n", k, v.name)
	}
}

func main() {
	allStudent = make(map[int64]*student, 48)
	//1.打印菜单
	for {
		fmt.Println("学生管理系统")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.exit
	`)
		fmt.Println("选择你的操作:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了：%d\n", choice)
		// 2.等待用户选择
		//3.等待用户选择
		switch choice {
		case 1:
			showAllStudent(allStudent)
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("选错了，重来")
		}
	}
}
