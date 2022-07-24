package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// exported to other packages/modules
// if the struct is named as `test`, it will be only accessible in same package
type Test struct {
	test int
}

func main() {
	// random number generation, using current unix time stamp as the seed value
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))

	// using external library, consuming exported variables/methods
	fmt.Println("pi:", math.Pi)

	// simple function call
	fmt.Println("add(3, 5): ", add(3, 5))

	// multi return function call
	a, b := swap("hello", "world")
	fmt.Println("swap(hello, world): ", a, b)

	// multi return function call
	x, y := split(17)
	fmt.Println("split(17): ", x, y)

	// data types and initialisation
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// using exported struct
	test := Test{test: 1}
	fmt.Println("Test struct: ", test)

	// typecasting
	var (
		i int     = 42
		j float64 = float64(i)
		k uint64  = uint64(j)
	)

	fmt.Println("Typecasted values (int, float64, uint64): ", i, j, k)

	// simple assignment
	l := i // also an int
	fmt.Printf("i's type: %T and l's type: %T\n", i, l)

	k1 := 2.2 // takes float type based on precision
	k2 := 2   // takes int type based on precision
	fmt.Printf("k1's type: %T and k2's type: %T\n", k1, k2)

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println("needInt(Small): ", needInt(Small))
	fmt.Println("needFloat(Small): ", needFloat(Small))
	fmt.Println("needFloat(Big):", needFloat(Big))

	// Float can handle large numbers but int can only hold 64-bit integer (also depending on system arch)
	// Hence it's not possible to run needInt(Big) as it will throw an overflow error.
	fmt.Println("needInt(Big): cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows)")

	fmt.Println("Done.")
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// initialising in function definition itself and using naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
