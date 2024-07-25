package main

import (
	"database/sql" // 标准库
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql驱动连接包
)

// var db *sql.DB    // 创建一个数据库类型 而不是数据库连接，
// Go中的数据库连接来自内部实现的连接池，连接的建立是惰性的，即连接将会在实际操作的时候，由连接池创建并维护使用。

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

// 使用标准库的sql连接
func connectMysql() (db *sql.DB) {
	// sql.Open() 函数创建数据库类型，第一个参数是数据库驱动名，第二个参数是连接信息字符串
	db, err := sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("conn error :", err)
	}

	// 尝试链接
	err = db.Ping()
	if err != nil {
		fmt.Println("ping error :", err)
	}
	fmt.Println("conn success!")
	return
}

func insertWithSql() {
	db := connectMysql()

	tx, _ := db.Begin()
	tx.Exec("insert into person(username, sex, email) values(?, ?, ?)", "stu006", "man", "stu0006@qq.com")
	tx.Exec("insert into person(username, sex, email) values(?, ?, ?)", "stu007", "woman", "stu0007@qq.com")
	tx.Commit()

}

func selectWithSql() {
	db := connectMysql()

	// Exec() 无需返回数据行的查询,一般用于增删改

	// QueryRow() 返回单行的查询
	// row := db.QueryRow("select * from person where user_id = ?", 1)
	// var singlePerson Person
	// _ = row.Scan(&singlePerson.UserId, &singlePerson.Username, &singlePerson.Sex, &singlePerson.Email)
	// fmt.Println("get one record ----", singlePerson)

	// Query() 查询，返回数据集
	rows, err := db.Query("select user_id, username, email from person where user_id>?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
	}
	defer rows.Close()
	var personList []Person
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.UserId, &person.Username, &person.Email)
		if err != nil {
			fmt.Println("select err :", err)
		}
		personList = append(personList, person)
	}
	fmt.Println("query result is ", personList)
	byteArr, _ := json.Marshal(personList)
	fmt.Println("json data is ", string(byteArr))

}

func updateWithSql() {
	db := connectMysql()

	// db.Exec() 无需返回数据集，只返回结果，一般用于增删改
	db.Exec("update person set username=? where user_id=?", "stu00007", 7)

}

func deleteWithSql() {
	db := connectMysql()
	db.Exec("delete from person where user_id=?", 7)
}

func usePrepare() {
	db := connectMysql()

	// db.Prepare()
	multiPerson := []Person{
		{UserId: 10, Username: "hh11", Sex: "man", Email: "hh11@163.com"},
		{UserId: 11, Username: "hh12", Sex: "woman", Email: "hh12@163.com"},
		{UserId: 12, Username: "hh13", Sex: "man", Email: "hh13@163.com"},
	}
	stmt, err := db.Prepare("insert into person(user_id, username, sex, email) values (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	// 插入多条记录
	for _, person := range multiPerson {
		stmt.Exec(person.UserId, person.Username, person.Sex, person.Email)
	}

}

func main() {
	// 使用标准库的sql连接
	// insertWithSql()
	// updateWithSql()
	// deleteWithSql()
	// selectWithSql()
	// usePrepare()

}
