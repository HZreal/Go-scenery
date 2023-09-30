package main

import (
	"encoding/json"
	"fmt"
	"github.com/sony/sonyflake"
)

/**
sony flake
github: https://github.com/sony/sonyflake
*/

var (
	sf *sonyflake.Sonyflake
	// 定义一个全局的 machineID 模拟获取
	// 现实环境中应从 zk 或 etcd 中获取
	sonyMachineID uint16
)

func init() {
	settings := sonyflake.Settings{}
	// settings.MachineID = awsutil.AmazonEC2MachineID

	//
	sf = sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("sonyFlake not created")
	}
}
func main() {
	generateId, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	fmt.Println(generateId)

	body, err := json.Marshal(sonyflake.Decompose(generateId))
	if err != nil {
		return
	}
	fmt.Println(body)
}
