package main

/**
 * @Author nico
 * @Date 2024-12-03
 * @File: excelize.go
 * @Description:
 */

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"strings"
)

func extractDimensions(dimension string) (int, int) {
	parts := strings.Split(dimension, ":")
	if len(parts) < 2 {
		return 0, 0
	}

	// 提取最后一个单元格的列号和行号
	endCell := parts[1]

	// 分离列号和行号
	colStr := strings.TrimRight(endCell, "0123456789")
	rowStr := strings.TrimLeft(endCell, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// 转换行号
	rowCount, _ := strconv.Atoi(rowStr)

	// 转换列号
	colCount := 0
	for _, ch := range colStr {
		colCount = colCount*26 + int(ch-'A'+1)
	}

	return rowCount, colCount
}

func test() {
	filepath := "assets/demo.xlsx"

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Fatalf("%s Not Exist", filepath)
	}

	sheetDimension, err := f.GetSheetDimension("Sheet1")
	if err != nil {
		return
	}
	log.Println("sheetDimension  ---->  ", sheetDimension)

	rowCount, colCount := extractDimensions(sheetDimension)
	log.Println("rowCount,rowCount  ---->  ", rowCount, colCount)
}

func main() {
	test()
}
