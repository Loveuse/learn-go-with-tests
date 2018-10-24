package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := Sum(input)
	want := 15

	if got != want {
		t.Errorf("got '%d' want '%d' input '%v'", got, want, input)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([][]int{[]int{1, 2, 3}, []int{4, 5, 6}})

	want := []int{6, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
