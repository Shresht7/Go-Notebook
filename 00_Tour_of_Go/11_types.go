package main

import "fmt"

//	A type assertion provides access to an interface value's underlying concrete value/
//	t = i.(T)
//	This statement asserts that the interface value i holds the concrete type T and assigns
//	the underlying value to variable t
//	If i does not hold T, the statement will trigger a panic

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	//	To test whether an interface value holds a specific type,
	//	a type assertion can return two values (the value itself, and a boolean to report whether the assertion succeeded)
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	typeSwitch()
}

func do(i interface{}) {
	//	A type switch is a contstruct that permits several type assertions in series
	//	A type-switch is like a regular switch statement, but the cases in a type-switch specify types.
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeSwitch() {
	do(21)
	do("hello")
	do(true)
}
