// Command sliceshare demonstrates the shared-backing-array trap and its fix.
package main

import (
	"fmt"

	"github.com/lokeshchauhan2801/GoLang/go-basics/collections"
)

func main() {
	// The trap: a slice of an array shares the array's memory.
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3] // [2 3], but it points into arr
	s[0] = 99
	fmt.Println("Trap:")
	fmt.Println("  slice:", s)
	fmt.Println("  array:", arr) // arr changed too

	// The fix: copy into a new slice with its own backing array.
	arr2 := [5]int{1, 2, 3, 4, 5}
	c := make([]int, 2)
	copy(c, arr2[1:3])
	c[0] = 99
	fmt.Println("Fix with copy:")
	fmt.Println("  slice:", c)
	fmt.Println("  array:", arr2) // arr2 unchanged

	fmt.Println("Collections:")
	fmt.Println("  RotateLeft([1 2 3 4 5], 2) =", collections.RotateLeft([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("  RotateLeft([1 2 3 4 5], 7) =", collections.RotateLeft([]int{1, 2, 3, 4, 5}, 7))
	fmt.Println("  Unique([a b a c b])        =", collections.Unique([]string{"a", "b", "a", "c", "b"}))
}
