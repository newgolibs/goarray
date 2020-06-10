package Arr

import (
	"github.com/newgolibs/goarray"
	"github.com/newgolibs/goarray/test/Arr/last_232_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Last_232_0_test struct {
}

func TestArr_last_232_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range last_232_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Last_232_0_test{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Last_232_0_test) run(input, arg interface{}) interface{} {
	arr := &goarray.Arr{Value: input}
	return arr.SetValue(input).
		Getlast()
}
