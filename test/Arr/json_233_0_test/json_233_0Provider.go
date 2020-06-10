package json_233_0_test

type  Json_233_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []string{"a","e"},
            "[\"a\",\"e\"]",
            "",
        },
        []interface{}{
            []int{1,2,3},
            "[1,2,3]",
            "",
        },
        []interface{}{
            []string{},
            "[]",
            "",
        },
	}

}