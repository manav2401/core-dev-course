package helper

// SortSlice is a generic sort function which uses insertion sort
// to perform ascending and descending sort on a slice of int or string
func SortSlice[T int | string](s []T, reverse bool) []T {
	// insertion sort
	for i := 1; i < len(s); i++ {
		for j := i; (!reverse && j > 0 && s[j-1] > s[j]) || (reverse && j > 0 && s[j-1] < s[j]); j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}

	return s
}
