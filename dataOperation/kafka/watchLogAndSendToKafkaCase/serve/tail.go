// 主要包括了两个方法：
// 		TailInit初始化，组装tail配置。Listener
// 		Listener, 保存kafka服务类初始化之后的管道。监听日志文件，如果有新日志，就往管道里发送

// 补充
// github.com/hpcloud/tail包
// tail命令用途是依照要求将指定的文件的最后部分输出到标准设备，通常是终端，通俗讲来，就是把某个档案文件的最后几行显示到终端上，假设该档案有更新，tail会自己主动刷新，确保你看到最新的档案内容 ，在日志收集中可以实时的监测日志的变化。

package serve

import (
	"fmt"
	"github.com/hpcloud/tail"
)

type TailServe struct {
	tails *tail.Tail
}

func (ts *TailServe) TailInit(fileName string) {
	// 初始化tail配置
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}

	// 打开log文件读取数据，放入ts.tails中
	ts.tails, _ = tail.TailFile(fileName, config) // 返回有个tail的结构体，tail结构体的Lines字段封装了拿到的信息
	// if err != nil {
	// 	fmt.Println("tails %s failed,err:%v\n", fileName, err)
	// 	return nil, err
	// }
	fmt.Println("启动，开始监听log文件-----  " + fileName)
}

func (ts *TailServe) Listener(MsgChan chan string) {
	// 循环监听(因为tail可以实现实时监控)
	// 等待ts.tails.Lines通道中line.Text数据，只要有则发送至MsgChan通道
	for {
		line, ok := <-ts.tails.Lines
		if !ok {
			// todo
			fmt.Println("数据接收失败")
			return
		}
		fmt.Println("Line.Text -----", line.Text)
		MsgChan <- line.Text
	}
}
