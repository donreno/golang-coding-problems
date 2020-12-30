# golang-coding-problems
Just coding problems solved in golang

## Testing the code
Tests are backed by [goblin](https://github.com/franela/goblin), totally recommended.

### Running tests
To run tests we just need to run tests either with `make test` on the command line or `go test` if you are running on windows. (this will also run benchmarks)

### Running tets with coverage
To run tests with coverage just run with `make coverage`

## Current examples
Currently i've implemented couple of examples from different sources and couple of data structs to be able to solve basic problems.
 - [Stack example](stack.go) and corresponding [tests](stack_test.go) (actually had to implement this because of balanced string problem but still is a good example)
 - [Graph example](graph.go) created in order to implente DFS and BFS in go, test can be found [here](graph_test.go).
 - [Queue example](queue.go) and corresponding tests [here](queue_test.go).
 - [Balanced strings example](balanced_strings.go) and corresponding [tests](balanced_strings_test.go)
 - [String permutations coding problem](permutations.go) and corresponding [tests](permutations_test.go)
 - [String has unique characters coding problem](string_with_unique_characters.go) and corresponding [tests](string_with_unique_characters_test.go)
