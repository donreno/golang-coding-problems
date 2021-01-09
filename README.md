# golang-coding-problems

[![Coverage Status](https://coveralls.io/repos/github/donreno/golang-coding-problems/badge.svg?branch=main)](https://coveralls.io/github/donreno/golang-coding-problems?branch=main)

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
| Hash Table | [hashtable.go](internal/structs/hashtable.go) | [tests](internal/structs/hashtable_test.go) |


### arrays
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
| String compression | [string_compression.go](internal/arrays/string_compression.go) | [tests](internal/arrays/string_compression_test.go) |
| Rotate Matrix | [rotate_matrix.go](internal/arrays/rotate_matrix.go) | [tests](internal/arrays/rotate_matrix_test.go) |
| Zero Matrix | [zero_matrix.go](internal/arrays/zero_matrix.go) | [tests](internal/arrays/zero_matrix_test.go) |
| Is String rotation | [string_rotation.go](internal/arrays/string_rotation.go) | [tests](internal/arrays/string_rotation_test.go) |

### Linked List

| problem | implementation | tests |
| --- | --- | --- |
| Remove duplicates | [remove_dups.go](internal/linkedlist/remove_dups.go) | [tests](internal/linkedlist/remove_dups_test.go) |
| Sum Lists | [remove_dups.go](internal/linkedlist/sum_lists.go) | [tests](internal/linkedlist/sum_lists_test.go) |
| Palindrome | [palindrome.go](internal/linkedlist/palindrome.go) | [tests](internal/linkedlist/palindrome_test.go) |
| Palindrome (fastslow) | [palindrome_fastslow.go](internal/linkedlist/palindrome_fastslow.go) | [tests](internal/linkedlist/palindrome_fastslow_test.go) |
| Loop detection | [loop_detection.go](internal/linkedlist/loop_detection.go) | [tests](internal/linkedlist/loop_detection_test.go) |
| Partition | [partition.go](internal/linkedlist/partition.go) | [tests](internal/linkedlist/partition_test.go) |


### Stacks and queues

| problem | implementation | tests |
| --- | --- | --- |
| Three in one | [three_in_one.go](internal/stacks/three_in_one.go) | [tests](internal/stacks/three_in_one_test.go) |
| Min Stack | [min_stack.go](internal/stacks/min_stack.go) | [tests](internal/stacks/min_stack_test.go) |
| Stack of plates | [stack_of_plates.go](internal/stacks/stack_of_plates.go) | [tests](internal/stacks/stack_of_plates_test.go) |
| Queue via stacks | [queue_via_stacks.go](internal/stacks/queue_via_stacks.go) | [tests](internal/stacks/queue_via_stacks_test.go) |
| Sorted Stack | [sorted_stack.go](internal/stacks/sorted_stack.go) | [tests](internal/stacks/sorted_stack_test.go) |
| Animal Shelter | [animal_shelter.go](internal/stacks/animal_shelter.go) | [tests](internal/stacks/animal_shelter_test.go) |
