package main

import (
	"fmt"
)

func main() {
	/* 1. */ showcaseStructs()
	/* 2. */ showcaseArrays()
}

//	=======
//	STRUCTS
//	=======

//	A struct is a collection of fields

type Vertex struct {
	X int
	Y int
}

func showcaseStructs() {
	fmt.Println("This is a struct", Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 5 //	Struct fields are accessed using a dot
	fmt.Println(v)

	//	Struct fields can also be accessed through a struct pointer
	//	to access X in struct pointer p we could write (*p).X
	//	However the language permits us to write p.X instead; without the explicit dereference.

	//	A struct literal denotes a newly allocated struct value by listing the values of its fields (order is irrelevant)
	newV := Vertex{X: 6} //	Y: 0 is implicit
	zeroV := Vertex{}    //	X: 0 and Y: 0

	fmt.Println(newV)
	fmt.Println(zeroV)
}

//	======
//	ARRAYS
//	======

func showcaseArrays() {
	//	`[n]T` is an array of n values of type T
	//	Arrays in Go have fixed length
	var a [2]string //	Declares an array
	a[0] = "Hello"  //	Assign value to index 0
	a[1] = "World"  //	Assign value to index 1

	fmt.Println(a[0], a[1], a)

	//	Alternate definition. Can use :=
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
