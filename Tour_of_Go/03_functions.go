package main

import "fmt"

//	Functions are declared using the `func` keyword
//	They can take zero or more arguments
//	Types come after the variable name!
func add(x int, y int) int {
	return x + y
}

//	If two or more consecutive named parameters share the same type
//	then you can omit the type from all but the last paramter
func subtract(x, y int) int {
	return x - y
}

//	A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

//	Returns values can be named and are treated as variables defined at the top of the function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//	A return statement without arguments returns the named values. This is called 'naked' return
	//	Should generally be used in short functions
	return
}

func main() {
	fmt.Println(add(30, 12))
	fmt.Println(subtract(30, 12))
	fmt.Println(swap("first", "second"))
	fmt.Println(split(42))
}
