package kata

import "testing"

func TestSolution(t *testing.T) {

	t.Run("the ending is the same", func(t *testing.T) {
		str := "amimako"
		ending := "ako"

		got := solution(str, ending)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("the ending is NOT the same", func(t *testing.T) {
		str := "amimako"
		ending := "puh"

		got := solution(str, ending)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
