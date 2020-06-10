package first_231_0_test

type  First_231_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []string{"aa","bb"},
            "aa",
            "",
        },
        []interface{}{
            []interface{}{1,"bb"},
            1,
            "",
        },
        []interface{}{
            []int{},
            nil,
            "",
        },
	}

}