package main

import "testing"

func TestArea(t *testing.T) {
	//TableDrivenTests
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{5.0, 4.0}, 10.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g and want %g", got, tt.want)
		}
	}
}
