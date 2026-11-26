// Command wordfreq reads a text file and prints its ten most frequent words.
//
// Usage: go run ./cmd/wordfreq testdata/sample.txt
package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: wordfreq <file>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// Count words. The map must be created with make before writing to it.
	counts := make(map[string]int)
	for _, w := range strings.Fields(strings.ToLower(string(data))) {
		// Trim punctuation so "go," and "go" count as the same word.
		w = strings.TrimFunc(w, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		if w != "" {
			counts[w]++
		}
	}

	// Maps have no order, so copy the keys into a slice and sort that.
	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	slices.SortFunc(words, func(a, b string) int {
		// Highest count first; equal counts sorted alphabetically so output is stable.
		if c := cmp.Compare(counts[b], counts[a]); c != 0 {
			return c
		}
		return cmp.Compare(a, b)
	})

	top := min(10, len(words))
	for i, w := range words[:top] {
		fmt.Printf("%2d. %-12s %d\n", i+1, w, counts[w])
	}
}
