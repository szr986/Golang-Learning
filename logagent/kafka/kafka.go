package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

// kafka
var (
	client      sarama.SyncProducer //声明一个全局的连接kafka的生产者client
	logDataChan chan *logData
)

func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	// tailf包的使用
	// 发送完数据需要leader和follower都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的信息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed err:", err)
		return
	}
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道取数据发往kafka
	go SendToKafka()
	return
}

// 往kafka发送消息
func SendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send message failed,err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
	// 构建一个消息
}

// 将消息发往通道
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}
