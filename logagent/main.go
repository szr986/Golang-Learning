package main

import (
	"fmt"
	"sync"
	"time"
	"work/logagent/conf"
	"work/logagent/etcd"
	"work/logagent/kafka"
	"work/logagent/taillog"
	"work/logagent/utils"

	"gopkg.in/ini.v1"
)

var cfg = new(conf.AppConf)

// func run() {
// 	// 1.读取日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadLog():
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// 	// 2.发送到kafka
// }

// logAgent入口
func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini error:", err)
	}
	fmt.Println(cfg.KafkaConf.Topic)
	// fmt.Println(cfg.TaillogConf.Filename)
	fmt.Println(cfg.EtcdConf.Key)
	// 1.初始kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Maxsize)
	if err != nil {
		fmt.Println("init kafka failed,err:", err)
		return
	}
	// 2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed,err:", err)
		return
	}
	fmt.Println("etcd init success")
	// 实现每个logagent都能获取自己的配置，从IP地址拉取
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 2.1 从etcd中拉取日志收集项的配置信息
	fmt.Println("获取Key:", etcdConfKey)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println("get conf failed,err:", err)
		return
	}
	fmt.Println("get conf from etcd success:", &logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v  value:%v\n", index, value)
	}
	// 3.收集日志发往kafka
	// 3.1 循环每一个日志收集项，创建一个tailObj
	// 3.2 发往kafka

	taillog.Init(logEntryConf)

	// 2.2 派一个哨兵去监视日志收集项的变化
	newconfchan := taillog.NewConfChan() //从taillog中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newconfchan) //哨兵发现最新的配置会通知上面那个通道
	wg.Wait()

	// // 2.打开日志文件准备收集
	// err = taillog.Init(cfg.TaillogConf.Filename)
	// if err != nil {
	// 	fmt.Println("init taillog failed,err:", err)
	// 	return
	// }

}
