package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// Basic for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("1st loop result", sum)

	// for acting as while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("2nd loop result", sum)

	// Infinite loop can be created using for {}

	// Demonstrating if conditions in sqrt function
	fmt.Println("sqrt(2): ", sqrt(2), "sqrt(-4): ", sqrt(-4))

	// Demonstrating if condition with a statement to execute before the condition
	// Also demonstrating else condition
	// Scope of variables only for the condition
	fmt.Println("pow(3, 2, 10): ", pow(3, 2, 10), "pow(3, 3, 20): ", pow(3, 3, 20))

	// Demonstrating for loop in Sqrt function using newton's method
	fmt.Println("Sqrt(2) using newton's method: ", Sqrt(2))

	// Switch case
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch without condition, resembles to `switch true {}`
	// and can be used instead of multiple if-else statements
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Calling temp() which has a defer() statement
	temp()

	fmt.Println("Done.")
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

// Sqrt using Newton's method
func Sqrt(x float64) float64 {
	z := 1.0
	temp := z

	for i := 0; i < 10; i++ {
		temp -= ((z*z - x) / (2 * z))
		if math.Abs(temp-z) <= math.Pow(10, -12) {
			break
		}

		z = temp
	}

	return z
}

// Demonstrating multiple defer statements
func temp() {
	fmt.Println("Starting to run temp()")
	defer fmt.Println("Ending temp()")
	defer time.Sleep(time.Second)
	defer fmt.Println("Saving work...")
	fmt.Println("Doing some work...")
	time.Sleep(time.Second)
}
