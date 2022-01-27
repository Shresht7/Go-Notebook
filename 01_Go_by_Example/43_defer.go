package main

import (
	"fmt"
	"os"
)

//	Defer is used to enure a function call is performed later in a program's context
//	usually for purposes fo cleanup.

func main() {

	//	Suppose we want to create a file, write to it, and then close it when we're done.
	f := createFile("/tmp/defer.txt")
	//	Immediately after getting a file object, we defer the closing of that file with closeFile
	//	This will be executed at the end of the enclosing function (main), after writeFile has finished
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")

	//	It is important to check for errors when closing a file, even ina deferred function.
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}
