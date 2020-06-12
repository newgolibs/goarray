package getbyIndex_234_0_test

import (
    "github.com/newgolibs/goarray"
)

type  GetbyIndex_234_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []string{"a","b","d"},
            "b",
            1,
            "获取第一个索引 - id=74",
        },
        []interface{}{
            []string{"a","b","e"},
            "e",
            2,
            " - id=75",
        },
        []interface{}{
            []string{"a","b","e"},
            "a",
            0,
            " - id=76",
        },
        []interface{}{
            []string{"a","b","e"},
            nil,
            4,
            "要获取的索引超出的数组的长度，返回nil - id=77",
        },
        []interface{}{
            []string{"a","b","e"},
            "b",
            -2,
            " - id=78",
        },
        []interface{}{
            []string{"a","b","e"},
            "e",
            -1,
            " - id=79",
        },
        []interface{}{
            []interface{}{"a","b","e"},
            "e",
            -1,
            " - id=80",
        },
        []interface{}{
            []interface{}{"a","b","e",goarray.Arr{}},
            goarray.Arr{},
            -1,
            " - id=81",
        },
	}

}