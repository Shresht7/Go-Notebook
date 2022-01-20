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
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  //	a MyFloat implements Abser
	a = &v //	a *Vertex implement Abser

	// a = v	//	Errors: Vertex does not implement Abser

	fmt.Println(a.Abs())

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
