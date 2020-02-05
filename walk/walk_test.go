package walk

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
	tests := []struct {
		desc          string
		input         interface{}
		expectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested field",
			Person{
				"Chris",
				Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointer to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{22, "London"},
				{34, "Tokyo"},
			},
			[]string{"London", "Tokyo"},
		},
		{
			"Arrays",
			[2]Profile{
				{22, "London"},
				{34, "Tokyo"},
			},
			[]string{"London", "Tokyo"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var got []string
			walk(tc.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.expectedCalls) {
				t.Errorf("got %v, want %v", got, tc.expectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
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

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}

}
