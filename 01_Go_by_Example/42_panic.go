package main

import "os"

//	A panic typically means something went unexpectedly wrong. Mostly we use it to fail
//	fast on errors that shouldn't occur during normal operation, or that we aren't prepared to
//	handle gracefully

func main() {

	// panic("a problem")

	//	A common use of panic is to abort a function if it returns an error value that we
	//	don't kow how to (or want to) handle.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}

//	Note unlike some languages, which use exceptions for handling many errors
//	In Go, it is idiomatic to use error indicating return values whenever possible.
