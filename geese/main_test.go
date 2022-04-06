package main

import (
	"reflect"
	"testing"
)

func TestUngeeser(t *testing.T) {
	all := Birds{"Mallard", "Hook Bill", "African", "Crested", "Pilgrim", "Toulouse", "Blue Swedish"}
	geese := Birds{"African", "Roman Tufted", "Toulouse", "Pilgrim", "Steinbacher"}
	want := Birds{"Mallard", "Hook Bill", "Crested", "Blue Swedish"}

	if got := Ungeeser(all, geese); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v want %#v", got, want)
	}
}

func TestUngeeser_NoGeeseInAllBirds(t *testing.T) {
	all := Birds{"A", "B"}
	geese := Birds{"C", "D", "E", "F"}
	want := Birds{"A", "B"}

	if got := Ungeeser(all, geese); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v want %#v", got, want)
	}
}
