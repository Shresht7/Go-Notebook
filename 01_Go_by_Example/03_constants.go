package main

import (
	"fmt"
	"math"
)

//	Constant Declaration
const s string = "constant"

func main() {
	fmt.Println(s)

	//	A const statement can appear anywhere in the program
	const n = 50_000_000

	//	Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n //	Note d is untyped float
	fmt.Println(d)

	//	A numeric constant has no type until it is given one, such as explicit conversion
	fmt.Println(int64(d))

	//	A number can be given a type by using it in a context that requires one
	//	Such as a variable assignment of function call.
	fmt.Println(math.Sin(n))
}
