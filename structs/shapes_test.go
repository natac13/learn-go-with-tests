package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0, name: "Rectangle"},
		{shape: Circle{Radius: 10}, want: 314.1592653589793, name: "Circle"},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0, name: "Triangle"},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
