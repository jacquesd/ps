# Things
Implementation of the things style in go.

## Author
Jacques Dafflon <[jacques.dafflon@usi.ch](mailto:jacques.dafflon@usi.ch)>

## Programming Language
`go version go1.9.2 darwin/amd64`

## Instructions

1. Install `go1.9.2` ([installation instructions](https://golang.org/doc/install))
2. Build the program with the following command: 
`go build -i -o main main.go`
3. Run the program with the following command: 
`./main <input_file> <stop_words_file>`

## Things Style Constraints

- The larger problem is decomposed into *things* that make sense for the problem domain.
- Each *thing* is a capsule of data that exposes procedures to the rest of the world.
- Data is never accessed directly, only through these procedures.
- Capsules can re-appropriate procedures defined in other capsules.

## Observations
- There are no concept of objects or inheritance in Go. Struct are used to implement classes.
- Because there are no classes, there are no constructors. A common pattern to have properly instantiated `struct`s is to put them in a spearate package (in this case `things`), to make them private (lower case first letter) and to only make public a constructor function (in this case `New...`).
- Variables declared without an explicit initial value are given their zero value. For a `map[string]int`, the default value for a string is `0`. Therefore `wordFreqs[word] += 1` because `wordFreqs[word]`'s default value is `0`.
- Some of the built-in Python features used in the sample programs are not available in go and where implemented in `things/util.go`