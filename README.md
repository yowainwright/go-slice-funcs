# Go Slice Funcs

Slice functions written in Go.

This repository is for learning purposes.

### Goals:

- Build a distributable library
- Improve my Go Lang skills
- Implement common CI features using Go Lang

---

## Installation

```sh
go get github.com/yowainwright/go-slice-funcs
```

---

## Usage

```go
import "github.com/yowainwright/go-slice-funcs/SliceMap"

func main() {
    // ...
    SliceMap([]string{"a", "b", "c"}, func(s string) string {
        return s + "!"
    })
    // => ["a!", "b!", "c!"]
}
```

---
## Functions

Each function is built to mimic the native JS function within Go Lang constraints.

---
### `SliceMap`

`SliceMap` is similar to `Array.prototype.map` in JavaScript.

> `SliceMap[T any](data []T, f func(T) T) []T`

It takes in a generic slice and a function to run operations on each element of the slice, returning a new slice with the result.

```go
SliceMap([]string{"a", "b", "c"}, func(s string) string {
    return s + "!"
})
// => ["a!", "b!", "c!"]
```

---

Made by [@yowainwright](https://github.com/yowainwright), MIT 2022
