package perimeter

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			"Rectangle",
			Rectangle{Width: 12, Height: 6,},
			72.0,
		},
		{
			"Circle",
			Circle{Radius: 10},
			314.1592653589793,
		},
		{
			"Triangle",
			Triangle{Base: 12, Height: 6},
			36.0,
		},
	}

	for _, tc := range areaTests {
		got := tc.shape.Area()
		want := tc.want
		if got != want {
			t.Errorf("%#v got %g want %g", tc.name, got, tc.want)
		}
	}

}
