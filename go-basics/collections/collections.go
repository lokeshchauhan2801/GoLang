// Package collections holds small helpers for slices.
package collections

// RotateLeft returns a new slice with the elements of s moved n places to the left.
// n larger than len(s) wraps around, so RotateLeft(s, 7) on 5 items equals RotateLeft(s, 2).
// The input slice is not modified.
func RotateLeft(s []int, n int) []int {
	if len(s) == 0 {
		return []int{} // guard: n % 0 would panic
	}
	n %= len(s)
	if n < 0 {
		n += len(s) // negative n rotates right
	}

	result := make([]int, 0, len(s))
	result = append(result, s[n:]...)
	result = append(result, s[:n]...)
	return result
}

// Unique returns the items of s in order, keeping only the first occurrence of each.
func Unique(s []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(s))
	for _, item := range s {
		if seen[item] {
			continue
		}
		seen[item] = true
		result = append(result, item)
	}
	return result
}
