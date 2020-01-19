package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	// if there is not the test below, coverage rate does not change

	// t.Run("collection of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := Sum(numbers)
	// 	want := 6
	//
	// 	if got != want {
	// 		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	// 	}
	// })

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 5}, []int{3, 0})
	want := []int{8, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
