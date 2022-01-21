package main

import "fmt"

func main() {

	//	In Go, variables are explicitly declared and used by
	//	the compiler to check type-correctness of function calls

	//	Var declares one or more variables
	var a = "initial"
	fmt.Println(a)

	//	var b int, c int = 1, 2 works too
	var b, c int = 1, 2
	fmt.Println(b, c)

	//	Type-inference
	var d = true
	fmt.Println(d)

	//	Declaration but no initialization
	//	Go assigns a zero value. zero-value for int is 0
	var e int
	fmt.Println(e)

	//	:= short syntax is for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)

}
