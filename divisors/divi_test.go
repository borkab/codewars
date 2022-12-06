package kata

import "testing"

func TestDivisors(t *testing.T) {
	n := 7
	got := Divisors(n)
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
