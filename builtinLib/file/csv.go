package main

/**
 * @Author HZreal
 * @Date 2024-12-16
 * @File: csv.go
 * @Description:
 */

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// 打开文件，若文件不存在则创建
	file, err := os.Create("assets/output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建 CSV 写入器
	writer := csv.NewWriter(file)
	defer writer.Flush() // 刷新缓存

	// 定义要写入的记录
	records := [][]string{
		{"Name", "Address", "Phone Number"},
		{"John Doe", "1234 Elm St, Springfield", "(555) 123-4567"},
		{"Jane Smith", "4567 Oak St, Gotham", "(555) 987-6543"},
	}

	// 写入所有记录
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 你可以看到，`csv.Writer` 会自动将含有逗号和换行符的字段用双引号包裹起来
	log.Println("CSV 文件写入成功")
}
