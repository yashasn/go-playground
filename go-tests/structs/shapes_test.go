package main

import "testing"

func TestPermiter(t *testing.T) {
	rectangle := Rectangle{10.0, 5.0}
	got := Permiter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	/* Helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle.
	By declaring an interface, the helper is decoupled from the concrete types
	and only has the method it needs to do its job.*/

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		//Use of g will print a more precise decimal number in the error message
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		want := 50.0
		checkArea(t, rectangle, want)

	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

}

func TestAreaTableTCs(t *testing.T) {
	// Table driven tests using anonymous structs
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 5.0}, 50.0},
		{Circle{50.0}, 314.1592653589793},
	}
	//NOTE: %#v used to print structs with all its values
	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("%#v got %g want %g", test.shape, got, test.want)
		}

	}
}
