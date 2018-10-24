package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			&Person{
				"Marco",
				Profile{27, "Caprecano"},
			},
			[]string{"Marco", "Caprecano"},
		},
		{
			"Struct with slices of Profile",
			[]Profile{
				{27, "Caprecano"},
			},
			[]string{"Caprecano"},
		},
		{
			"Struct with array of Profile",
			[2]Profile{
				{27, "Caprecano"},
				{27, "Cologna"},
			},
			[]string{"Caprecano", "Cologna"},
		},
		{
			"Struct with map of string",
			map[string]string{
				"Foo": "Bar",
				"Bar": "Foo",
			},
			[]string{"Bar", "Foo"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %s", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps to be indipendent by the order", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Bar": "Foo",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Foo")

	})
}

func assertContains(t *testing.T, got []string, wantEntry string) {
	t.Helper()
	contains := false
	for _, stringFound := range got {
		if stringFound == wantEntry {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("got %v want %s", got, wantEntry)
	}
}
