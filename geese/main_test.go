package main

import (
	"reflect"
	"testing"
)

func TestUngeeser(t *testing.T) {

	want := nogeese

	if got := Ungeeser(allbirds); !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}

}
