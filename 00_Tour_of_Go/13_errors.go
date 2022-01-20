package main

import (
	"fmt"
	"time"
)

//	Go programs express errors with `error` values
//	The `error` type is a built-in interface similar to `fmt.Stringer`
//	```
//	type error interface {
//		Error() string
//	}
//	```

//	As with fmt.Stringer, the fmt package looks for error interface when printing values.

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	//	Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`
	//	A `nil` error denotes success, a non-nil error denotes failure
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
