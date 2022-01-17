package main

import "fmt"

//	Variables are declared using the `var` keyword
var i int = 1
var j float32 = 3.14
var k string //	Declared but uninitialized

//	If an initializer is preset then the type can be omitted as it will be inferred
var a = 3

func main() {

	//	Inside a function body, `:=` syntax can be used inplace of a var declaration
	walrus := true

	fmt.Println(i, j, k)
	fmt.Println(a)
	fmt.Println(walrus)

	showTypes()
}

//	BASIC TYPES
//	===========

//	bool
//	string
//	int int8 int16 int32 int64
//	uint uint8 uint16 uint32 uint64 uintptr
//	byte	-	Alias for uint8
//	rune	-	Alias for int32 (represents Unicode code point)
//	float32 float64
//	complex64 complex128 - represented by two floats

//	The int, uint and uintptr are usually 32 bits on 32-bit systems and 64 on 64-bit systems
//	When you need an integer value, you should use int unless you have a good reason not to

//	Variables declared without initialization are given a zero value
//	0 for numeric types
//	false for boolean type
//	"" (empty string) for strings

func showTypes() {

	newWalrus := true

	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", j, j)
	fmt.Printf("Type: %T Value: %v\n", k, k)
	fmt.Printf("Type: %T Value: %v\n", newWalrus, newWalrus)

	//	Type conversion
	var x = int(j) //	j is float32
	fmt.Printf("Type: %T Value: %v\n", x, x)

	showConstant()

}

//	CONSTANTS
//	=========

//	Constants are declared like variables but use the `const` keyword
//	They can be characters, string, boolean, or numeric values
//	They cannot be declared using the := syntax

const Pi = 3.14

func showConstant() {
	fmt.Printf("Type: %T Value: %v\n", Pi, Pi)
}
