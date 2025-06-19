package utils

import (
	"fmt"
	"reflect"
)

func PrintObj(obj any) {
	// 获取反射值
	v := reflect.ValueOf(obj)

	// 如果不是结构体则直接返回
	if v.Kind() != reflect.Struct {
		return
	}

	// 获取结构体类型
	t := v.Type()

	// 遍历所有字段
	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}
