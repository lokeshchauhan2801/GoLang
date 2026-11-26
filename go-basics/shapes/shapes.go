// Package shapes defines geometric shapes that share a Shape interface.
package shapes

import (
	"fmt"
	"math"
)

// Shape is satisfied by any type that has both methods below.
// No type declares "implements Shape": Go checks the method set implicitly.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Named is embedded in every shape. Its Name field is promoted,
// so rect.Name works as if Name were declared on Rect itself.
type Named struct {
	Name string
}

// Rect is a rectangle of width W and height H.
type Rect struct {
	Named
	W, H float64
}

// Circle is a circle of radius R.
type Circle struct {
	Named
	R float64
}

// Triangle is a triangle with side lengths A, B and C.
type Triangle struct {
	Named
	A, B, C float64
}

// ---- Area and Perimeter: value receivers, because they only read the shape. ----

func (r Rect) Area() float64      { return r.W * r.H }
func (r Rect) Perimeter() float64 { return 2 * (r.W + r.H) }

func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

// Area uses Heron's formula.
func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

// ---- String: fmt calls this automatically because it satisfies fmt.Stringer. ----

func (r Rect) String() string {
	return fmt.Sprintf("Rect %q %.2f x %.2f", r.Name, r.W, r.H)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle %q r=%.2f", c.Name, c.R)
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle %q %.2f, %.2f, %.2f", t.Name, t.A, t.B, t.C)
}

// ---- Scale: pointer receivers. ----
//
// Scale must change the shape it is called on. With a value receiver
// (func (r Rect) Scale) Go would pass a COPY of the shape, the method would
// scale the copy, and the copy would be thrown away when the method returns,
// so the caller's shape would stay the same size. A pointer receiver gets the
// address of the original, so the change is visible to the caller.

// Scale multiplies both sides by f, so the area grows by f*f.
func (r *Rect) Scale(f float64) {
	r.W *= f
	r.H *= f
}

// Scale multiplies the radius by f.
func (c *Circle) Scale(f float64) {
	c.R *= f
}

// Scale multiplies every side by f.
func (t *Triangle) Scale(f float64) {
	t.A *= f
	t.B *= f
	t.C *= f
}

// Largest returns the shape with the biggest area, or nil for an empty slice.
func Largest(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}
	largest := shapes[0]
	for _, s := range shapes[1:] {
		if s.Area() > largest.Area() {
			largest = s
		}
	}
	return largest
}

// Kind uses a type switch to recover the concrete type behind a Shape.
func Kind(sh Shape) string {
	switch s := sh.(type) {
	case Rect:
		if s.W == s.H {
			return "square rectangle"
		}
		return "rectangle"
	case Circle:
		return "circle"
	case Triangle:
		return "triangle"
	case nil:
		return "nil shape"
	default:
		return fmt.Sprintf("unknown shape %T", s)
	}
}
