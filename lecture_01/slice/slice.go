package main

import (
	"fmt"
)

func main() {
	// initialize a fixed slice
	slice0 := []int{1, 2, 3}
	fmt.Printf("Slice0 initialized: %v\n\n", slice0)

	// add a constant value to each element
	addConstantToEachElement(slice0, 5)
	fmt.Printf("Adding constant value to each element in Slice0: %v\n", slice0)

	// append a new value to the end of slice
	slice0 = appendAtEnd(slice0, 9)
	fmt.Printf("Appending constant value at the end in Slice0: %v\n", slice0)

	// append a new value at the start of slice
	slice0 = appendAtStart(slice0, 5)
	fmt.Printf("Appending constant value at the start in Slice0: %v\n\n", slice0)

	// remove the last element from slice
	val, slice0 := removeElement(slice0, len(slice0)-1)
	if val == nil {
		fmt.Printf("Unable to remove element from end\n")
	} else {
		fmt.Printf("Removing last element (%d) in Slice0. Updated slice: %v\n", *val, slice0)
	}

	// remove the first element from slice
	val, slice0 = removeElement(slice0, 0)
	if val == nil {
		fmt.Printf("Unable to remove element from start\n")
	} else {
		fmt.Printf("Removing first element (%d) in Slice0. Updated slice: %v\n", *val, slice0)
	}

	// remove the element with index `1` from slice
	index := 1
	val, slice0 = removeElement(slice0, index)

	if val == nil {
		fmt.Printf("Unable to remove element from %d\n", index)
	} else {
		fmt.Printf("Removing element (%d) from index (%d) in Slice0. Updated slice: %v\n\n", *val, index, slice0)
	}

	// declaring a new slice for appending
	slice1 := []int{10, 8, 11, 12}

	// append slice0 and slice1 and remove duplicates
	slice2 := appendAndRemoveDuplicates(slice0, slice1)
	fmt.Printf("Appending slice0 (%v) and slice1 (%v): %v\n", slice0, slice1, slice2)

	// declaring a new slice
	slice3 := []int{8, 11, 13}
	fmt.Printf("Removing slice3 elements (%v) from slice2 (%v): ", slice3, slice2)
	slice4 := removeOverlappingElements(slice2, slice3)
	fmt.Printf("%v\n\n", slice4)

	// declaring a new slice and rotating left by 1
	slice5 := []int{3, 4, 5, 6, 7, 8}
	fmt.Printf("Slice5 before rotating left  by 1: %v\n", slice5)
	slice5 = rotateLeft(slice5, 1)
	fmt.Printf("Slice5 after  rotating left  by 1: %v\n", slice5)

	// rotating left by 3
	slice5 = rotateLeft(slice5, 3)
	fmt.Printf("Slice5 after  rotating left  by 3: %v\n", slice5)

	// rotating right by 1
	slice5 = rotateRight(slice5, 1)
	fmt.Printf("Slice5 after  rotating right by 1: %v\n", slice5)

	// rotating right by 3
	slice5 = rotateRight(slice5, 2)
	fmt.Printf("Slice5 after  rotating right by 2: %v\n\n", slice5)

	// get copy of slice5
	slice6 := createCopy(slice5)
	fmt.Printf("Slice6 after copying from slice5: %v\n", slice6)

	// swapping odd and even elements of the slice
	slice6 = swapElements(slice6)
	slice7 := slice6
	fmt.Printf("Slice6 after swapping odd and even elements: %v\n\n", slice6)

	// sorting the slice in ascending order
	fmt.Printf("Slice6 before: %v and ", slice6)
	slice6 = sortSlice(slice6, false)
	fmt.Printf("after sorting in ascending  order: %v\n", slice6)

	// sorting the slice in descending order
	fmt.Printf("Slice7 before: %v and ", slice7)
	slice7 = sortSlice(slice7, true)
	fmt.Printf("after sorting in descending order: %v\n\n", slice7)

	// creating a new slice of strings
	str := []string{"hello", "world", "core", "dev", "course"}
	str2 := str

	// sorting the string slice in ascending order using the same generic sort function
	fmt.Printf("String slice before: %v and ", str)
	str = sortSlice(str, false)
	fmt.Printf("after sorting in ascending  order: %v\n", str)

	// sorting the string slice in descending order using the same generic sort function
	fmt.Printf("String slice before: %v and ", str2)
	str2 = sortSlice(str2, true)
	fmt.Printf("after sorting in descending order: %v\n", str2)
}

