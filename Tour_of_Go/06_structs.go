package main

import "fmt"

//	A struct is a collection of fields

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

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
