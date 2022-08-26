package main

import (
	"fmt"
	"math/rand"
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
	fmt.Printf("Overlapping numbers in slice0: %v and slice1: %v are %v\n", slice0, slice1, slice2)
}

func generateRandomSlice(l int) []int {
	s := make([]int, l)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(10)
	}

	return s
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
