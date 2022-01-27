//	We often need our programs to perform operations on a collection of data,
//	like selecting all items that satisfy a given predicate or mapping all
//	items to a new collection with a custom function.

package main

import (
	"fmt"
	"strings"
)

//	Index returns the first index of the target string t, or -1 of no match is found
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//	Include returns true if target string t is in the slice
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//	Any returns true if one of the strings in the slice satisfies the predicate f
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//	All returns true if all of the string in the slice satisfies the predicate f
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//	Filter returns a new slice containing all strings in the slice that satisfy the predicate f
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//	Map returns a new slice containing the results of applying the function to each string
//	in the original slice
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var fruits = []string{"peach", "apple", "mangoes", "pear", "pineapple"}

	fmt.Println(Index(fruits, "pear"))
	fmt.Println(Include(fruits, "grape"))

	fmt.Println(Any(fruits, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(fruits, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(fruits, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(fruits, strings.ToUpper))
}
