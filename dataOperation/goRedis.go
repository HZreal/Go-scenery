package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Redis支持数据的持久化(RDB、AOF)，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
// Redis不仅仅支持简单的key-value类型的数据，同时还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
// Redis支持数据的备份，即master-slave模式的数据备份

// 性能极高 适合做缓存。
// 丰富的数据类型
// 所有操作都是原子性的
// 丰富的特性

func connectRedis() (conn redis.Conn) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
	}
	fmt.Println("redis connect successfully")
	return
}

// 字符串操作
func operateString() {
	conn := connectRedis()

	// set key value
	// _, err := conn.Do("Set", "key1", "value1")
	// _, err := conn.Do("Set", "key2", 222)
	// if err != nil {
	// 	fmt.Println("set err ", err)
	// 	conn.Close()
	// 	return
	// }

	// get key
	// reply, err := conn.Do("Get", "key1")   // 返回的reply类型为interface{} -> 且是[]byte类型
	// if err != nil {
	// 	fmt.Println("get key1 failed,", err)
	// 	conn.Close()
	// 	return
	// }
	// fmt.Printf("type is %T\nvalue is %v\n", reply, reply)
	// fmt.Println("get value is", string(reply.([]byte)))       // interface{}类型 断言转成[]byte类型，再转成string

	// 直接使用封装好的进行转换  redis.String/Int/
	// value, _ := redis.String(conn.Do("Get", "key1"))
	// value, _ := redis.Int(conn.Do("Get", "key2"))
	// fmt.Println(value)

	// mset批量设置
	// _, _ = conn.Do("MSet", "name", "hu", "age", 22)

	// mget 获取相同类型的值
	// intArr, _ := redis.Ints(conn.Do("MGet", "age", "age1"))
	// fmt.Println(intArr)
	// stringArr, _ := redis.Strings(conn.Do("MGet", "name", "name1"))
	// fmt.Println(stringArr)

	// mget 获取不同类型的值，返回[]interface{}
	valueArr, _ := redis.Values(conn.Do("MGet", "name", "age"))
	for _, v := range valueArr {
		// v 的类型为 []byte
		fmt.Println(string(v.([]byte)))
	}

	// 设置过期时间
	// conn.Do("expire", "k", 10)

}

// list队列操作
func operateList() {
	conn := connectRedis()

	// lpush key value1 value2 value3
	// conn.Do("lpush", "book_list", "abc", "ceg", 300)

	// lpop key
	// v, _ := redis.String(conn.Do("lpop", "book_list"))
	// fmt.Println(v)

	// lrange key start end
	values, _ := redis.Values(conn.Do("lrange", "book_list", 0, -1))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

}

// hash表
func operateHash() {
	conn := connectRedis()

	// hset key field value
	// conn.Do("HSet", "books", "name", "golang")

	// hmset key field1 value1 field2 value2
	// conn.Do("HMSet", "books", "desc", "just for learn", "sale", 54)

	// hget key field
	// v, _ := redis.Int(conn.Do("HGet", "books", "sale"))
	// fmt.Println(v)

	// hgetall key
	valueArr, _ := redis.Values(conn.Do("HGetAll", "books"))
	for _, v := range valueArr {
		fmt.Println(string(v.([]byte)))
	}

}

// func init() {
// 	pool = &redis.Pool{
// 		MaxIdle:16,                       // 最初的连接数量
// 		IdleTimeout:300,                  // 连接关闭时间 300秒 （300秒不使用自动关闭）
// 		MaxActive:0,                      // 连接池中最大连接数量，不确定可以用0（0表示自动定义），按需分配
// 		MaxConnLifetime: 1000,            //
// 		Dial: func() (redis.Conn ,error){     // 要连接的redis数据库
// 			return redis.Dial("tcp","localhost:6379")
// 		},
// 	}
// }
// redis连接池
func redisPool() {
	var pool *redis.Pool
	pool = &redis.Pool{
		MaxIdle:         16,   // 最初的连接数量
		IdleTimeout:     300,  // 连接关闭时间 300秒 （300秒不使用自动关闭）
		MaxActive:       0,    // 连接池中最大连接数量，不确定可以用0（0表示自动定义），按需分配
		MaxConnLifetime: 1000, //
		Dial: func() (redis.Conn, error) { // 要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	// 获取活跃连接数、空闲连接数
	activeCount := pool.ActiveCount()
	idleCount := pool.IdleCount()
	fmt.Println("活跃连接数、空闲连接数分别为：", activeCount, idleCount)

	conn := pool.Get() // 从连接池，取一个连接
	defer conn.Close() // 函数运行结束 ，把连接放回连接池
	value, _ := redis.String(conn.Do("Get", "name"))
	fmt.Println(value)

}

func main() {
	// operateString()
	// operateList()
	// operateHash()
	redisPool()
}
