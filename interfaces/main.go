package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface { // here we defined an interface
	Area() float32 // Interface method
}

// The interface implementation is an Implicit One and does not need to be declared
// if the struct has the same members as the interface, then the struct implements the interface

// Circle is a  struct that implement the Shape interface
type Circle struct {
	Radius float32
}

// Area is the method that implements the interface
func (circle Circle) Area() float32 {
	return circle.Radius * circle.Radius * math.Pi
}

// Rectangle is a struct that implements the Shape interface
type Rectangle struct {
	Height float32
	Width  float32
}

// Area is the method that implements the interface
func (rec Rectangle) Area() float32 {
	return rec.Height * rec.Width
}

// This function is a wrapper for the Area method of the Shape interface and so shapes rectangle and circle would work
func CalculateArea(s Shape) float32 {
	return s.Area()
}

// function for getting value and type of interfaces to prove what type it holds and what values
func describeValue(t interface{}) {
	fmt.Println("Type of value: ", reflect.TypeOf(t))
	fmt.Println("Value of type: ", t)
}

func main() {
	rec := Rectangle{Height: 10.00, Width: 5.00}
	circle := Circle{Radius: 10.00}

	fmt.Println("Rectangle Area: ", CalculateArea(rec)) // this validates our theory that the Rectangle struct implements the Shape interface
	fmt.Println("Circle Area: ", CalculateArea(circle)) // this validates our theory that the Circle struct implements the Shape interface

	// interface type can hold any type of value
	MysteryBox := interface{}(12)

	describeValue(MysteryBox)

	value, ok := MysteryBox.(Circle) // this .(int) or .(string) or in this case .(Circle) checks if value is same type
	// similar to .isalnum() or .isalpha() in python
	if ok {
		fmt.Println("MysteryBox is a Circle")
		fmt.Println("MysteryBox Area: ", value.Area())
	} else {
		fmt.Println("MysteryBox is not a Circle")
	}

	// here rec became a Geometry interface as it implements both the Shape and Measurable interfaces
	describeShape(rec)

	PerformCalculation(-28.99) // this will throw an error

}

// implementing interface composition or embedding

// this is another interface
type Measurable interface {
	Perimeter() float32
}

// Rectangle is a struct that implements the Geometry interface as it implements the Shape interface and the Measurable interface
func (rec Rectangle) Perimeter() float32 {
	return rec.Height + rec.Width
}

// Circle is a struct that implements the Geometry interface as it implements the Shape interface and the Measurable interface
func (circle Circle) Perimeter() float32 {
	return 2 * circle.Radius * math.Pi
}

// this interface uses two embedded interfaces
type Geometry interface {
	Shape // shape is an interface and so is Measurable
	Measurable
}

func describeShape(g Geometry) {
	fmt.Println("Shape is a:", reflect.TypeOf(g).String())
	fmt.Println("Perimeter is:", g.Perimeter())
}

// Implementing error interface using this struct

type CalculationError struct {
	msg string
}

// implement Error() func for this struct
func (e CalculationError) Error() string {
	return e.msg
}

// trying to throw an error and use implemented struct error
func PerformCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{msg: "invalid value"}
	}
	return math.Sqrt(val), nil
}
