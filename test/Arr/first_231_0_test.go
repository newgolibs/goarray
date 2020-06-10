package Arr

import (
    "github.com/newgolibs/goarray"
    "github.com/newgolibs/goarray/test/Arr/first_231_0_test"
    "log"
    "testing"
    "github.com/stretchr/testify/assert"
    )

func TestArr_first_231_0(t *testing.T){
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    for _, v := range first_231_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func run(input, arg interface{}) interface{} {
	var a goarray.Arr
	return a.SetValue(input).
		Getfirst()
}