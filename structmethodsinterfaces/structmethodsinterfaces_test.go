package structmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 2.0}
	got := rectangle.Perimeter()
	want := 8.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{2.0, 2.0}
		checkArea(t, rectangle, 4.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTableDriven(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 2.0, Width: 2.0}, hasArea: 4.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 6, Height: 6}, hasArea: 18},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, test.hasArea)
			}
		})
	}
}
