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
        },
        []interface{}{
            []string{"a","b","e"},
            "e",
            2,
        },
        []interface{}{
            []string{"a","b","e"},
            "a",
            0,
        },
        []interface{}{
            []string{"a","b","e"},
            nil,
            4,
        },
        []interface{}{
            []string{"a","b","e"},
            "b",
            -2,
        },
        []interface{}{
            []string{"a","b","e"},
            "e",
            -1,
        },
        []interface{}{
            []interface{}{"a","b","e"},
            "e",
            -1,
        },
        []interface{}{
            []interface{}{"a","b","e",goarray.Arr{}},
            goarray.Arr{},
            -1,
        },
	}

}