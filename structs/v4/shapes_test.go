package main

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g and watn %g", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72)

	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
