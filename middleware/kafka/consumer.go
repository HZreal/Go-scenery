package main

import (
	"fmt"
	"github.com/Shopify/sarama" // doc:  https://pkg.go.dev/github.com/Shopify/sarama
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("all partition IDs -----------", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() { // pc.Messages()返回 <-chan *ConsumerMessage
				fmt.Printf("Topic:%s Partition:%d Offset:%d Key:%v Value:%v", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value) // key、value均是[]byte
			}
		}(pc)
	}

	// keep main-thread alive, or the goroutine will be dead with the main-thread's terminal
	a := make(chan bool, 1)
	<-a
}
