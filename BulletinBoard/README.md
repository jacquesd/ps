# BulletinBoard
Implementation of the bulletinboard style in go.

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

## BulletinBoard Style Constraints

- Larger problem is decomposed into entities using some form of abstraction.
- The entities are never called on directly for actions.
- Existence of an infrastructure for publishing and subscribing to events, aka the bulletin board.
- Entities post event subscriptions to the bulletin board and publish events to the bulletin board. The bulletin board infrastructure does all the event management and distribution.

## Observations
- Some of the built-in Python features used in the sample programs are not available in go and where implemented in `bulletinboard/util.go`