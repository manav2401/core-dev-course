package main

import (
	"fmt"
	"math/rand"

	helper "github.com/manav2401/core-dev-course/helpers"
)

func main() {
	// initialize a string and print occurrences of each word in it
	str := "core dev course hello dev world dev core"
	printWordOccurrence(str)

	// initialize a slice of integers
	slice0 := generateRandomSlice(10)
	printOccurringNumbers(slice0)

	// declare a new slice
	slice1 := generateRandomSlice(15)

	// find overlapping numbers in slice0 and slice1
	slice2 := findOverlappingNumbers(slice0, slice1)
	fmt.Printf("Overlapping numbers in slice0: %v and slice1: %v are %v\n\n", slice0, slice1, slice2)

	// find nth fibonacci number using memoization in maps
	val1, m := fibonacci(5, make(map[int]int))
	fmt.Printf("fibonacci(5): %d\n", val1)

	val2, m := fibonacci(8, m)
	fmt.Printf("fibonacci(8): %d\n", val2)

	val3, _ := fibonacci(4, m)
	fmt.Printf("fibonacci(4): %d\n\n", val3)

	// create a random multi dimensional map and print the sorted version of it
	// map[fee]map[nonce]exists
	m2 := generateMultiDimensionalMap(3, 3)
	printSortedMultiDimensionalMap(m2)
}

// generateRandomSlice fills random numbers in range 0-9 till length `l`
func generateRandomSlice(l int) []int {
	s := make([]int, l)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(10)
	}

	return s
}

// generateMultiDimensionalMap generates a nested map with random numbers
func generateMultiDimensionalMap(l1, l2 int) map[int]map[int]struct{} {
	m := make(map[int]map[int]struct{})

	for i := 0; i < l1; i++ {
		idx := rand.Intn(100)
		for {
			// check if the index already exists
			if _, ok := m[idx]; ok {
				idx = rand.Intn(100)
			} else {
				break
			}
		}
		m[idx] = make(map[int]struct{})
		for j := 0; j < l2; j++ {
			// it's okay to override an existing sub-index
			m[idx][rand.Intn(100)] = struct{}{}
		}
	}

	return m
}

// printWordOccurrence prints occurrence of each word in a string
func printWordOccurrence(str string) {
	m := make(map[string]int, len(str))

	s := []byte{}

	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			m[string(s)]++
			s = []byte{}
		} else {
			s = append(s, str[i])
		}
	}
	m[string(s)]++

	fmt.Printf("Occurrence of each word in string: %s\n", str)

	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Printf("\n")
}

// printOccuringNumbers prints numbers occurring atleast once in the slice
func printOccurringNumbers(s []int) {
	m := make(map[int]struct{}, len(s))

	for i := 0; i < len(s); i++ {
		m[s[i]] = struct{}{}
	}

	fmt.Printf("Occurrence of each number in slice: %v\n", s)

	for key := range m {
		fmt.Printf("%d ", key)
	}

	fmt.Printf("\n")
}

// findOverlappingNumbers finds the overlapping numbers in s1 and s2
func findOverlappingNumbers(s1 []int, s2 []int) []int {
	m := make(map[int]struct{}, len(s1))

	for i := 0; i < len(s1); i++ {
		m[s1[i]] = struct{}{}
	}

	min := len(s1)
	if len(s2) < len(s1) {
		min = len(s2)
	}

	ans := make([]int, 0, min)
	for i := 0; i < len(s2); i++ {
		if len(ans) == min {
			break
		}

		if _, ok := m[s2[i]]; ok {
			ans = append(ans, s2[i])
		}
	}

	return ans
}

// fibonacci uses map for memoization of values for seen indexes.
// It also returns the map for future use.
func fibonacci(n int, m map[int]int) (int, map[int]int) {
	if len(m) == 0 {
		m[0] = 0
		m[1] = 1
	}

	// check if we have the result memoized
	if val, ok := m[n]; ok {
		return val, m
	}

	val1, _ := fibonacci(n-1, m)
	val2, _ := fibonacci(n-2, m)
	m[n] = val1 + val2
	return m[n], m
}

// sortMultiDimensionalMap sorts a nested map in ascending for the first order keys
// and in descending order for second order keys
func printSortedMultiDimensionalMap(m map[int]map[int]struct{}) {
	// collect the first order keys
	firstOrderKeys := make([]int, 0, len(m))
	for k := range m {
		firstOrderKeys = append(firstOrderKeys, k)
	}

	// sort the first order keys
	firstOrderKeys = helper.SortSlice(firstOrderKeys, false)

	fmt.Printf("Map before sorting: %v\n\n", m)
	for _, v := range firstOrderKeys {
		// collect the second order keys
		secondOrderKeys := make([]int, 0, len(m[v]))
		for k := range m[v] {
			secondOrderKeys = append(secondOrderKeys, k)
		}

		// sort the second order keys
		secondOrderKeys = helper.SortSlice(secondOrderKeys, true)

		// print
		fmt.Printf("Key: %d\nValues: ", v)
		for _, v1 := range secondOrderKeys {
			fmt.Printf("%d ", v1)
		}
		fmt.Printf("\n\n")
	}
}
