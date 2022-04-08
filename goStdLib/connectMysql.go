package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//var db *sql.DB

type SlaveNode struct {
	Id         int    `json:"id"`
	NodeName   string `json:"node_name" db:"node_name"`
	Ip         string `json:"ip" db:"ip"`
	Port       int    `json:"port" db:"port"`
	Status     int    `json:"status" db:"status"`
	SystemInfo string `json:"system_info" db:"system_info"`
	ExtraData  string `json:"extra_data" db:"extra_data"`
	createTime time.Time
	updateTime time.Time
	isDelete   bool
}

func connectMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/LineMasterNode")
	if err != nil {
		fmt.Println("conn err :", err)
	}

	// 尝试链接
	err = db.Ping()
	if err != nil {
		fmt.Println("ping err :", err)
	}
	fmt.Println("conn success!") // conn success!
	return db
}

func main() {
	db := connectMysql()

	//db.Exec()   // 不需要返回数据集，只返回结果

	rows, err := db.Query("select id, node_name, ip, port, status from SlaveNode") // 查询，返回数据集
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	//var slaveArr []string
	var slaveArr []SlaveNode
	for rows.Next() {
		ss := SlaveNode{}
		err := rows.Scan(&ss.Id, &ss.NodeName, &ss.Ip, &ss.Port, &ss.Status)
		if err != nil {
			fmt.Println("select err :", err)
		}
		//byteArr, _ := json.Marshal(ss)
		//fmt.Printf("%T\n--------%v\n", byteArr, string(byteArr))
		//slaveArr = append(slaveArr, string(byteArr))
		slaveArr = append(slaveArr, ss)
	}
	byteArr2, _ := json.Marshal(slaveArr) // 结构体数组转为json
	fmt.Println(string(byteArr2))

	//db.QueryRow()  // 查询，期待返回一条数据
	//db.Prepare()

}
