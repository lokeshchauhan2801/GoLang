package shapes

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

// Compile-time checks that every type satisfies Shape.
var (
	_ Shape = Rect{}
	_ Shape = Circle{}
	_ Shape = Triangle{}
)

func TestAreaAndPerimeter(t *testing.T) {
	tests := []struct {
		shape     Shape
		area, per float64
	}{
		{Rect{W: 2, H: 3}, 6, 10},
		{Circle{R: 1}, math.Pi, 2 * math.Pi},
		{Triangle{A: 3, B: 4, C: 5}, 6, 12},
	}
	for _, tt := range tests {
		if got := tt.shape.Area(); !almostEqual(got, tt.area) {
			t.Errorf("%v Area() = %v, want %v", tt.shape, got, tt.area)
		}
		if got := tt.shape.Perimeter(); !almostEqual(got, tt.per) {
			t.Errorf("%v Perimeter() = %v, want %v", tt.shape, got, tt.per)
		}
	}
}

func TestString(t *testing.T) {
	r := Rect{Named{"door"}, 2, 1}
	if got, want := r.String(), `Rect "door" 2.00 x 1.00`; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestScaleQuadruplesArea(t *testing.T) {
	r := Rect{W: 2, H: 1}
	c := Circle{R: 1}
	tr := Triangle{A: 3, B: 4, C: 5}
	before := []float64{r.Area(), c.Area(), tr.Area()}

	r.Scale(2)
	c.Scale(2)
	tr.Scale(2)

	after := []float64{r.Area(), c.Area(), tr.Area()}
	for i := range before {
		if !almostEqual(after[i], 4*before[i]) {
			t.Errorf("shape %d: area %v after Scale(2), want %v", i, after[i], 4*before[i])
		}
	}
}

func TestLargest(t *testing.T) {
	small := Rect{W: 1, H: 1}
	big := Circle{R: 5}
	got := Largest([]Shape{small, big, Triangle{A: 3, B: 4, C: 5}})
	if got != big {
		t.Errorf("Largest = %v, want %v", got, big)
	}
	if Largest(nil) != nil {
		t.Error("Largest(nil) should be nil")
	}
}

func TestKind(t *testing.T) {
	cases := map[string]Shape{
		"rectangle":        Rect{W: 2, H: 1},
		"square rectangle": Rect{W: 2, H: 2},
		"circle":           Circle{R: 1},
		"triangle":         Triangle{A: 3, B: 4, C: 5},
		"nil shape":        nil,
	}
	for want, s := range cases {
		if got := Kind(s); got != want {
			t.Errorf("Kind(%v) = %q, want %q", s, got, want)
		}
	}
}
