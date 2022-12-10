package multiplicationtable

import (
	"reflect"
	"testing"
)

func TestMultiplicationTable(t *testing.T) {

	size := 5
	got := MultiplicationTable(size)

	want := [][]int{{1, 2, 3, 4, 5}, {2, 4, 6, 8, 10}, {3, 6, 9, 12, 15}, {4, 8, 12, 16, 20}, {5, 10, 15, 20, 25}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
