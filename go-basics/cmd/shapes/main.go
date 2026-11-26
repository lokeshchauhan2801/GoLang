// Command shapes demonstrates the shapes package: interfaces, embedding,
// Stringer, pointer receivers and type switches.
package main

import (
	"fmt"

	"github.com/lokeshchauhan2801/GoLang/go-basics/shapes"
)

func main() {
	door := shapes.Rect{Named: shapes.Named{Name: "door"}, W: 2, H: 1}
	wheel := shapes.Circle{Named: shapes.Named{Name: "wheel"}, R: 0.5}
	roof := shapes.Triangle{Named: shapes.Named{Name: "roof"}, A: 3, B: 4, C: 5}

	// 1. A []Shape can hold all three because each has Area and Perimeter.
	all := []shapes.Shape{door, wheel, roof}
	fmt.Println("Areas:")
	for _, s := range all {
		// 2. fmt.Println(s) calls s.String() automatically (fmt.Stringer).
		fmt.Printf("  %-34v area=%6.2f perimeter=%6.2f\n", s, s.Area(), s.Perimeter())
	}

	// Embedding promotes fields: door.Name is really door.Named.Name.
	fmt.Println("\nPromoted field: door.Name =", door.Name)

	// 3. Scale has a pointer receiver; Go takes &door for us automatically.
	before := door.Area()
	door.Scale(2)
	fmt.Printf("\nAfter door.Scale(2): %v, area %.2f -> %.2f (x%.0f)\n",
		door, before, door.Area(), door.Area()/before)

	// 4. Largest and the type switch.
	all = []shapes.Shape{door, wheel, roof} // rebuild: the slice held a copy of the old door
	fmt.Println("\nLargest:", shapes.Largest(all))

	fmt.Println("\nKinds:")
	for _, s := range all {
		fmt.Printf("  %-34v is a %s\n", s, shapes.Kind(s))
	}
}
