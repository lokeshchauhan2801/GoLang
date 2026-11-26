// Command variadic shows variadic functions and named multiple results.
package main

import "fmt"

// sum accepts any number of ints; inside the function nums is a []int.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// stats returns the smallest and largest value using named results.
// A bare "return" sends back the current values of min and max.
func stats(nums ...int) (min, max int) {
	if len(nums) == 0 {
		return // both are their zero value, 0
	}
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func main() {
	defer fmt.Println("done") // runs last, when main returns

	// Calling with separate arguments.
	fmt.Println("sum(1, 2, 3) =", sum(1, 2, 3))
	fmt.Println("sum(10, 20) =", sum(10, 20))

	// Passing an existing slice: add ... after it.
	s := []int{4, 8, 15, 16, 23, 42}
	fmt.Println("sum(s...) =", sum(s...))

	lo, hi := stats(7, -2, 9, 3)
	fmt.Printf("stats(7, -2, 9, 3) -> min=%d max=%d\n", lo, hi)

	lo, hi = stats(s...)
	fmt.Printf("stats(s...) -> min=%d max=%d\n", lo, hi)
}
