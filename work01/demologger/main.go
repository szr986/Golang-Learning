package main

import (
	"work/work01/mylogger"
)

func main() {
	i := mylogger.NewFileLogger("info", "./", "zhoulinwan.log", 10*1024)
	for {
		id := 10010
		name := "lixiang"
		i.Debug("这是一条Debug,id:%v,name:%v", id, name)
		i.Info("这是一条Info")
		i.Warning("这是一条warning")
		i.Error("这是一条Error")
		i.Fatal("这是一条Fatal")
		// time.Sleep(500 * time.Millisecond)
	}
}
