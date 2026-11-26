# go-basics

Small Go programs from the Auriga Go course (Chapters 3–7).

| Program | Run |
|---|---|
| `cmd/profile` | `go run ./cmd/profile` |
| `cmd/temperature` | `go run ./cmd/temperature 37` |
| `cmd/fizzbuzz` | `go run ./cmd/fizzbuzz` |
| `cmd/safediv` | `go run ./cmd/safediv` |
| `cmd/variadic` | `go run ./cmd/variadic` |
| `cmd/wordfreq` | `go run ./cmd/wordfreq cmd/wordfreq/testdata/sample.txt` |
| `cmd/sliceshare` | `go run ./cmd/sliceshare` |
| `cmd/shapes` | `go run ./cmd/shapes` |
| `cmd/calc` | `go run ./cmd/calc 10 / 2` (use `x` or `'*'` for multiply) |

Check: `go vet ./...` and `go test ./...`

## Note: the shared-backing-array trap

A slice does not own its data. Slicing an array (or another slice) creates a
view onto the **same** backing array, so writing through the slice changes the
original too.

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:3] // [2 3], pointing into arr
s[0] = 99
fmt.Println(s, arr)
```

Output (the array changed as well):

```
slice: [99 3]
array: [1 99 3 4 5]
```

**Fix:** copy the elements into a new slice that has its own backing array.

```go
arr := [5]int{1, 2, 3, 4, 5}
c := make([]int, 2)
copy(c, arr[1:3])
c[0] = 99
fmt.Println(c, arr)
```

Output (the array is untouched):

```
slice: [99 3]
array: [1 2 3 4 5]
```

The same trap happens with `append`: if the slice still has spare capacity,
`append` writes into the shared array instead of allocating a new one.
Use `copy`, `slices.Clone`, or a full slice expression `s[low:high:high]`
when you need an independent slice.

Run `go run ./cmd/sliceshare` to see both outputs.
