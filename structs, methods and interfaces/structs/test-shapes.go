package main

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.00, 10.00}
	got := Perimeter(rec)
	chck := 40.00

	if got != chck {
		t.Errorf("got %.2f want %.2f", got, chck)
	}
}

func TestArea(t *testing.T) {

	t.Run(
		"Rectangles",
		func(t *testing.T) {
			rec := Rectangle{12.0, 6.00}
			got := rec.Area()
			chck := 72.00

			if got != chck {
				t.Errorf("got %.2f want %.2f", got, chck)
			}
		},
	)

	t.Run(
		"Circle",
		func(t *testing.T) {
			circle := Circle{10}
			got := circle.Area()
			want := 314.1592653589793

			if got != want {
				t.Errorf("got %g want %g", got, want)
			}
		},
	)
}
