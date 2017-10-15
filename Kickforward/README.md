# Kickforward
Implementation of the kick forward style in go.

## Author
Jacques Dafflon <[jacques.dafflon@usi.ch](mailto:jacques.dafflon@usi.ch)>

## Programming Language
`go version go1.9.1 darwin/amd64`

## Instructions

1. Install `go1.9.1` ([installation instructions](https://golang.org/doc/install))
2. Build the program with the following command: 
`go build -i -o kickforward kickforward.go`
3. Run the program with the following command: 
`./kickforward <input_file> <stop_words_file>`

## Kickforward Style Constraints

- Larger problem is decomposed using functionl abstraction.
- Functions take input, and produce output.
- No shared state between functions.
- The larger problem is solved by composing functions one after the other, in 
pipeline, as a faithful reproduction of mathematical function composition 
`f ⚬ g`.
- Each function takes an additional parameter, usually the last, which is
another function.
- That function parameter is applied at the end of the current function.
- That function parameter is given, as input, what would be the output of the 
current function.
- The larger problem is solved as a pipeline of functions, but where the next 
function to be applied is given as parameter to the current function.

## Observations

Some of the built-in Python features used in the sample programs are not 
available in go and where implemented at the beginning of the program.

Variables and function names where kept as similar as possible to the ones in 
the sample program while following go's naming standards such as camelCase.

While the fist version used type aliasing and a type hierarchy, this new version
takes advantage of the empty interface (`interface{}`) and type assertions for a
more generic approach.

### Changes
- Split the problem further, from 9 to 12 functions
- All functions are of the same type (`continuation`), that is they all take the 
same arguments and (in this case) none return a value
- The specific arguments needed by a function are extracted from the 
`genericArgs` array and their type is asserted to access the concrete type. This
does not cast the value to another type. A panic is triggered if a value does
not hold the concrete type.
([More on go type assertion](https://tour.golang.org/methods/15))
- The `continuation` functions are defined in an array instead of being 
hardcoded into each function. Each function takes the arguments it needs, this 
array of continuations and the `next` continuation. Each function will then call
the `next` continuation with the arguments it needs, and this array of 
continuations from which the first function is taken out and passed as the 
next's `next` continuation.

> Those changes allow for more clarity in the order the functions are executed 
and more flexibility in adding, removing and swapping functions. The use of the
empty interface (`interface{}`) common to all types and type assertion to 
recover the concrete type of variables alos removes the need to hard-code a full
static type hierarchy using type aliasing.