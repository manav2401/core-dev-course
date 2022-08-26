# Exercises
You'd need to write a function(s) that solves the task. Tests are required, please set up your CI to check test coverage threshold at least 70%.
Do tasks as one PR to your repo.
Try to keep code readable and without obvious performance issues like algorithms O(N^3) computation complexity.
Try to use only standard golang packages, reuse your own code.
Don't forget about nil, empty or zero cases. :)

## Slices
1. Add given int(N) to each of []int elements.
2. Add the number int(N) to the end of the slice.
3. Add the number int(N) to the beginning of the slice.
4. Take the last number of slice, return it to the user, and remove this element from slice.
5. Take the first slice number, return it to the user, and delete this element from slice.
6. Take the i-th number of slice, return it to the user, and delete this element from slice. The number i is passed by the user to the function. Try to do it without allocating a new slice. 
7. Merge two slices and return a new one with all elements of the first and second without duplicates.
8. From the first slice, remove all the numbers that are in the second.
9. Shift all slice elements by 1 to the left. Zero index element becomes the last one, the first - zero, the last - penultimate.
10. Same, but shift by user-specified i.
11. Same as 9, but right shift.
12. Same, but shift by i.
13. Return a copy of the passed slice to the user
14. swap all even index elements with the nearest odd indices. eg. 0 and 1, 2 and 3, 4 and 5 etc.
15. Sort slice in the given order: direct, reverse, lexicographic (you'd need at least []int and []string).

## Maps
1. There is a text, you need to count how many times each word occurs.
2. There is a very large array or slice of integers, it must be said which numbers are mentioned in it at least once.
3. There are two large arrays of numbers, you need to find which numbers are mentioned in both.
4. Fibonacci (`Fibonacci(n)`) with memoization (https://en.wikipedia.org/wiki/Memoization).
5. you have map[fee]map[nonce], you need to print it as sorted data: descending sorting by fee first and ascending by nonce last.  
