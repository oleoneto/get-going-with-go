# Get Going with Go
Exploring the [`go`](https://go.dev) programming language.


## Useful links
- [Getting Started Guides](https://go.dev/doc)
- [Package Index](https://pkg.go.dev)
- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [Developing and publishing modules](https://go.dev/doc/modules/developing)

## Commands
- `go help`
- `go test`
- `go build`
- `go list` (show the install path of a package)
- `go install`
- `go mod init some/module` (enable dependency tracking)
- `go mod tidy` (similar to `npm install`)
- `go run {{module_path}}` (similar to `node {{filename}}`)
- `go mod edit -replace some/module=path`

## Syntax
- `=` assign a value to a variable
- `:=` declare and initialize a variable
- `func name(param type) type {}`
![Illustration](https://go.dev/doc/tutorial/images/function-syntax.png "Go function syntax")

### Types
- [Array (value type)](https://go.dev/doc/effective_go#arrays)
```go
// declares an array of strings with length 2
var names [2]string

// declares and initializes and array of strings with length 2
names := [2]string{"John", "Jane"}

// declares and initializes and array of strings (with length 2)
names := [...]string{"John", "Jane"}
```

- [Slice (preferred over arrays)](https://go.dev/blog/slices-intro)
```go
// declares and initializes a slice (slices have no specified length)
names := []string{"John", "Jane"}
```

### Useful Built-in Functions
- `len()` returns the length (# of assigned elements) of an array or slice
- `cap()` returns the capacity (maximum length of a segment) of an array or slice
- `make()` creates an object
