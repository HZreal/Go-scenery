package go_test

import (
	"fmt"
	"github.com/hpcloud/tail"
	"testing"
)

// 测试案例
func Test_Demo(t *testing.T) {
	filename := `../tailWatchFile.log`
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	// 打开文件开始读取数据
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tails %s failed,err:%v\n", filename, err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	fmt.Println("启动")
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tails file close reopen,filename:$s\n", tails.Filename)
		}
		fmt.Println("msg:", msg.Text)
	}
}
