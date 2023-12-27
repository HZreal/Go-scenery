package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

/**
* @author
* @description 按字节解析 .mp 二进制文件
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
	/**
	"""
	    .mp 数据为二进制流格式，字节顺序定义如下:
	    1~4字节，int32，行数，rowLen
	    5~8字节，int32，列数，colLen
	    9~24字节，保留
	    以下数据一共 rowLen * colLen 组，每组16字节
	    25~48字节，double * 2，经度、纬度
	    25~32字节，double，经度
	    33~40字节，double，纬度
	    41~56字节，double * 2，经度、纬度
	    ...
	    ...
	"""
	*/
	mpFilePath := "/path/to/Go-scenery/basics/data_representation_and_storage.go"
	parseMpFile(mpFilePath)
}
