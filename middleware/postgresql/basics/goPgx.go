package main

/**
 * @Author huang
 * @Date 2024-07-25
 * @File: goPgx.go
 * @Description:
 */

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5" // 官方推荐 且 gorm 驱动也基于此库
	"log"
)

var ctx context.Context
var client *pgx.Conn

func init() {
	var err error
	client, err = pgx.Connect(ctx, "postgres://username:password@localhost:5432/dbname")
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	// 测试数据库连接
	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

}

func use() {

	// 执行查询
	query := "SELECT id, name FROM users"
	rows, err := client.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var id int
		var name string
		scanErr := rows.Scan(&id, &name)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// 检查查询错误
	if rowErr := rows.Err(); rowErr != nil {
		log.Fatal(rowErr)
	}
}

func main() {
	use()
}
