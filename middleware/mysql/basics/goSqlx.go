package main

/**
 * @Author huang
 * @Date 2024-07-25
 * @File: _sqlx.go
 * @Description:
 */

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

/**
 * sqlx 基于Go标准库database/sql的封装，查询优化的较好
 */

// 使用sqlx连接
func connectMysqlWithSqlx() (db *sqlx.DB) {
	// sqlx.Open() 获取一个 sql.DB 类型对象，只是验证数据库参数，并没有创建数据库连接
	// sqlx.DB 是数据库的抽象，而不是数据库连接，有几个数据库就要创建几个 sqlx.DB 类型对象，因为它要维护一个连接池，因此不需要频繁的创建和销毁
	// sqlx.Open() 函数创建连接池，此时只是初始化了连接池，并没有连接数据库，连接都是惰性的，只有调用 sqlx.DB 的方法时，此时才真正用到了连接，连接池才会去创建连接
	// 连接池的工作原理
	// 		当调用 sqlx.DB的方法时，会首先去向连接池请求要一个数据库连接，如果连接池有空闲的连接，则返回一个空闲连接给调用方法使用，否则连接池将创建一个新的连接给调用方法使用；
	// 		一旦将数据库连接给到了方法中，连接就属于该调用方法了。方法执行完毕后，要么把连接所属权归还给连接池，要么传递给下一个需要数据库连接的方法中，最终所有调用方法都使用完连接后，将连接释放回到连接池中
	db, err := sqlx.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("conn error :", err)
	}

	// db.SetMaxIdleConns(n1 int) 设置连接池中的保持连接的最大连接数
	// 默认是0，表示连接池不会保持数据库连接的状态：即当连接释放回到连接池的时候，连接将会被关闭。这会导致连接在连接池中频繁的关闭和创建，我们可以设置一个合理的值
	db.SetMaxIdleConns(5)

	// db.SetMaxOpenConns(n2 int) 设置允许连接池创建数据库连接的最大连接数。包含正在使用的连接和连接池的连接
	// 如果你的方法调用需要用到一个连接，并且连接池已经没有了连接或者连接数达到了最大连接数。此时的方法调用将会被阻塞，直到有可用的连接才会返回。设置这个值可以避免并发太高导致连接mysql出现 too many connections的错误。该函数的默认设置是0，表示无限制
	db.SetMaxOpenConns(10)

	// 当前总连接数 = 池内连接数(空闲) + 活跃连接数
	// n1 <= n2

	// db.SetConnMaxIdleTime(d time.Duration) 设置保持连接的最大时间，超过这个时间，自动断开本连接
	db.SetConnMaxIdleTime(2 * time.Second)
	// db.SetConnMaxLifetime(d time.Duration) 设置连接的最长使用有效时间，如果过期，连接将被拒绝
	db.SetConnMaxLifetime(5 * time.Second)

	// 尝试连接，当调用了 Ping() 方法后，连接池一定会初始化一个数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("ping error :", err)
	}
	fmt.Println("connect success!")

	return
}

func insertWithSqlx() {
	db := connectMysqlWithSqlx()
	result, err := db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu002", "woman", "stu02@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("the id of inserted record is :", id)

	RowsAffectedNum, err := result.RowsAffected()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("the number of rows affected is ", RowsAffectedNum)

}

func selectWithSqlx() {
	db := connectMysqlWithSqlx()

	// Get() 查询一条记录，保存到结构体
	var singlePerson Person
	_ = db.Get(&singlePerson, "select * from person where user_id = ?", 1)
	fmt.Println("get one record ----", singlePerson)

	// Select 查询的多条记录，直接保存到结构体的切片中
	var personSlice []Person
	_ = db.Select(&personSlice, "select user_id, username, sex, email from person where user_id>?", 1)
	fmt.Println("select success! personList is :", personSlice)

	// Query()查询  需要定义多个字段的变量(结构体)进行接收
	var personList []Person
	rows, err := db.Query("select user_id, username, email from person")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() { // 迭代游标结果集
		person := Person{}
		// Scan读取一行记录，映射到定义的person
		// Scan的参数个数必须与查询结果的字段个数保持一致！
		err := rows.Scan(&person.UserId, &person.Username, &person.Email)
		if err != nil {
			fmt.Printf("get data failed, error:[%v]", err.Error())
		}
		fmt.Println("current record is ", person)
		personList = append(personList, person)
	}
	fmt.Println("query result is ", personList)

}

func updateWithSqlx() {
	db := connectMysqlWithSqlx()

	res, err := db.Exec("update person set username=? where user_id=?", "stu00001", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	RowsAffectedNum, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("the number of rows affected is ", RowsAffectedNum)
}

func deleteWithSqlx() {
	db := connectMysqlWithSqlx()

	res, err := db.Exec("delete from person where user_id=?", 4)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	RowsAffectedNum, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("the number of rows affected is ", RowsAffectedNum)
}

func transactionWithSqlx() {
	db := connectMysqlWithSqlx()

	// 开启事务
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu0002", "man", "stu0002@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		// 回滚
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	// 提交
	conn.Commit()

}

func main() {
	// 使用 sqlx 连接
	// insertWithSqlx()
	// updateWithSqlx()
	// deleteWithSqlx()
	selectWithSqlx()
	// transactionWithSqlx()
}
