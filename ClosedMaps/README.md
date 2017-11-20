# ClosedMaps
Implementation of the closedmaps style in go.

## Author
Jacques Dafflon <[jacques.dafflon@usi.ch](mailto:jacques.dafflon@usi.ch)>

## Programming Language
`go version go1.9.2 darwin/amd64`

## Instructions

1. Install `go1.9.2` ([installation instructions](https://golang.org/doc/install))
2. Build the program with the following command: 
`go build -i -o closedmaps closedmaps.go util.go`
3. Run the program with the following command: 
`./closedmaps <input_file> <stop_words_file>`

## ClosedMaps Style Constraints

- The larger problem is decomposed into things that make sense for the problem domain.
- Each thing is a map from keys to values. Some values are procedures/ functions.
- The procedures/functions close on the map itself by referring to its slots.

## Observations
- Because Go is strongly typed, the objects have `generic` values and their actual types are asserted when needed.
- Objects are setup in the special [`init`](https://golang.org/doc/effective_go.html#init) method
- All the functions are anonymous functions defined in the objects directly.
- There is no logic written explicitely in the main (as in the [Python example](https://github.com/crista/exercises-in-programming-style/blob/master/12-closed-maps/tf-12.py#L43-L49)), all the logic is in the bojects and only functions in the objects are called from `main` with the correct parameters.
- Variables declared without an explicit initial value are given their zero value. For a `map[string]int`, the default value for a string is `0`. Therefore `wordFreqs[word] += 1` because `wordFreqs[word]`'s default value is `0`.
- Some of the built-in Python features used in the sample programs are not available in go and where implemented in `letterbox/util.go`