package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

var uri = "http://127.0.0.1:8000/api/parseRawBody"

// 获取键值对的body
func getUrlValues() url.Values {
	// 方式1
	data1 := url.Values{"filename": {"TiMi"}, "id": {"123"}}
	fmt.Println("data1--------\n", data1)
	// 方式2
	data2 := url.Values{}
	data2.Set("name", "TiMi")
	data2.Set("id", "123")
	fmt.Println("data2--------\n", data2)
	// 方式3
	data3 := make(url.Values)
	data3["name"] = []string{"TiMi"}
	data3["id"] = []string{"123"}
	fmt.Println("data3--------\n", data3)
	/*
	   map[id:[123] name:[TiMi]]
	   map[id:[123] name:[TiMi]]
	   map[id:[123] name:[TiMi]]
	*/
	return data1
}

// 以PostForm的方式发送body为键值对的post请求
func sendFormData() {
	// 构造formData
	data := getUrlValues()

	res, err := http.PostForm(uri, data)
	if err != nil {
		fmt.Println("err=", err)
	}
	// http返回的response的body必须close,否则就会有内存泄露
	defer func() {
		res.Body.Close()
		fmt.Println("res finish ------- ")
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	fmt.Println("body--------\n", string(body))

}

// 发送body为键值对的post请求   application/x-www-form-urlencoded
func sendFormDataWithUrlencoded() {
	data := getUrlValues()

	ipPort := "http://127.0.0.1:5000"
	resource := "/test/install"
	u, _ := url.ParseRequestURI(ipPort)
	u.Path = resource
	urlStr := u.String()
	fmt.Println("########## urlStr等价于uri ###########-----------", urlStr)

	request, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 等价于
	// res, err := http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	client := &http.Client{}
	res, err := client.Do(request) // 发送
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// http返回的response的body必须close,否则就会有内存泄露
	defer func() {
		res.Body.Close()
		fmt.Println("finish")
	}()

	// 读取body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	fmt.Println(string(body))

}

func sendFormDataWithJson() {
	bodyMap := map[string]interface{}{
		"filename": "123.txt",
		"name":     "huang",
	}
	bytesData, _ := json.Marshal(bodyMap)
	res, err := http.Post(uri, "application/json;charset=UTF-8", bytes.NewReader(bytesData))
	if err != nil {
		fmt.Println("err --------", err)
		return
	}
	// 最后关闭res.Body文件
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	// 使用ioutil.ReadAll将res.Body中的数据读取出来,并使用body接收
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err --------", err)
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	fmt.Println(*str)

}

// 以二进制格式上传文件
func sendFormDataWithFile() {
	filePath := "/Users/hz/Desktop/hz/go/src/backend_master/asset/install.sh.zip"
	filebyteArr, _ := ioutil.ReadFile(filePath)
	fmt.Println("zip 文件  []byte   -----------\n", filebyteArr)
	res, err := http.Post(uri, "multipart/form-data", bytes.NewReader(filebyteArr))

	// bodyMap := make(map[string]string)
	// bodyMap["file"] = string(filebyteArr)
	// bodyMap["filename"] = "install.sh.zip"
	// formDataByteArr, _ := json.Marshal(bodyMap)
	// res, err := http.Post(uri, "multipart/form-data", bytes.NewReader(formDataByteArr))

	if err != nil {
		fmt.Println("err=", err)
	}
	// http返回的response的body必须close,否则就会有内存泄露
	defer func() {
		res.Body.Close()
		fmt.Println("res finish -------------")
	}()

	// 读取body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	fmt.Println("body ----------", string(body))

}

func sendFormDataWithFile2() {
	filePath := "/Users/hz/Desktop/hz/go/src/backend_master/asset/install.sh.zip"
	filebyteArr, _ := ioutil.ReadFile(filePath)

	bodyMap := make(map[string]interface{})
	bodyMap["file"] = filebyteArr
	bodyMap["filename"] = "install.sh.zip"

	dataVal := url.Values{}
	// if bodyMap != nil {
	//	for k, v := range bodyMap{
	//		//dataVal.Set(k, datautils.ToString(v))
	//		//dataVal.Set(k, v)
	//	}
	// }
	res, err := http.Post(uri, "multipart/form-data", strings.NewReader(dataVal.Encode()))
	// 最后关闭res.Body文件
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	var body []byte
	// 使用ioutil.ReadAll将res.Body中的数据读取出来,并使用body接收
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err=", err)
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	fmt.Println(*str)

}
func main() {
	// sendFormData()
	// sendFormDataWithUrlencoded()
	// sendFormDataWithJson()
	sendFormDataWithFile()
	// sendFormDataWithFile2()
}
