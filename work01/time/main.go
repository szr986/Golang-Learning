package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	// 时间戳
	fmt.Println(now.Unix())
	// time.Unix()
	ret := time.Unix(1564803667, 0)
	fmt.Println(ret.Year())
	// 格式化时间 把语言中的对象1转化成字符串对象的时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2022-03-01")
	if err != nil {
		fmt.Println("failed err:%v", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// 按照东八区的时区和格式解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load location failed,err:%v", err)
		return
	}
	// 按照指定时区解析时间
	timeObj2, err := time.ParseInLocation("2006-01-02", "2022-03-01", loc)
	if err != nil {
		fmt.Printf("parse time failed,err:%v", err)
		return
	}
	fmt.Println(timeObj2)
	// 时间对象相减
	td := timeObj2.Sub(now)
	fmt.Println(td)
}
