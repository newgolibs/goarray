package Arr

import (
	"github.com/newgolibs/goarray"
	"github.com/newgolibs/goarray/test/Arr/Contain_247_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Contain_247_0 struct {
}

func TestArr_Contain_247_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Contain_247_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Contain_247_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Contain_247_0) run(input, arg interface{}) interface{} {
	var a goarray.Arr
	return a.SetValue(input).
		Contain(arg)
}
