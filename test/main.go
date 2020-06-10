package main

import (
	"github.com/newgolibs/goarray"
	"fmt"
	"reflect"
)

func main() {
	fmt.Print("cc")
	var a goarray.Arr
	a.SetValue("ok")
	a.SetValue([]string{"aaa", "dddd"})
	fmt.Printf("%+v\n", a)
	fmt.Printf("转换回来看看？%+v，%+v\n", a.Value,reflect.TypeOf(a.GetValue()))
	val := reflect.ValueOf(a.Value)
	fmt.Printf("%+v\n", val)
	fmt.Printf("%+v\n", val.Kind())
	fmt.Printf("%+v\n", val.Kind()==reflect.Slice)
}
