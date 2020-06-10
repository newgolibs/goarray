package Arr

import (
	"github.com/newgolibs/goarray"
	"github.com/newgolibs/goarray/test/Arr/getbyIndex_234_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type GetbyIndex_234_0_test struct {
}

func TestArr_getbyIndex_234_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range getbyIndex_234_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], GetbyIndex_234_0_test{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (GetbyIndex_234_0_test) run(input, arg interface{}) interface{} {
	o := goarray.Arr{}
	return o.SetValue(input).
		GetbyIndex(arg.(int))
}
