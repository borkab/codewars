package kata

import (
	"reflect"
	"sort"
	"testing"
)

func TestMinMax(t *testing.T) {
	fullarray := []int{3, 7, 4, 11, 2, 5, 9, 1}
	type minmax [2]int
	got := MinMax(sort.IntSlice(fullarray))
	want := [2]int(minmax{1, 11})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
