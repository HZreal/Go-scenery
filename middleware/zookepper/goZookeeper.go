package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// 获取一个zk连接
var zkList = []string{"localhost:2181"}

// 连接到zk集群
// var zkList = []string{"localhost:21811", "localhost:21812", "localhost:21813", "localhost:21814"}
func getConnect() (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// 增
func zkCreate() {
	conn := getConnect()
	var data = []byte("test value")

	// flags有4种取值：
	// flags=0                                // 创建永久节点，除非手动删除，默认
	// flags = zk.FlagEphemeral = 1           // 创建临时节点，session断开则该节点也被删除
	// flags = zk.FlagSequence  = 2           // 会自动在节点名称后面添加序号
	// flags = 3                              // Ephemeral和Sequence即1和2的组合， 短暂且自动添加序号
	var flags int32 = 0
	// var flags int32 = zk.FlagEphemeral
	// var flags int32 = zk.FlagSequence
	// var flags int32 = 3

	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)

	// 创建数据节点
	str, err := conn.Create("/test", data, flags, acls)
	// str, err := conn.Create("/test", data, zk.FlagEphemeral, acls)     // 测试临时节点
	if err != nil {
		fmt.Printf("创建失败: %v\n", err)
		return
	}
	fmt.Println(str)

}

// 查
func zkGet() {
	conn := getConnect()

	// 获取所有节点
	children, _, err := conn.Children("/")
	fmt.Println("all children ------", children)

	// 获取某个path的数据
	byteData, stat, err := conn.Get("/test")
	if err != nil {
		fmt.Printf("查询%s失败, err: %v\n", err)
		return
	}
	fmt.Println("get stat info -------", stat) // stat结构体
	fmt.Println("get data -------", string(byteData))
}

// 删改与增不同在于其函数中的version参数,其中version是用于 CAS 支持
// 可以通过此种方式保证原子性

// 改
func zkModify() {
	conn := getConnect()

	new_data := []byte("hello zookeeper")
	_, stat1, _ := conn.Get("/test")
	fmt.Println("stat1 ----", stat1)
	stat2, err := conn.Set("/test", new_data, stat1.Version)
	if err != nil {
		fmt.Printf("数据修改失败: %v\n", err)
		return
	}
	fmt.Println("stat2 ----", stat2)

}

// 删
func zkDelete() {
	conn := getConnect()

	_, stat, _ := conn.Get("/test")
	fmt.Println("stat1 ----", stat)
	err := conn.Delete("/test", stat.Version)
	if err != nil {
		fmt.Printf("数据删除失败: %v\n", err)
		return
	}
}

// zk watch 回调函数
func zkEventCallback(event zk.Event) {
	// 事件类型
	// zk.EventNodeCreated = 1
	// zk.EventNodeDeleted = 2
	// zk.EventNodeDataChanged = 3

	fmt.Println("----------zkEventCallback start-----------------")
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
	fmt.Println("----------zkEventCallback end-----------------\n")
}

// 全局监听
// Java API中是通过Watcher实现的，在go-zookeeper中则是通过Event。道理都是一样的
func zkWatchGlobal() {
	// 创建监听的option，用于初始化zk
	eventCallbackOption := zk.WithEventCallback(zkEventCallback)
	conn, _, err := zk.Connect(zkList, 5*time.Second, eventCallbackOption)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 开始监听path
	conn.ExistsW("/test")

	// 手动触发zk数据操作
	zkModify()

}

// 部分监听：
// 调用conn.ExistsW(path) 或GetW(path)为对应节点设置监听，该监听只生效一次，开启一个协程处理chanel中传来的event事件
// 注意：watchCreataNode一定要放在一个协程中，不能直接在main中调用，不然会阻塞main
func zkWatchLocal() {
	conn := getConnect()

	_, _, eventChan, _ := conn.ExistsW("/test")

	// 协程调用监听事件
	go func(eventChan <-chan zk.Event) {
		event := <-eventChan
		fmt.Println("----------zkEventCallback start-----------------")
		fmt.Println("path: ", event.Path)
		fmt.Println("type: ", event.Type.String())
		fmt.Println("state: ", event.State.String())
		fmt.Println("----------zkEventCallback end-----------------\n")
	}(eventChan)

	// 触发创建数据操作
	zkModify()

}

// 1.如果即设置了全局监听又设置了部分监听，那么最终是都会触发的，并且全局监听在先执行
// 2.如果设置了监听子节点，那么事件的触发是先子节点后父节点

func main() {
	// 基本操作
	// zkCreate()
	// zkGet()
	// zkModify()
	// zkDelete()

	// 监听
	// zkWatchGlobal()
	zkWatchLocal()

}
