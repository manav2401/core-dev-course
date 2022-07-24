package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

// Vertex denoting co-ordinates
type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer

	fmt.Println(i) // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer

	fmt.Println(j) // see the new value of j

	// Demonstrating struct
	fmt.Println(Vertex{1, 2})

	// Creating a new Vertex object and accessing by value
	v := Vertex{3, 4}
	fmt.Println("Access by value: v.X:", v.X, "v.Y", v.Y)

	// Creating a reference to Vertex Object and only setting partial parameters
	// Others will take the default type of respective datatype
	vPtr := &Vertex{X: 3}
	fmt.Println("Access by reference: vPtr.X:", vPtr.X, "vPtr.Y", vPtr.Y)

	// Fixed size arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// slice
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var p1 []int = primes[1:4] // for index 1, 2, 3

	fmt.Println(p1)

	// omitting upper or lower bounds
	p1 = primes[:4]
	fmt.Println(p1)

	// slice is a reference to an underlying array
	var aSlice []string = a[0:1]
	aSlice[0] = "Hello1" // this should also change a[0]

	fmt.Println(a)

	// initialization without size also creates slice
	a1 := []string{"hello", "world"}
	fmt.Println(a1)

	// Slice contains 2 parameters: length and capacity
	// length: number of elements it contains (till end)
	// capacity: number of elements it contains till end from the first element
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:] // s => 5, 7, 11, 13. So, ideally if we count from first element, it only has 4 elements
	printSlice(s)

	var s1 []int

	fmt.Println(s1, len(s1), cap(s1))

	if s1 == nil {
		fmt.Println("nil slice: s1")
	}

	a2 := make([]int, 5)
	printSlice1("a", a2)

	b := make([]int, 0, 5)
	printSlice1("b", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)

	s = []int{}
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// Iterating over slice. For a slice or map there are 2 values: index and the value itself
	// In case there's no need of any one, they can be skipped by using `_`.
	for index, value := range s {
		fmt.Println("index:", index, "value:", value)
	}

	// Pic implementation
	pic.Show(Pic)

	// Map (key -> value) for any native and custom type
	// Can also be created using slice.
	m := make(map[string]Vertex)
	m["Vertex1"] = Vertex{1, 2}
	m["Vertex2"] = Vertex{3, 4}

	fmt.Println("Iterating over map 1")

	for index, value := range m {
		fmt.Println("index:", index, "value:", value)
	}

	// Initializing map
	m2 := map[string]Vertex{
		"Vertex0": {0, 1},
	}

	fmt.Println("Iterating over map 2")

	for index, value := range m2 {
		fmt.Println("index:", index, "value:", value)
	}

	// Mutating map and Accessing data
	fmt.Println("Before update:", m["Vertex1"])
	m["Vertex1"] = Vertex{10, 10}
	fmt.Println("After update:", m["Vertex1"])

	// Checking presence of key
	elem, ok := m["Vertex3"] // not actually present in map
	fmt.Println("m[Vertex3]:", elem, "present?", ok)

	// Give frequency of each word in a map of a string
	word := "Hello World World Test World Test2 Hello Test"
	fmt.Println("Frequency Map:", WordCount(word))
	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))

	// Function closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)

	for i := 0; i < dx; i++ {
		x := make([]uint8, dy)
		for j := 0; j < dy; j++ {
			x[j] = (uint8)((i + j) / 2)
		}

		pic[i] = x
	}

	return pic
}

func WordCount(s string) map[string]int {
	slice := strings.Fields(s)
	m := make(map[string]int)

	for _, value := range slice {
		m[value] += 1 // It's 0 if it was not present previously
	}

	return m
}

// Demonstrating function passing
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Demonstrating function closures
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0 := 0
	f1 := 1

	return func() int {
		fn := f0
		f0, f1 = f1, f0+f1

		return fn
	}
}
