// Command calc evaluates a simple expression using the calc package.
//
// Usage:
//
//	go run ./cmd/calc 10 / 2
//	go run ./cmd/calc 3 x 4      (or '3 * 4' with quotes: the shell expands a bare *)
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lokeshchauhan2801/GoLang/go-basics/calc"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: calc <number> <+|-|x|/> <number>   e.g. calc 10 / 2")
		os.Exit(2)
	}

	result, err := calc.Eval(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		var ie *calc.InputError
		switch {
		case errors.Is(err, calc.ErrDivideByZero):
			fmt.Println("cannot divide by zero")
		case errors.Is(err, calc.ErrUnknownOperator):
			fmt.Printf("unknown operator %q, use one of + - x /\n", os.Args[2])
		case errors.As(err, &ie):
			fmt.Printf("%q is not a number, please enter something like 10 or 2.5\n", ie.Value)
		default:
			fmt.Println("error:", err)
		}
		os.Exit(1)
	}

	fmt.Println(result)

	// calc.round2(result) would not compile here:
	//   undefined: calc.round2 (cannot refer to unexported name)
	// because lower-case names are private to their package.
}
