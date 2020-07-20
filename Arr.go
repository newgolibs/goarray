package goarray

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (this *Arr) SetValue(Value interface{}) *Arr {
	val := reflect.ValueOf(Value)
	if val.Kind() != reflect.Slice {
		panic("传入的类型应该是一个切片：[]xxxx")
	}
	this.Value = Value
	return this
}

func (this *Arr) Getfirst() interface{} {
	reflectv := reflect.ValueOf(this.Value)
	if reflectv.Len() == 0 {
		return nil
	}
	return reflectv.Index(0).Interface()
}

func (this *Arr) Getlast() interface{} {
	reflectv := reflect.ValueOf(this.Value)
	if reflectv.Len() == 0 {
		return nil
	}
	return reflectv.Index(reflectv.Len() - 1).Interface()
}

func (this *Arr) Tojson() string {
	marshal, err := json.Marshal(this.Value)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return ""
	}
	return string(marshal)
}

func (this *Arr) GetbyIndex(index int) interface{} {
	reflectv := reflect.ValueOf(this.Value)
	// 空数组
	if reflectv.Len() == 0 {
		return nil
	}
	// 溢出
	if reflectv.Len() < index+1 {
		return nil
	}
	// 倒数
	if index < 0 {
		return reflectv.Index(reflectv.Len() + index).Interface()
	}
	return reflectv.Index(index).Interface()
}

func (this *Arr) Contain(findvar interface{}) bool {
	reflectv := reflect.ValueOf(this.Value)
	for i := 0; i < reflectv.Len(); i++ {
		// 第一步，判断类型一致
		if reflectv.Index(i).Kind() == reflect.ValueOf(findvar).Kind() {
			// 第二步：判断值一致
			if reflectv.Index(i).Interface() == findvar {
				return true
			}
		}
	}
	return false
}
