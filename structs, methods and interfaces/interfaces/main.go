package main

import (
	"math"
)

type Circle struct {
	Radius float32
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (rec Rectangle) Area() float32 {
	return rec.Height * rec.Width
}

func (circle Circle) Area() float32 {
	return circle.Radius * circle.Radius * math.Pi
}

func main() {

}
