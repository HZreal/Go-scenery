package main

// 1.1.1. etcd介绍
// Go语言实现的分布式、高可用、强一致性、小型kv数据存储服务
// 如果你需要一个分布式存储仓库来存储配置信息，并且希望这个仓库读写速度快、支持高可用、部署简单、支持http接口，那么就可以使用etcd

// 1.1.2. etcd应用场景
// 配置管理
// 服务注册
// 服务发现Service Discovery    ----- 动态增添或减少新的服务节点，都可被监测到
// leader竞选                  -----  若检测到master节点宕机，会在slave中选出某个替换
// 应用调度
// 集群监控
// 分布式队列
// 分布式锁                    ------  节点竞争临界资源，需通过etcd获取锁才能去操作，其他节点等待

// ETCD架构、原理
// gRPC Server: 负责客户端的api请求 以及 节点与节点之间通信
// wal(write ahead log)日志: 保持数据的一致性。主节点将每次操作形成日志条目，并持久化到本地磁盘，然后通过网络IO发送给其他节点。其他节点根据日志的逻辑时钟(TERM)和日志编号(INDEX)来判断是否将该日志记录持久化到本地
// snapshot快照: 负责备份当前节点数据状态，可用于其他节点同步主节点数据时的一致性迁移，类似于redis集群中的主从复制
// boltdb存储引擎: 为操作的key创建一个索引(B+树)，该B+树存储了key对应的历史版本信息；实现事务

// client通过api请求gRPC Server 发送某操作指令，etcd接收操作并写入wal日志，将日志广播到其他follow，在接收到大多数节点同意写入的响应后，将数据提交写入磁盘，并将提交操作日志广播通知其他节点也提交，最后将操作结果返给client
// 内存中，采用B树，通过key查询revision版本信息， 磁盘中，采用B+树，通过revision查找对应的value
// lease租约，类似于redis中对key设置过期时间，redis会监测每个key的过期时间，而etcd将大多数相同过期时间的key的过期时间作为一个对象进行管理，只需检测过期实体即可，lease租约可以绑定多个key的过期信息，只要lease过期，其绑定的key全部过期
// raft共识一致算法(leader选举 + 日志复制)：  动画描述: http://thesecretlivesofdata.com/raft/ 讲解视频: https://www.bilibili.com/video/BV1ny4y1G7o7/
// 			选举超时()  心跳超时()

// 相对于zookeeper的优势
// 使用 Go 语言编写部署简单；支持HTTP/JSON API,使用简单；使用 Raft 算法保证强一致性让用户易于理解，zk共识算法paxos难理解
// etcd 默认数据一更新就进行持久化。
// etcd 支持 SSL 客户端安全认证

// etcd基本命令使用
// etcdctl put key value      创建/修改key  可用参数lease=lease_id(16进制整数) 设置key绑定到某个lease
// etcdctl get key            获取key
// etcdctl del key            删除key
// etcdctl watch key          监听key的状态变更，如其他节点的修改，此节点可监测到
// etcdctl get key -w json    查看key具体的版本信息，以json形式输出
// 输出：{"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":4,"raft_term":2},"kvs":[{"key":"bmFtZQ==","create_revision":2,"mod_revision":4,"version":3,"value":"aHVh"}],"count":1}
// 字段解释：raft_term表示任期，基于raft算法进行选举，当在网络分区等原因下进行修复节点集群时，多个leader凭借任期值大小确定继续作为leader还是退回follow； revision表示全局版本号，只要etcd修改，就会自增加1； create_revision表示创建key时的revision； mod_revision表示修改key时的revision；
// etcdctl get key --rev=2    查看key第2个版本的值
// etcdctl txn -i             交互式地进入一个事务
// 首先compares输入比较条件，然后输入判断成功的操作，然后输入判断失败的操作，最终查看返回的结果     相关函数value("key")获取具体值  mod("key")获取修改版本  create("key")版本信息
// etcdctl lease grant ttl    创建lease租约，过期时间为ttl

// etcd集群
// 一个leader   多个follower

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	// "github.com/coreos/etcd/client/v3"
	"time"
)

// 基本put get 操作
func etcdOperation() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "key1", "value1")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "key1")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Println(" KeyValue struct---------\n", ev.Key, ev.Value, ev.CreateRevision, ev.ModRevision, ev.Lease)
	}

}

// watch用来获取未来更改的通知
// 将代码保存编译执行，此时程序就会等待etcd中key1的变化。
// 例如：我们打开终端执行以下命令修改、删除、设置这个key1，程序都能收到通知
func useWatch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// watch key:lmh change
	rch := cli.Watch(context.Background(), "key1") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

// lease租约
func useLease() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()

	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /lmh/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}

// keepAlive
func keepAlive() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()

	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key 'foo' will be kept forever
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

// 基于etcd实现分布式锁
// go.etcd.io/etcd/clientv3/concurrency在etcd之上实现并发操作，如分布式锁、屏障和选举。
func demoDistributedLockBasedOnETCD() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	// cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}

func main() {
	etcdOperation()
	// useWatch()
	// useLease()
	// keepAlive()
	// demoDistributedLockBasedOnETCD()
}
