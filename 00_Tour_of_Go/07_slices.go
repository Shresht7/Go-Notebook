package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 1. */ showcaseSlices()
}

//	======
//	SLICES
//	======

func showcaseSlices() {
	//	Slices are dynamically sized arrays.
	//	type `[]T` is a slice with elements of type T
	//	A slice is formed by specifying two indices, a low and high bound separated by a colon
	//	When slicing, you may omit the high or low bounds to use their defaults. (0 for low bound, length of slice for high bound)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	//	Declaring a slice
	var slc []int = primes[1:4]
	fmt.Println(slc)

	//	Slices are like reference to an array
	//	They do not store any data, they just describe a section of an underlying array

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	//	Changing the elements of a slice modifies the corresponding elements of its underlying array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	//	Other slices that share the same underlying array will see those changes
	b[0] = "Harry"
	fmt.Println(a, b)
	fmt.Println(names)

	//	Slice literals are like array literals without the length
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	str := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
	}
	fmt.Println(str)

	//	The length of a slice is the number of elements it contains
	//	Capacity of a slice is the number of elements in the underlying array
	fmt.Printf("Length: %v Capacity: %v\n", len(q), cap(q))

	//	Zero value of a slice is `nil` (len: 0, cap: 0)

	//	Slices can be created using the builtin `make` function.
	//	This is how you create dynamically-sized arrays
	//	The `make` function allocates a zeroed array and returns a slice that refers to that array
	forged := make([]int, 5 /*, capacity here if needed */) //	len(forged) == 5
	fmt.Println(forged)

	//	Slices can contain other slices!
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []int
	fmt.Println(s)

	//	Go's builtin `append` function is used to add new element to slices
	//	The first parameter is the slice and the rest the values to append to the slice
	//	Returns a new slice contaning all the elements of the original slice and provided values
	//	Note: if the given slice does not have the capacity to accomodate all elements, a new array will be allocated
	//		and the returned slice will point to that array
	s = append(s, 0)
	fmt.Println(s)

	//	We can add more than one element
	s = append(s, 1, 2, 3, 4)

	//	for loops can be used to iterate over the slice
	for index, value := range s {
		fmt.Printf("Index: %v Value: %v\n", index, value)
	}

	//	Unused variables are an error in Go, labeling them with _ designates them as unused
	for _, value := range s {
		fmt.Println(value)
	}

}
