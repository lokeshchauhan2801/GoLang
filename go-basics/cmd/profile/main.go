// Command profile prints a short personal profile using
// var, short declaration (:=) and const.
package main

import "fmt"

// const: a value fixed at compile time; it can never change.
const city = "Jaipur"

func main() {
	// var with an explicit type.
	var name string = "Lokesh Chauhan"

	// := short declaration; the type is inferred (int and float64).
	age := 25
	height := 1.75 // metres

	// One sentence with Printf: %s string, %d integer, %.2f float with 2 decimals.
	fmt.Printf("My name is %s, I am %d years old, %.2f m tall and I live in %s.\n",
		name, age, height, city)

	// %T prints the Go type of each value.
	fmt.Printf("Types: name=%T, age=%T, height=%T, city=%T\n", name, age, height, city)
}
