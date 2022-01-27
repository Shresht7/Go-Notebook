package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

//	Go offers several printing verbs designed to format general Go values.
func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	//	If the value is a struct, the %+v variant will include the struct's fields names
	fmt.Printf("struct2: %+v\n", p)

	//	The %#v variant prints a Go syntax representation of the value, i.e. the source snippet
	//	that would produce that value
	fmt.Printf("struct3: %#v\n", p)

	//	To print the type of a value, use %T
	fmt.Printf("Type: %T\n", p)

	//	%t to format booleans
	fmt.Printf("bool: %t\n", true)

	//	There are many formatting options for integers
	//	%d standard, base-10 formatting
	fmt.Printf("int: %d\n", p)
	//	%b prints a binary representation
	fmt.Printf("bin: %b\n", p)
	//	%c prints the character corresponding to the given integer
	fmt.Printf("char: %c\n", p)
	//	%x provides hex encoding
	fmt.Printf("hex: %x\n", 456)

	//	There are also several formatting options for floats.
	//	%f for basic decimal formatting use
	fmt.Printf("float1: %f\n", 78.6)
	//	%e and %E to format in scientific notation
	fmt.Printf("float2: %e\n", 1234000000.0)
	fmt.Printf("float3: %E\n", 1234000000.0)

	//	For basic string printing use %s
	fmt.Printf("str1: %s\n", "\"string\"")
	//	To double-quote strings, use %q
	fmt.Printf("str2: %s\n", "\"string\"")
	//	%x for hex
	fmt.Printf("str3: %x\n", "hex this")

	//	%p for pointers
	fmt.Printf("pointer: %p\n", &p)

	//	When formatting numbers you will often want to control the width and precision
	//	of the resulting figure. To specify the width of an integer, use a number after
	//	% in the verb. By default, the result will be right-justified and padded with space
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	//	To specify the number of printed floats
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	//	To left justify use the - flag
	fmt.Printf("width3: |%6.2f|%6.2f|\n", 1.2, 3.45)

	//	You may also want to control width when formatting strings in a table like format
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "bar")
	//	to left justify, use - flag as with numbers
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "bar")

	//	Printf outputs formatted string to os.Stdout.
	//	Sprintf formats and returns a string without printing it
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	//	You can also format+print to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
