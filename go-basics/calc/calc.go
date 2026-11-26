// Package calc is a small calculator whose errors callers can inspect
// with errors.Is and errors.As.
package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Sentinel errors: fixed values callers compare against with errors.Is.
var (
	ErrDivideByZero    = errors.New("cannot divide by zero")
	ErrUnknownOperator = errors.New("unknown operator")
)

// InputError reports an operand that is not a valid number.
// Callers read the bad value with errors.As.
type InputError struct {
	Value string
}

func (e *InputError) Error() string {
	return fmt.Sprintf("%q is not a valid number", e.Value)
}

// Add returns a + b.
func Add(a, b float64) float64 { return a + b }

// Sub returns a - b.
func Sub(a, b float64) float64 { return a - b }

// Mul returns a * b.
func Mul(a, b float64) float64 { return a * b }

// Divide returns a / b, or an error wrapping ErrDivideByZero when b is 0.
// %w (not %v) keeps ErrDivideByZero in the chain so errors.Is still finds it.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("calc.Divide(%v, %v): %w", a, b, ErrDivideByZero)
	}
	return a / b, nil
}

// ParseOperand converts s to a number, returning *InputError on failure.
func ParseOperand(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, &InputError{Value: s}
	}
	return v, nil
}

// Eval parses both operands, applies op (+, -, *, x or /) and
// returns the result rounded to two decimal places.
func Eval(left, op, right string) (float64, error) {
	a, err := ParseOperand(left)
	if err != nil {
		return 0, err
	}
	b, err := ParseOperand(right)
	if err != nil {
		return 0, err
	}

	var result float64
	switch op {
	case "+":
		result = Add(a, b)
	case "-":
		result = Sub(a, b)
	case "*", "x":
		result = Mul(a, b)
	case "/":
		result, err = Divide(a, b)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("calc.Eval: %q: %w", op, ErrUnknownOperator)
	}
	return round2(result), nil
}

// round2 rounds to two decimal places. It starts with a lower-case letter,
// so it is unexported: only code inside package calc can call it.
func round2(v float64) float64 {
	return math.Round(v*100) / 100
}
