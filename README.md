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


### Arrays and strings

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

### Trees and Graphs

| problem | implementation | tests |
| --- | --- | --- |
| Route between graph nodes | [route_between_nodes.go](internal/graphsntrees/route_between_nodes.go) | [tests](internal/graphsntrees/route_between_nodes_test.go) |
| Minimal tree | [minimal_tree.go](internal/graphsntrees/minimal_tree.go) | [tests](internal/graphsntrees/minimal_tree_test.go) |
| List of Depths | [list_of_depths.go](internal/graphsntrees/list_of_depths.go) | [tests](internal/graphsntrees/list_of_depths_test.go) |
| Build Order | [build_order.go](internal/graphsntrees/build_order.go) | [tests](internal/graphsntrees/build_order_test.go) |
| BST Sequences | [bst_sequences.go](internal/graphsntrees/bst_sequences.go) | [tests](internal/graphsntrees/bst_sequences_test.go) |
| Check Subtree | [check_subtree.go](internal/graphsntrees/check_subtree.go) | [tests](internal/graphsntrees/check_subtree_test.go) |

### Bit Operations

| problem | implementation | tests |
| --- | --- | --- |
| Binary to string  | [binary_to_string.go](internal/bits/binary_to_string.go) | [tests](internal/bits/binary_to_string_test.go) |
| Bit swaps required | [bit_swaps_required.go](internal/bits/bit_swaps_required.go) | [tests](internal/bits/bit_swaps_required_test.go) |
| Flip to win | [flip_to_win.go](internal/bits/flip_to_win.go) | [tests](internal/bits/flip_to_win_test.go) |
| Insert | [insert.go](internal/bits/insert.go) | [tests](internal/bits/insert_test.go) |
| Pairwise Swap | [pairwise_swap.go](internal/bits/pairwise_swap.go) | [tests](internal/bits/pairwise_swap_test.go) |


## Recursion and dynamic programming

| problem | implementation | tests |
| --- | --- | --- |
| Fibonacci  | [fibonacci.go](internal/recursionndynamic/fibonacci.go) | [tests](internal/recursionndynamic/fibonacci_test.go) |
| Count Ways  | [count_ways.go](internal/recursionndynamic/count_ways.go) | [tests](internal/recursionndynamic/count_ways_test.go) |
| Robot in a grid  | [robot_in_a_grid.go](internal/recursionndynamic/robot_in_a_grid.go) | [tests](internal/recursionndynamic/robot_in_a_grid_test.go) |
| Magic Index  | [magic_index.go](internal/recursionndynamic/magic_index.go) | [tests](internal/recursionndynamic/magic_index_test.go) |
| Powerset  | [power_set.go](internal/recursionndynamic/power_set.go) | [tests](internal/recursionndynamic/power_set_test.go) |
| Recursive Multiply  | [recursive_multiply.go](internal/recursionndynamic/recursive_multiply.go) | [tests](internal/recursionndynamic/recursive_multiply_test.go) |
| Towers of Hanoi  | [towers_of_hanoi.go](internal/recursionndynamic/towers_of_hanoi.go) | [tests](internal/recursionndynamic/towers_of_hanoi_test.go) |
| Perms without Dups  | [permutations_without_dups.go](internal/recursionndynamic/permutations_without_dups.go) | [tests](internal/recursionndynamic/permutations_without_dups_test.go) |
| Perms with Dups  | [permutations_with_dups.go](internal/recursionndynamic/permutations_with_dups.go) | [tests](internal/recursionndynamic/permutations_with_dups_test.go) |
| Stack of boxes  | [stack_of_boxes_recursive.go](internal/recursionndynamic/stack_of_boxes_recursive.go) | [tests](internal/recursionndynamic/stack_of_boxes_recursive_test.go) |
| Stack of boxes (Grouped by surface not 100% optimized)  | [stack_of_boxes.go](internal/recursionndynamic/stack_of_boxes.go) | [tests](internal/recursionndynamic/stack_of_boxes_test.go) |
| Boolean evaluation  | [boolean_evaluation.go](internal/recursionndynamic/boolean_evaluation.go) | [tests](internal/recursionndynamic/boolean_evaluation_test.go) |

## Sorting and searching

| problem | implementation | tests |
| --- | --- | --- |
| Quick Sort  | [quicksort.go](internal/sorting/quicksort.go) | [tests](internal/sorting/quicksort_test.go) |
| Merge Sort  | [mergesort.go](internal/sorting/mergesort.go) | [tests](internal/sorting/mergesort_test.go) |
| Binary Search  | [binary_search.go](internal/sorting/binary_search.go) | [tests](internal/sorting/binary_search_test.go) |
| Sorted merge  | [sorted_merge.go](internal/sorting/sorted_merge.go) | [tests](internal/sorting/sorted_merge_test.go) |
| Search in rotated array  | [search_in_rotated_array.go](internal/sorting/search_in_rotated_array.go) | [tests](internal/sorting/search_in_rotated_array_test.go) |
| Group Anagrams  | [group_anagrams.go](internal/sorting/group_anagrams.go) | [tests](internal/sorting/group_anagrams_test.go) |
