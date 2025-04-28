package main

import "math"

/*
In Go interface resolution is implicit.
If the type you pass in matches what the interface is asking for, it will compile

Rectangle has Area() method and returns float64. Similar to Circl
*/
type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

// Creating a method for Type Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Creating a method for Type Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// normal function not associated to a type
func Permiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
