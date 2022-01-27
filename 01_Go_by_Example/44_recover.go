package main

import "fmt"

//	Go makes it possible to recover from a panic, by using the recover builtin function
//	A recover can stop a panic from aborting the program and let it continue with
//	executing instead

//	e.g: A server wouldn't want to crash if one of the client connections exhibits a critical
//	error. Instead, the server would want to close that connection and continue serving
//	other client. In fact, this is what Go's net/http does by default for HTTP servers

//	This function panics
func mayPanic() {
	panic("a problem")
}

func main() {

	//	Recover must be called within a deferred function. When the enclosing function panics,
	//	the defer will activate and a recover call within in will catch the panic

	//	The return value of recover is the error raised in the call to panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error\n", r)
		}
	}()

	mayPanic()

	//	This code will not run because mayPanic will panic. The execution of main stops
	//	at the point of panic and resumes in the deferred closure.
	fmt.Println("After mayPanic()")
}
