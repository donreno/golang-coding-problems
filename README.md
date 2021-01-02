# golang-coding-problems
Just coding problems solved in golang

## Testing the code
Tests are backed by [goblin](https://github.com/franela/goblin), totally recommended.

### Running tests
To run tests we just need to run tests either with `make test` on the command line or `go test` if you are running on windows. (this will also run benchmarks)

### Running tets with coverage
To run tests with coverage just run with `make coverage`

## Package organization
I have implemented different examples of coding problems, structure implementations and many others distributed in different packages in this code base.

### structs
The structs package contains all the structs I have implemented im golang either to re-learn how it's built or to solve an specific problem.

| struct | implementation | tests |
| --- | --- | --- | 
| Graph | [graph.go](internal/structs/graph.go) | [tests](internal/structs/graph_test.go) |
| Linked List | [linked_list.go](internal/structs/linked_list.go) | [tests](internal/structs/linked_list_test.go) |
| Queue | [queue.go](internal/structs/queue.go) | [tests](internal/structs/queue_test.go) |
| Stack | [stack.go](internal/structs/stack.go) | [tests](internal/structs/stack_test.go) |

## arrays
Arrays package contains coding problems corresponding to arrays or strings.

| problem | implementation | tests |
| --- | --- | --- |
| Balanced Strings | [balanced_string.go](internal/arrays/balanced_string.go) | [tests](internal/arrays/balanced_string_test.go) |
| Count Permutations | [count_permutations.go](internal/arrays/count_permutations.go) | [tests](internal/arrays/count_permutations_test.go) |
| Permutations | [permutation.go](internal/arrays/permutation.go) | [tests](internal/arrays/permutation_test.go) |
| String with unique characters | [string_with_unique_characters.go](internal/arrays/string_with_unique_characters.go) | [tests](internal/arrays/string_with_unique_characters_test.go) |
| Urlify | [urlify.go](internal/arrays/urlify.go) | [tests](internal/arrays/urlify_test.go) |
| Palindrome permutation | [palindrome_permutation.go](internal/arrays/palindrome_permutation.go) | [tests](internal/arrays/palindrome_permutation_test.go) |
| One edit away | [one_away.go](internal/arrays/one_away.go) | [tests](internal/arrays/one_away_test.go) |