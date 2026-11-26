// Command safediv divides numbers and returns an error instead of
// crashing or printing Inf when dividing by zero.
package main

import (
	"errors"
	"fmt"
)

// divide returns a result plus an error: Go's most common pattern.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	defer fmt.Println("done") // runs last, when main returns

	pairs := [][2]float64{{10, 2}, {7, 3}, {5, 0}}

	for _, p := range pairs {
		result, err := divide(p[0], p[1])
		if err != nil {
			fmt.Printf("%g / %g: error: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("%g / %g = %.2f\n", p[0], p[1], result)
	}
}
