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

It takes in a generic slice and a function to run operations on each element of the slice, returning a new slice with the mapped result.

```go
SliceMap([]string{"a", "b", "c"}, func(s string) string {
    return s + "!"
})
// => ["a!", "b!", "c!"]
```

---

### `SliceFilter`

`SliceFilter` is similar to `Array.prototype.filter` in JavaScript.

> `SliceFilter[T any](data []T, f func(T) bool) []T`

It takes in a generic slice and a function to run operations on each element of the slice, returning a new slice with the filtered result.

```go
SliceFilter([]string{"a", "b", "c", "a"}, func(s string) bool {
    return s == "a"
})
// => ["a", "a"]
```

---

### `SliceFind`

`SliceFind` is similar to `Array.prototype.find` in JavaScript.

> `SliceFind[T any](data []T, f func(T) bool) T`

```go
SliceFind([]string{"a", "b", "c", "a"}, func(s string) bool {
    return s == "a"
})
// => "a"
```

---

### `SliceReduce`

`SliceReduce` is similar to `Array.prototype.reduce` in JavaScript.

> `SliceReduce[T any](data []T, f func(T, T) T) T`

```go
SliceReduce([]string{"a", "b", "c"}, func(a, b string) string {
    return a + b
})
// => "abc"
```

---

### `SliceForEach`

`SliceForEach` is similar to `Array.prototype.forEach` in JavaScript.

> `SliceForEach[T any](data []T, f func(T))`

```go
var word = ""
SliceForEach([]string{"a", "b", "c"}, func(s string) {
    word += s
    return
})
// word => "abc"
```

### `SliceIncludes`

`SliceIncludes` is similar to `Array.prototype.includes` in JavaScript.

> `SliceIncludes[T any](data []T, f func(T) bool) bool`

```go
SliceIncludes([]string{"a", "b", "c"}, func(s string) bool {
    return s == "a"
})
// => true
```

---

## Development

setup
```sh
.bin/setup.sh
# sets up brew, go, and pre-commit
```

setup for coding
```sh
go get ./... && go mod tidy && go mod vendor
```

running unit tests
```sh
go test ./...
```


---

## Roadmap

- Complete all/most of the JS Array Prototype functions
- Add more config
  - Github config files

---

Made by [@yowainwright](https://github.com/yowainwright), MIT 2022
