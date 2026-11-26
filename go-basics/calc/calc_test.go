package calc

import (
	"errors"
	"testing"
)

func TestBasicOps(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add = %v, want 5", got)
	}
	if got := Sub(2, 3); got != -1 {
		t.Errorf("Sub = %v, want -1", got)
	}
	if got := Mul(2, 3); got != 6 {
		t.Errorf("Mul = %v, want 6", got)
	}
	if got, err := Divide(10, 4); err != nil || got != 2.5 {
		t.Errorf("Divide(10, 4) = %v, %v, want 2.5, nil", got, err)
	}
}

func TestDivideByZeroIsWrapped(t *testing.T) {
	_, err := Divide(10, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("errors.Is(err, ErrDivideByZero) = false for %v", err)
	}
	if err.Error() == ErrDivideByZero.Error() {
		t.Error("error should carry context, not just the sentinel text")
	}
}

func TestInputError(t *testing.T) {
	_, err := Eval("ten", "/", "2")
	var ie *InputError
	if !errors.As(err, &ie) {
		t.Fatalf("errors.As failed for %v", err)
	}
	if ie.Value != "ten" {
		t.Errorf("Value = %q, want %q", ie.Value, "ten")
	}
}

func TestEval(t *testing.T) {
	tests := []struct {
		l, op, r string
		want     float64
		wantErr  error
	}{
		{"10", "/", "2", 5, nil},
		{"10", "/", "3", 3.33, nil},
		{"2", "x", "4", 8, nil},
		{"2", "*", "4", 8, nil},
		{"7", "-", "10", -3, nil},
		{"1", "+", "2", 3, nil},
		{"10", "/", "0", 0, ErrDivideByZero},
		{"1", "%", "2", 0, ErrUnknownOperator},
	}
	for _, tt := range tests {
		got, err := Eval(tt.l, tt.op, tt.r)
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("Eval(%s %s %s) err = %v, want %v", tt.l, tt.op, tt.r, err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("Eval(%s %s %s) = %v, want %v", tt.l, tt.op, tt.r, got, tt.want)
		}
	}
}

func TestRound2(t *testing.T) {
	if got := round2(3.14159); got != 3.14 {
		t.Errorf("round2 = %v, want 3.14", got)
	}
}
