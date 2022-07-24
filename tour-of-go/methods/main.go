package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Methods for Vertex type (similar to class methods)
// Value receiver, which creates a copy of struct (similar to pass by value)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receiver which changes the underlying struct (similar to pass by reference)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

// Methods for non-struct type (the definition of methods for MyFloat are only scoped to this package)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	v := Vertex{3, 4}

	// Even though `v` is a value of Vertex, still it can call Scale method which has a pointer receiver. Go interprets this internally.
	v.Scale(10)

	fmt.Println(v.Abs())

	// Same logic as above works for value receiver methods called by reference variables.
	fmt.Println((&v).Abs())

	// Pointer receiver is more suitable because
	// 1. Can change underlying values
	// 2. Avoid copying struct while calling

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
