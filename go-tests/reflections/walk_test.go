package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Street int
	City   string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Yashas"},
			[]string{"Yashas"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Yashas", "XYZ"},
			[]string{"Yashas", "XYZ"},
		},
		{
			"struct with nested string fields",
			Person{"Yashas", Address{2, "XYZ"}},
			[]string{"Yashas", "XYZ"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			testFunc := func(input string) {
				got = append(got, input)
			}
			Walk(test.Input, testFunc)

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("got %v, want %v", got, test.Expected)
			}

		})
	}

}
