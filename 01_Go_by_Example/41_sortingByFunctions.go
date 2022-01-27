package main

import (
	"fmt"
	"sort"
)

//	We can use custom-sorts to sort a collection by something other than its natural order

//	In order to sort by a custom function in Go, we need a corresponding type.
//	Here we've created a byLength type that is just an alias for the builtin []string type
type byLength []string

//	We implement sort.Interface - Len, Less and Swap - on our type so we can use the sort
//	package's generic Sort function. Len and Swap will usually be similar across types and Less
//	will hold the actual custom sorting logic. In our case we want to sort in order of increasing
//	string length, so we use len(s[i]) and len(s[j]) here
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

//	With all that in place we can now implement our custom sort

func main() {

	fruits := []string{"peach", "banana", "mangoes", "pineapple"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

}

//	Running our program, we can se that a list is sorted by string length.
// 	By following the same pattern of creating a custom type, implementing the three Interface
//	method on that type, and then calling sort.Sort on a collection of that custom type,
//	we can sort Go slices by arbitrary functions.
