# Merge Sort

# About
Using Go Programming Language implement merge sort.
(Bonus:- simulation of progress bar.)
# Theory
Top-down Merge Sort Implementation:
The top-down merge sort approach is the methodology which uses recursion mechanism. It starts at the Top and proceeds downwards, with each recursive turn asking the same question such as “What is required to be done to sort the array?” and having the answer as, “split the array into two, make a recursive call, and merge the results.”, until one gets to the bottom of the array-tree.

Example: Let us consider an example to understand the approach better.

Divide the unsorted list into n sublists, each comprising 1 element (a list of 1 element is supposed sorted).

Working of Merge Sort

Top-down Implementation

Repeatedly merge sublists to produce newly sorted sublists until there is only 1 sublist remaining. This will be the sorted list.

Merging of two lists done as follows:
The first element of both lists is compared. If sorting in ascending order, the smaller element among two becomes a new element of the sorted list. This procedure is repeated until both the smaller sublists are empty and the newly combined sublist covers all the elements of both the sublists.

# Components used:-
  1. Slice:- Dyanamic array.
  2. Rand:- used to generate random elements between a range and of a specific type. 
  3. Sleep:-used to pause the execution of the program.

# Imports used:-
  1. fmt
  2. Math/rand
  3. time

# How to Run
Install Go in your system (Refer this)
Set up system environment variable
Follow steps below
$ cd Directory where file is saved

$ go run sort.go

# Author

Top-down Merge Sort Implementation:
The top-down merge sort approach is the methodology which uses recursion mechanism. It starts at the Top and proceeds downwards, with each recursive turn asking the same question such as “What is required to be done to sort the array?” and having the answer as, “split the array into two, make a recursive call, and merge the results.”, until one gets to the bottom of the array-tree.

Example: Let us consider an example to understand the approach better.

Divide the unsorted list into n sublists, each comprising 1 element (a list of 1 element is supposed sorted).

Working of Merge Sort

Top-down Implementation

Repeatedly merge sublists to produce newly sorted sublists until there is only 1 sublist remaining. This will be the sorted list.

Merging of two lists done as follows:
The first element of both lists is compared. If sorting in ascending order, the smaller element among two becomes a new element of the sorted list. This procedure is repeated until both the smaller sublists are empty and the newly combined sublist covers all the elements of both the sublists.

Merging Two Lists