// addConstantToEachElement adds a constant value `x` to
// each element of slice `s`. As we're just updating the
// underlying array values, we don't need to return back s.
func addConstantToEachElement(s []int, x int) {
	for index := range s {
		s[index] += x
	}
}

// appendAtEnd appends a value `x` to the slice `s` at the end.
func appendAtEnd(s []int, x int) []int {
	s = append(s, x)
	return s
}

// appendAtStart appends a value `x` to the slice `s` at the start.
func appendAtStart(s []int, x int) []int {
	if len(s) == 0 {
		return append(s, x)
	}

	// append 2 slices which are s[0] and s[0:],
	// which basically copies the 0th element twice
	s = append(s[:1], s[0:]...)

	// update the 0th element with `x`
	s[0] = x

	return s
}

// removeElement removes an element from the slice given an index.
// It returns the removed element and slice.
func removeElement(s []int, index int) (*int, []int) {
	if len(s) == 0 && index >= 0 && index < len(s) {
		return nil, nil
	}

	val := s[index]
	s = append(s[:index], s[index+1:]...)

	return &val, s
}

// appendAndRemoveDuplicates acceps 2 slices, appends them and
// returns a new slice with unique elements.
func appendAndRemoveDuplicates(s0 []int, s1 []int) []int {
	// create an output slice to return
	out := make([]int, 0, len(s0)+len(s1))

	// create a map for keeping track of duplicates
	m := make(map[int]bool)

	for _, val := range s0 {
		if _, ok := m[val]; !ok {
			out = append(out, val)
			m[val] = true
		}
	}

	for _, val := range s1 {
		if _, ok := m[val]; !ok {
			out = append(out, val)
			m[val] = true
		}
	}

	return out
}

// removeOverlappingElements removes all elements from `s0` if
// they're in `s1`.
func removeOverlappingElements(s0 []int, s1 []int) []int {
	// create a map for keeping tracks of elements of second slice
	m := make(map[int]bool)

	for _, val := range s1 {
		m[val] = true
	}

	for i := 0; i < len(s0); i++ {
		if _, ok := m[s0[i]]; ok {
			_, s0 = removeElement(s0, i)
			i--
		}
	}

	return s0
}

// rotateLeft rotates the given slice `s` to left `i` times
func rotateLeft(s []int, i int) []int {
	if len(s) < 2 || i < 1 || i >= len(s) {
		return s
	}

	s = append(s[i:], s[:i]...)

	return s
}

// rotateRight rotates the given slice `s` to right `i` times
func rotateRight(s []int, i int) []int {
	if len(s) < 2 || i < 1 || i >= len(s) {
		return s
	}

	s = append(s[len(s)-i:], s[:len(s)-i]...)

	return s
}

// createCopy creates a deep copy of given slice
func createCopy(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)

	return c
}

// swapElements swaps the odd and even indexed elements
func swapElements(s []int) []int {
	if len(s) < 2 {
		return s
	}

	l := len(s)
	if len(s)%2 != 0 {
		l--
	}

	for i := 0; i < l; i += 2 {
		s[i], s[i+1] = s[i+1], s[i]
	}

	return s
}

// sortSlice is a generic sort function which uses insertion sort
// to perform ascending and descending sort on a slice of int or string
func sortSlice[T int | string](s []T, reverse bool) []T {
	// insertion sort
	for i := 1; i < len(s); i++ {
		for j := i; (!reverse && j > 0 && s[j-1] > s[j]) || (reverse && j > 0 && s[j-1] < s[j]); j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}

	return s
}
