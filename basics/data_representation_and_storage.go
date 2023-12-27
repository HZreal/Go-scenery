package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

/**
* @author
* @description //TODO
* @date
* @param
* @return
**/
func parseMpFile(mpFilePath string) {
	file, err := os.Open(mpFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 读取行数和列数
	var rowLen, colLen int32
	err = binary.Read(file, binary.LittleEndian, &rowLen)
	if err != nil {
		fmt.Println("Error reading rowLen:", err)
		return
	}
	err = binary.Read(file, binary.LittleEndian, &colLen)
	if err != nil {
		fmt.Println("Error reading colLen:", err)
		return
	}
	fmt.Println("rowLen  ----> ", rowLen)
	fmt.Println("colLen  ----> ", colLen)

	// 跳过保留字节
	file.Seek(16, 1)

	// 遍历每一组经纬度数据
	for i := 0; i < int(rowLen*colLen); i++ {
		var lon, lat float64
		err = binary.Read(file, binary.LittleEndian, &lon)
		if err != nil {
			fmt.Println("Error reading longitude:", err)
			return
		}
		err = binary.Read(file, binary.LittleEndian, &lat)
		if err != nil {
			fmt.Println("Error reading latitude:", err)
			return
		}
		fmt.Printf("lon, lat ----> %f, %f\n", lon, lat)
	}
}
func main() {
	mpFilePath := "/path/to/Go-scenery/basics/data_representation_and_storage.go"
	parseMpFile(mpFilePath)
}
