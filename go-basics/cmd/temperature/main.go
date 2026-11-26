// Command temperature converts a Celsius value given on the
// command line to Fahrenheit.
//
// Usage: go run ./cmd/temperature 37
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args[0] is the program name; the first real argument is os.Args[1].
	if len(os.Args) < 2 {
		fmt.Println("usage: temperature <celsius>   e.g. temperature 37")
		os.Exit(1)
	}

	// Go never converts a string to a number silently:
	// ParseFloat returns the number AND an error that we must check.
	celsius, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("%q is not a valid number, please enter a temperature like 37 or -4.5\n", os.Args[1])
		os.Exit(1)
	}

	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.2f\n", fahrenheit)
}
