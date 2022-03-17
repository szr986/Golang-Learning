package kafka

// 初始化kafka连接的一个client

import (
	"fmt"
	"work/logtransfer/es"

	"github.com/Shopify/sarama"
)

type LogData struct {
	Data string `json:"data"`
}

// 初始化kafka，并从kafka取数据
func Init(addrs []string, topic string) error {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	// fmt.Println(partitionList)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println("分区列表", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 直接发给ES
				ld := map[string]interface{}{
					"data": msg.Value,
				}
				// err := json.Unmarshal(msg.Value, ld)
				if err != nil {
					fmt.Println("unmarshal failed err：", err)
					continue
				}
				es.SendToES(topic, ld)
			}
		}(pc)
	}
	return err
}
