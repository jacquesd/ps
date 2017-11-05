# Quarantine
Implementation of the the one style in go.

## Author
Jacques Dafflon <[jacques.dafflon@usi.ch](mailto:jacques.dafflon@usi.ch)>

## Programming Language
`go version go1.9.1 darwin/amd64`

## Instructions

1. Install `go1.9.1` ([installation instructions](https://golang.org/doc/install))
2. Build the program with the following command: 
`go build -i -o quarantine quarantine.go`
3. Run the program with the following command: 
`./quarantine <input_file> <stop_words_file>`

## Quarantine Style Constraints

- Core program functions have no side effects of any kind, including IO. 
- All IO actions must be contained in computation sequences that are clearly separated from the pure functions.
- All sequences that have IO must be called from the main program.

## Observations
- Some of the built-in Python features used in the sample programs are not available in go and where implemented at the beginning of the program.
- All functions are of type `callable`, defined as
  ```
  type callable func(args ...generic) []generic
  ```
  That is they all take an unspecified number of a `generic` arguments and return a slice of `generic` values.
- Parameters always have their types assertted at the beginning of a function using go's `variable.(type)` assertion notation.
- Return values are always wrapped into `generic` slices. 
- Since all functions return a slice of values, `guardCallable` checks and if needed calls each value in the slice.
- `getNextOsArg` returns the next OS argument every time it is called. In order to do so it needs to keep track of which arguments where returned. This global `coutner` state is wrapped in a closure for cleanliness.
  
  The actual `getNextOsArg` is pure. All it does is return the exact same `callable` which is then called from the main program. This callable will access and modify the "global" `counter` and read the OS arguments. 