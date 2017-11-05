# The One
Implementation of the the one style in go.

## Author
Jacques Dafflon <[jacques.dafflon@usi.ch](mailto:jacques.dafflon@usi.ch)>

## Programming Language
`go version go1.9.1 darwin/amd64`

## Instructions

1. Install `go1.9.1` ([installation instructions](https://golang.org/doc/install))
2. Build the program with the following command: 
`go build -i -o theone theone.go`
3. Run the program with the following command: 
`./theone <input_file> <stop_words_file>`

## The One Style Constraints

- Existence of an abstraction to which values can be converted.
- This abstraction provides operations to: 
    1. wrap around values, so that they become the abstraction
    2. bind itself to functions, so to establish sequences of functions
    3. unwrap the value, so to examine the final result.
- Larger problem is solved as a pipeline of functions bound together, with unwrapping happening at the end.
- Particularly for The One style, the bind operation simply calls the given function, giving it the value that it holds, and holds on to the returned value.

## Observations
- Some of the built-in Python features used in the sample programs are not available in go and where implemented at the beginning of the program.
- All functions are of type `callable`, defined as
  ```
  type callable func(args ...generic) []generic
  ```
  That is they all take an unspecified number of a `generic` arguments and return a slice of `generic` values.
- The `TheOne` `struct` stores a slice of `generic` values which are applied to the next bound function and replaced by the sliced returned from the function call.
- Parameters always have their types assertted at the beginning of a function using go's `variable.(type)` assertion notation.
- Return values are always wrapped into `generic` slices. 