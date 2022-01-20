package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//	Go does not have the concept of classes
//	A method is a function with a special receiver argument
//	The receiver appears in its own argument list between the func keyword and the method name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	Methods are just functions. The method above is equivalent to
	```
	func Abs(v Vertex) float 64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}
	```
*/

//	Methods can be declared on non-struct types too
//	Note: you can only declare a method with a receiver whose type is defined in the same package as the method.
//	You cannot declare a method with a receiver whose type is defined in another package
//	(which includes the built-in types such as `int`)

//	Type definition for a custom type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//	Even pointers!
//	Methods with pointer receivers can modify the value to which the receiver points
//	Since methods often need to modify their receiver, pointer receivers are more common than value receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) //	Calling the method

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(2)
	fmt.Println("Scaled up by 2", v)
}

/*
	Functions with a pointer argument must take a pointer
	while methods with a pointer receiver can take either a value or a pointer.

	Conversely, Functions with a value argument must take a value
	while methods with a value argument can take either a value or a pointer.const
*/
