// Command fizzbuzz prints 1 to 100, replacing multiples of 3 with Fizz,
// multiples of 5 with Buzz and multiples of both with FizzBuzz.
package main

import "fmt"

func main() {
	defer fmt.Println("done") // runs last, when main returns

	for i := 1; i <= 100; i++ {
		// A value-less switch is a cleaner if/else-if chain.
		// Cases are checked top to bottom and the first match wins,
		// so i%15 must come first or 15 would print "Fizz".
		// Go cases do not fall through: no break is needed.
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
