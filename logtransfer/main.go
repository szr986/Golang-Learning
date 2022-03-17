package main

import (
	"fmt"
	"work/logtransfer/conf"
	"work/logtransfer/es"
	"work/logtransfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	var cfg = new(conf.LogTransferConf)
	// 0.加载配置文件
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("init config failed:", err)
	}
	fmt.Println(cfg)
	// ES初始化
	// 1.1 初始化一个ES连接的client
	// 1.2	对外提供一个发送函数
	es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Println("init ES failed:", err)
		return
	}
	fmt.Println("init es success!")
	// 1.初始化
	// 1.1 连接kafka，创建分区消费者
	// 1.2 每个分区的消费者分别取出数据，通过SendToES将数据发往es
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Println("init kafka consumer failed:", err)
		return
	}
	select {}
	// 2.初始化ES

	// 2.从kafka取日志
}
