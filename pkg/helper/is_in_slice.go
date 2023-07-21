package helper

import (
	"fmt"
	"reflect"
)

// IsInSlice 判断变量是否为数组或切片中的某个值
func IsInSlice(value string, slice any) bool {
	sliceValue := reflect.ValueOf(slice)
	// .Kind() 获取被反射的值的类型
	// reflect.Slice 反射中的切片类型
	if sliceValue.Kind() != reflect.Slice {
		fmt.Printf("传入的值：%v 不是切片类型！", sliceValue)
	}

	// .Len() 被反射的值的长度
	for i := 0; i < sliceValue.Len(); i++ {
		// .Index(i) 被反射的值的索引, 返回的是reflect.Value
		// .Interface() 将reflect.Value转成Interface()
		element := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(element, value) {
			return true
		}
	}
	return false
}
