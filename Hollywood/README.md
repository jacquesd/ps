# Hollywood
Implementation of the hollywood style in go.

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

## Hollywood Style Constraints

- Larger problem is decomposed into entities using some form of abstraction.
- The entities are never called on directly for actions.
- The entities provide interfaces for other entities to be able to register callbacks.
- At certain points of the computation, the entities call on the other entities that have registered for callbacks.

## Observations
- Some of the built-in Python features used in the sample programs are not available in go and where implemented in `hollywood/util.go`