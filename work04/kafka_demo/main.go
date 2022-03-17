package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 基于sarama的第三方库开发kafka client
func main() {
	config := sarama.NewConfig()
	// tailf包的使用
	// 发送完数据需要leader和follower都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的信息将在success channel返回

	// 构建一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
