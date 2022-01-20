package main

import "fmt"

func main() {
	i, j := 42, 270

	//	A pointer holds the memory address of a value

	p := &i            //	p points to i. The `&` operator generates a pointer to its operand (var p *int)
	fmt.Println(p, *p) //	`*` operator denotes the pointer's underlying value. reads i through the pointer p
	*p = 21            //	Set i through the pointer p
	//	This is known as 'dereferencing' or 'indirecting'

	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

//	Unlike C, Go has no pointer arithmetic
