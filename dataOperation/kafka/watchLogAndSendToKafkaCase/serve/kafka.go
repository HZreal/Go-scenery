// 有3个方法 ：
// 		InitKafka，组装配置项以及初始化接收消息的管道，
// 		Listener, 监听管道消息，收到消息后，将消息组装，发送到kafka
// 		Destruct, 关闭管道

package serve

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type KafukaServe struct {
	MsgChan chan string
	//err         error
}

func (ks *KafukaServe) InitKafka(addr []string, chanSize int64) {

	// 初始化kafka连接配置
	config := sarama.NewConfig()
	// 1. 初始化生产者配置
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 选择分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的信息
	config.Producer.Return.Successes = true

	// 初始化msg通道
	ks.MsgChan = make(chan string, chanSize)

	// 异步连接kafka并监听通道
	go ks.Listener(addr, chanSize, config)

}

func (ks *KafukaServe) Listener(addr []string, chanSize int64, config *sarama.Config) {
	//  连接kafka
	var kafkaClient, _ = sarama.NewSyncProducer(addr, config)
	defer kafkaClient.Close()

	// 循环监听等待通道MsgChan中的数据，只要有则向kafka发送
	for {
		select {
		case content := <-ks.MsgChan:
			//
			msg := &sarama.ProducerMessage{
				Topic: "weblog",
				Value: sarama.StringEncoder(content),
			}
			partition, offset, err := kafkaClient.SendMessage(msg)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("分区，偏移量分别为：", partition, offset)
		}

	}
}

func (ks *KafukaServe) Destruct() {
	close(ks.MsgChan)
}
