package kata

import "testing"

func TestToJadenStrings(t *testing.T) {
	str := "good morning sunshine"
	got := ToJadenCase(str)
	want := "Good Morning Sunshine"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
