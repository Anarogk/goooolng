package main

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (rectangle Rectangle) Area() float64 {
	return (rectangle.Height * rectangle.Width)
}

func (circle Circle) Area() float64 {
	return math.Pi * (circle.Radius * circle.Radius)
}
