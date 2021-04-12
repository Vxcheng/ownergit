package reflection

import (
	"reflect"
	"testing"
)

/**
red-green-blue

1.首先编写测试
2.尝试运行测试
3.为测试的运行编写最小量的代码，并检查测试的失败输出
4.编写足够的代码使测试通过
5.重构
*/

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{

		{
			"Struct with two string fields",
			&struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Chris",
				struct {
					Age  int
					City string
				}{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Struct with two array fields",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got is %v", got)
			}
		})
	}

	t.Run("with Map", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, got []string, want string) {
	exist := false

	for _, v := range got {
		if v == want {
			exist = true
			break
		}
	}

	if !exist {
		t.Errorf("expected %+v to contain '%s' but it didnt", got, want)
	}
}

type Profile struct {
	Age  int
	City string
}
