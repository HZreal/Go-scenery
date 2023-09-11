package main

// Go语言中连接kafka使用第三方库: github.com/Shopify/sarama。
import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
)

// SyncProducer发送消息会阻塞直到被kafka确认。它将消息路由到正确的broker，适当地刷新元数据，并解析错误响应。您必须在生产者上调用Close()来避免泄漏，当它超出范围时可能不会被自动垃圾回收。
// 有两点需要注意: 效率通常低于AsyncProducer，且当消息被确认时提供的实际持久性保证取决于'Producer.RequiredAcks'的配置值。在某些配置中，由SyncProducer确认的消息有时仍然会丢失
func useSyncProducer() {
	// 设置连接配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK应答设置为all，即发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 选出一个随机的partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 创建同步生产者
	syncProducer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer syncProducer.Close()

	// 构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: "web_log",
		Key:   sarama.StringEncoder("firstKey"),
		Value: sarama.StringEncoder("firstValue"),
	}

	// 发送消息
	partition, offset, err := syncProducer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("partition:%v offset:%v\n", partition, offset)
}

// AsyncProducer在通道上接收消息，并在后台尽可能高效地异步生成消息;它是大多数情况下的首选
// AsyncProducer使用非阻塞API发布Kafka消息。它将消息路由到提供的主题分区的合适的broker，适当地刷新元数据，并解析错误响应。必须从Errors()通道读取，否则生产者将死锁。您必须在生产者上调用Close()或AsyncClose()来避免泄漏和消息丢失
// 当它超出作用域时，它不会被自动垃圾回收，缓冲的消息可能不会被刷新

// 这个例子展示了如何使用带有从成功和错误通道读取的独立goroutine的生产者。注意，为了填充成功通道，您必须将config. producer . return . successa设置为true。
func useAsyncProducerWithGoroutines() {
	// 配置
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	// 创建异步生产者
	asyncProducer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, cfg)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		panic(err)
	}

	// Trap SIGINT 触发安全关闭
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	var enqueued, producerSuccesses, producerErrors int

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range asyncProducer.Successes() {
			producerSuccesses++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range asyncProducer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

	//
ProducerLoop:
	for {
		// 创建msg
		message := &sarama.ProducerMessage{
			Topic: "my_topic",
			Value: sarama.StringEncoder("testing 123"),
		}

		select {
		case asyncProducer.Input() <- message:
			enqueued++

		case <-signals:
			asyncProducer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()
	log.Printf("Successfully produced: %d; errors: %d\n", producerSuccesses, producerErrors)

}

// 这个例子展示了如何使用生产者当读取Errors通道时了解任何错误。
func useAsyncProducerWithSelect() {
	// 创建异步生产者
	asyncProducer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		panic(err)
	}
	defer func() {
		if err := asyncProducer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, producerErrors int
	//
ProducerLoop:
	for {
		select {
		case asyncProducer.Input() <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder("testing 123")}:
			enqueued++
		case err := <-asyncProducer.Errors():
			log.Println("Failed to produce message", err)
			producerErrors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, producerErrors)

}

func main() {
	useSyncProducer()
	// useAsyncProducerWithGoroutines()
	useAsyncProducerWithSelect()
}
