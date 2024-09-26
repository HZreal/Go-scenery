package main

/**
 * @Author elastic·H
 * @Date 2024-09-24
 * @File: structImplUnmarshal.go
 * @Description:
 */

type User struct {
	ID     int    `json:"id" bson:"id" yaml:"id" xml:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	Avatar string `json:"avatar"`
}

// UnmarshalJSON 自定义解析
// struct 实现 Unmarshaler 接口, 便可以实现 JSON 序列化和反序列化的过程
func (p *User) UnmarshalJSON(data []byte) error {
	// 示例代码使用 jsonitor 代为解析
	p.ID = 2
	p.Age = 24
	p.Name = "my_test_name"
	return nil
}

// MarshalJSON 自定义编码
func (p *User) MarshalJSON() ([]byte, error) {
	// 自己编码json
	return []byte(`{"test":"name_test"}`), nil
}

func main() {

}
