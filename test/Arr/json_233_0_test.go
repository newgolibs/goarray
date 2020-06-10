package Arr

import (
	"github.com/newgolibs/goarray"
	"github.com/newgolibs/goarray/test/Arr/json_233_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Json_233_0_test struct {
}

func TestArr_json_233_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range json_233_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Json_233_0_test{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Json_233_0_test) run(input, arg interface{}) interface{} {
	o := goarray.Arr{}
	return o.SetValue(input).
		Tojson()
}
