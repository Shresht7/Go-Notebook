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

//	Type definition for a custom type
type MyFloat float64

//	You can declare methods on non-struct types
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
