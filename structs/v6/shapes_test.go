package main

import "testing"

func TestArea(t *testing.T) {
	//adding more details
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 5.0, Height: 4.0}, hasArea: 10.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("#%v got %g and want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
