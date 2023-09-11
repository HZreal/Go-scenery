// main函数做的有：构建配置结构体，映射配置文件。调用和初始化tail,srama服务。

package main

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"goBasics/middleware/kafka/watchLogAndSendToKafkaCase/serve"
	"os"
	"runtime"
	"strings"
)

type Children struct {
	Name string `ini:"name"`
}
type KafkaConfig struct {
	Address     string `ini:"address"`
	ChannelSize int    `ini:"chan_size"`
}
type TailConfig struct {
	Path     string `ini:"path"`
	Filename string `ini:"fileName"`
	// 如果是结构体，则指明分区名
	Children `ini:"tailfile.children"`
}
type Config struct {
	KafkaConfig `ini:"kafka"`
	TailConfig  `ini:"tailfile"`
}

func CurrentFile() string {
	_, currebtFilePath, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	index := strings.LastIndex(currebtFilePath, string(os.PathSeparator))
	currentDir := currebtFilePath[:index]
	return currentDir
}

func main() {
	// 获取当前文件所在目录
	currentDir := CurrentFile()
	logFilePath := currentDir + "/conf/go-conf.ini"
	fmt.Println("log file path ----", logFilePath)

	// 加载配置
	var cfg = new(Config)
	err := ini.MapTo(cfg, logFilePath)
	if err != nil {
		fmt.Print("failed to map ini file to struct ------ ", err)
	}
	fmt.Println("cfg-----", cfg)

	// 初始化kafka
	ks := &serve.KafukaServe{}
	// 启动kafka消息监听。异步
	ks.InitKafka([]string{cfg.KafkaConfig.Address}, int64(cfg.KafkaConfig.ChannelSize))
	// 关闭主协程时，关闭channel
	defer ks.Destruct()

	// 初始化tail
	ts := &serve.TailServe{}
	ts.TailInit(cfg.TailConfig.Path + "/" + cfg.TailConfig.Filename)
	// 阻塞
	ts.Listener(ks.MsgChan)

}
