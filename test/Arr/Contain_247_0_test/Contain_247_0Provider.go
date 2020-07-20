package Contain_247_0_test

type Contain_247_0_testProvider struct {
}

func DataProvider() []interface{} {

	return []interface{}{
		[]interface{}{
			[]string{"a", "b", "c"},
			true,
			"a",
			" - id=99",
		},
		[]interface{}{
			[]string{"a", "b", "c"},
			false,
			"cc",
			" - id=100",
		},
	}

}
