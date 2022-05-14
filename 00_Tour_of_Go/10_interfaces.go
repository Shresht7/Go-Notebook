package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//	An interface type is defined as a set of method signatures
//	A value of interface type can hold any value that implements those methods
type Abser interface {
	Abs() float64
}

//	Both MyFloat and Vertex satisfy the interface Abser

//	---------------------------------------------------------------------------

//	A type implements an interface by implementing its methods
//	There is no explicit declartion of intent, no "implements" keyword
//	Implicit interfaces decouple the definition of an interface from its implementation
//	Which could then appear in any package without prearrangement

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(Value: %v, Type: %T)\n", i, i)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  //	a MyFloat implements Abser
	a = &v //	a *Vertex implement Abser

	// a = v	//	Errors: Vertex does not implement Abser

	fmt.Println(a.Abs())

	//	Under the hood, interface values can be though of as a tuple of value and a concrete type
	//	(value, type)
	//	An interface value holds a value of a specific underlying concrete type
	//	Calling a method on an interface value executes the method of the same name on its underlying type

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

//	If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
//	Note: An interface value that holds a nil concrete value is itself not nil
//	A nil interface holds neither value nor concrete type.
//	Calling a method on a nil interface is a run-time error because there is no type inside the tuple.

//	The interface type that specifies zero methods is known as empty interface
//	interface{}
//	Empty interface may hold values of any type (every type implements at least zero methods)
//	Empty interfaces are used by code that handles values of unknown type,
//	For example, fmt.Println takes any number of arguments of type interface{}
