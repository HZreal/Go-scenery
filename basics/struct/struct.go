package _struct

/**
 * @Author nico
 * @Date 2025-04-03
 * @File: struct.go
 * @Description:
 */

import (
	"fmt"
	"reflect"
)

// GetValueByField 根据键映射到结构体字段名并获取字段值
func GetValueByField(obj any, fieldName string) (string, error) {
	v := reflect.ValueOf(obj)

	// 处理指针情况，获取结构体本身
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 确保是结构体
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct, got %T", obj)
	}

	// 根据 key 获取结构体字段名，移到外面去，不由此转换工具控制
	// fieldName, exists := constant.CustomFieldMap[fieldName]
	// if !exists {
	// 	return "", fmt.Errorf("key %s not found in mapping", key)
	// }

	// 通过字段名获取字段值
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return "", fmt.Errorf("field %s not found in struct", fieldName)
	}

	// 返回字段的值
	return field.String(), nil
}

// 设置结构体字段的值，这个 key 是该字段的名称字符串
func SetValueByField(obj any, key, value string) error {
	v := reflect.ValueOf(obj)

	// 处理指针情况，获取结构体本身
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer, got %T", obj)
	}
	v = v.Elem()

	// 确保是结构体
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %T", obj)
	}

	// 通过字段名获取字段值
	field := v.FieldByName(key)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found in struct", key)
	}

	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", key)
	}

	val := reflect.ValueOf(value)
	if val.Type().ConvertibleTo(field.Type()) {
		field.Set(val.Convert(field.Type())) // 类型匹配则赋值
	}

	return nil
}

// 设置结构体字段的值，基于 map，map 的键是结构体字段名称
func SetValueByMap(obj any, kv map[string]string) error {
	if kv == nil {
		return nil
	}

	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected pointer, got %T", obj)
	}

	for k, v := range kv {
		if err := SetValueByField(obj, k, v); err != nil {
			return err
		}
	}

	return nil
}
